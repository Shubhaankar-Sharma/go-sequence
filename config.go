package sequence

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts/gen/walletupgradable"
	"github.com/0xsequence/go-sequence/contracts/gen/walletutils"
)

type WalletConfig struct {
	Threshold uint16              `json:"threshold"`
	Signers   WalletConfigSigners `json:"signers"`
}

type WalletConfigSigner struct {
	Weight  uint8          `json:"weight"`
	Address common.Address `json:"address"`
}

type WalletConfigSigners []WalletConfigSigner

func (s WalletConfigSigners) Len() int { return len(s) }
func (s WalletConfigSigners) Less(i, j int) bool {
	return s[i].Address.Hash().Big().Cmp(s[j].Address.Hash().Big()) < 0
}
func (s WalletConfigSigners) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s WalletConfigSigners) GetWeightByAddress(address common.Address) (uint8, bool) {
	for _, signer := range s {
		if signer.Address == address {
			return signer.Weight, true
		}
	}
	return 0, false
}

func AddressFromWalletConfig(walletConfig WalletConfig, context WalletContext) (common.Address, error) {
	imageHash, err := ImageHashOfWalletConfig(walletConfig)
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromWalletConfig: %w", err)
	}
	return AddressFromImageHash(imageHash, context)
}

func AddressFromImageHash(imageHash string, context WalletContext) (common.Address, error) {
	mainModule32 := [32]byte{}
	copy(mainModule32[12:], context.MainModuleAddress.Bytes())

	codePack, err := ethcoder.SolidityPack([]string{"bytes", "bytes32"}, []interface{}{walletContractBytecode, mainModule32})
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	codeHash := crypto.Keccak256(codePack)

	hashPack, err := ethcoder.SolidityPack(
		[]string{"bytes1", "address", "bytes32", "bytes32"},
		[]interface{}{[]byte{0xff}, context.FactoryAddress, common.FromHex(imageHash), codeHash},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	hash := crypto.Keccak256(hashPack)[12:]

	return common.BytesToAddress(hash), nil
}

func ImageHashOfWalletConfig(walletConfig WalletConfig) (string, error) {
	imageHash, err := ImageHashOfWalletConfigBytes(walletConfig)
	if err != nil {
		return "", err
	}
	return ethcoder.HexEncode(imageHash), nil
}

func ImageHashOfWalletConfigBytes(walletConfig WalletConfig) ([]byte, error) {
	imageHash, err := ethcoder.SolidityPack([]string{"uint256"}, []interface{}{walletConfig.Threshold})
	if err != nil {
		return nil, fmt.Errorf("sequence, WalletConfigImageHash: %w", err)
	}

	for _, signer := range walletConfig.Signers {
		mm := [32]byte{}
		copy(mm[:], imageHash)

		weight := signer.Weight
		address := signer.Address

		aux, err := ethcoder.AbiCoder([]string{"bytes32", "uint8", "address"}, []interface{}{mm, weight, address})
		if err != nil {
			return nil, err
		}
		imageHash = ethcoder.Keccak256(aux)
	}

	return imageHash, nil
}

func FindCurrentConfig(ctx context.Context, address common.Address, provider, authProvider *ethrpc.Provider, walletContext *WalletContext, knownConfigs []*WalletConfig, ignoreIndex, requireIndex bool) (*WalletConfig, error) {
	if requireIndex && ignoreIndex {
		return nil, fmt.Errorf("findCurrentConfig: can't ignore index and require index")
	}

	imageHash, config, err := FindCurrentImageHash(ctx, walletContext, provider, authProvider, address, knownConfigs)
	if err != nil {
		return nil, err
	}
	if config != nil {
		knownConfigs = append([]*WalletConfig{config}, knownConfigs...)
	}

	config, err = FindConfigForImageHash(ctx, walletContext, imageHash, authProvider, knownConfigs)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func FindLastWalletOfInitialSigner(ctx context.Context, signer common.Address, provider, authProvider *ethrpc.Provider, walletContext *WalletContext, knownConfigs []*WalletConfig, ignoreIndex, requireIndex bool) (common.Address, error) {
	if requireIndex && ignoreIndex {
		return common.Address{}, fmt.Errorf("findCurrentConfig: can't ignore index and require index")
	}

	authContract, err := walletutils.NewWalletUtils(walletContext.UtilsAddress, authProvider)
	if err != nil {
		return common.Address{}, err
	}

	logBlockHeight := big.NewInt(0)
	if !ignoreIndex {
		block, err := authContract.LastSignerUpdate(&bind.CallOpts{Context: ctx}, signer)
		if err != nil {
			return common.Address{}, err
		}
		logBlockHeight.Set(block)
	}
	if requireIndex && logBlockHeight.Sign() == 0 {
		return common.Address{}, fmt.Errorf("no lastSignerUpdate for %v", signer)
	}

	var endBlockHeight *uint64
	if logBlockHeight.Sign() != 0 {
		endBlockHeight = new(uint64)
		*endBlockHeight = logBlockHeight.Uint64()
	}
	logs, err := authContract.FilterRequiredSigner(&bind.FilterOpts{
		Start:   logBlockHeight.Uint64(),
		End:     endBlockHeight,
		Context: ctx,
	}, nil, []common.Address{signer})
	if err != nil {
		return common.Address{}, err
	}
	lastLog := logs.Event
	for logs.Next() {
		lastLog = logs.Event
	}
	if lastLog == nil {
		return common.Address{}, fmt.Errorf("publishConfig: wallet config last log not found")
	}
	return lastLog.Wallet, nil
}

func FindConfigForImageHash(ctx context.Context, walletContext *WalletContext, image common.Hash, authProvider *ethrpc.Provider, knownConfigs []*WalletConfig) (*WalletConfig, error) {
	for _, kc := range knownConfigs {
		imageHash, err := ImageHashOfWalletConfigBytes(*kc)
		if err != nil {
			return nil, err
		}
		if bytes.Equal(imageHash, image[:]) {
			return kc, nil
		}
	}

	authContract, err := walletutils.NewWalletUtils(walletContext.UtilsAddress, authProvider)
	if err != nil {
		return nil, err
	}

	imageHashHeight, err := authContract.LastImageHashUpdate(&bind.CallOpts{Context: ctx}, image)
	if err != nil {
		return nil, err
	}

	var endBlockHeight *uint64
	if imageHashHeight.Sign() != 0 {
		endBlockHeight = new(uint64)
		*endBlockHeight = imageHashHeight.Uint64()
	}
	logs, err := authContract.FilterRequiredConfig(&bind.FilterOpts{
		Start:   imageHashHeight.Uint64(),
		End:     endBlockHeight,
		Context: ctx,
	}, nil, [][32]byte{image})
	if err != nil {
		return nil, err
	}
	lastLog := logs.Event
	for logs.Next() {
		lastLog = logs.Event
	}
	if lastLog == nil {
		return nil, fmt.Errorf("publishConfig: wallet config last log not found")
	}

	return decodeRequiredConfig(lastLog)
}

func FindCurrentImageHash(ctx context.Context, walletContext *WalletContext, provider, authProvider *ethrpc.Provider, address common.Address, knownConfigs []*WalletConfig) (common.Hash, *WalletConfig, error) {
	walletContract, err := walletupgradable.NewWalletUpgradable(address, provider)
	if err != nil {
		return common.Hash{}, nil, err
	}

	currentImageHash, err := walletContract.ImageHash(&bind.CallOpts{Context: ctx})
	if err == nil {
		return currentImageHash, nil, nil
	}

	for _, kc := range knownConfigs {
		walletAddress, err := AddressFromWalletConfig(*kc, *walletContext)
		if err != nil {
			return common.Hash{}, nil, err
		}
		if walletAddress == address {
			imageHash, err := ImageHashOfWalletConfigBytes(*kc)
			if err != nil {
				return common.Hash{}, nil, err
			}
			return common.BytesToHash(imageHash), kc, nil
		}
	}

	authContract, err := walletutils.NewWalletUtils(walletContext.UtilsAddress, authProvider)
	if err != nil {
		return common.Hash{}, nil, err
	}

	knownImageHash, err := authContract.KnownImageHashes(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	if knownImageHash != (common.Hash{}) {
		walletAddress, err := AddressFromImageHash(common.Bytes2Hex(knownImageHash[:]), *walletContext)
		if err != nil {
			return common.Hash{}, nil, err
		}
		if walletAddress != address {
			return common.Hash{}, nil, fmt.Errorf("findCurrentImageHash: inconsistent RequireUtils results")
		}
		return knownImageHash, nil, nil
	}

	logs, err := authContract.FilterRequiredConfig(&bind.FilterOpts{Context: ctx}, []common.Address{address}, nil)
	if err != nil {
		return common.Hash{}, nil, err
	}
	if logs.Next() {
		config, err := decodeRequiredConfig(logs.Event)
		if err != nil {
			return common.Hash{}, nil, err
		}

		gotImageHash, err := ImageHashOfWalletConfigBytes(*config)
		if err != nil {
			return common.Hash{}, nil, err
		}

		walletAddress, err := AddressFromImageHash(common.Bytes2Hex(gotImageHash), *walletContext)
		if err != nil {
			return common.Hash{}, nil, err
		}
		if walletAddress != address {
			return common.Hash{}, nil, fmt.Errorf("findCurrentImageHash: inconsistent RequireUtils results")
		}
		return common.BytesToHash(gotImageHash), config, nil
	} else {
		return common.Hash{}, nil, fmt.Errorf("counterfactual image hash not found")
	}
}

func decodeRequiredConfig(event *walletutils.WalletUtilsRequiredConfig) (*WalletConfig, error) {
	// TODO
	return &WalletConfig{
		Threshold: uint16(event.Threshold.Uint64()),
	}, nil
}

func SortWalletConfig(walletConfig WalletConfig) error {
	signers := walletConfig.Signers
	sort.Sort(signers) // Sort the signers

	// Ensure no duplicates
	for i := 0; i < len(signers)-1; i++ {
		if signers[i].Address == signers[i+1].Address {
			return fmt.Errorf("signer duplicate detected in the wallet config")
		}
	}

	return nil
}

func IsWalletConfigUsable(walletConfig WalletConfig) (bool, error) {
	// if walletConfig.Threshold.Cmp(big.NewInt(0)) == 0 {
	if walletConfig.Threshold == 0 {
		return false, fmt.Errorf("invalid wallet config - wallet threshold cannot be 0")
	}
	totalWeight := uint64(0)
	for _, s := range walletConfig.Signers {
		totalWeight += uint64(s.Weight)
	}
	// if walletConfig.Threshold.Cmp(big.NewInt(0).SetUint64(totalWeight)) > 0 {
	if uint64(walletConfig.Threshold) > totalWeight {
		return false, fmt.Errorf("invalid wallet config - total weight of the wallet config is less than the threshold required")
	}
	return true, nil
}

func IsWalletConfigEqual(walletConfigA, walletConfigB WalletConfig) bool {
	imageHashA, err := ImageHashOfWalletConfig(walletConfigA)
	if err != nil {
		return false
	}
	imageHashB, err := ImageHashOfWalletConfig(walletConfigB)
	if err != nil {
		return false
	}

	return imageHashA == imageHashB
}

// TODO: can leave out for now
/*type WalletState struct {
	Context WalletContext
	Config  *WalletConfig

	// the wallet address
	Address common.Address

	// the chainID of the network
	ChainID *big.Int

	// whether the wallet has been ever deployed
	Deployed bool

	// the imageHash of the `config` WalletConfig
	ImageHash string

	// the last imageHash of a WalletConfig, stored on-chain
	LastImageHash string

	// whether the WalletConfig object itself has been published to logs
	Published bool
}*/

package contracts

import (
	_ "embed"

	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts/gen/gasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/ierc1271"
	"github.com/0xsequence/go-sequence/contracts/gen/niftyswap"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
	"github.com/0xsequence/go-sequence/contracts/gen/walletfactory"
	"github.com/0xsequence/go-sequence/contracts/gen/walletgasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/walletguest"
	"github.com/0xsequence/go-sequence/contracts/gen/walletmain"
	"github.com/0xsequence/go-sequence/contracts/gen/walletupgradable"
	"github.com/0xsequence/go-sequence/contracts/gen/walletutils"
)

var (
	WalletFactory,
	WalletMainModule,
	WalletMainModuleUpgradable,
	WalletGuestModule,
	WalletUtils,
	WalletGasEstimator,
	GasEstimator,
	IERC1271,
	ERC20Mock,
	IERC20,
	IERC721,
	IERC1155,
	NiftyswapExchange,
	NiftyswapFactory,
	WrapAndNiftyswap,
	_ ethartifact.Artifact
)

var (
	//go:embed artifacts/erc1155/mocks/ERC20Mock.sol/ERC20Mock.json
	artifact_erc20mock string
)

func init() {
	WalletFactory = artifact("WALLET_FACTORY", walletfactory.WalletFactoryMetaData.ABI, walletfactory.WalletFactoryMetaData.Bin)
	WalletMainModule = artifact("WALLET_MAIN", walletmain.WalletMainMetaData.ABI, walletmain.WalletMainMetaData.Bin)
	WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", walletupgradable.WalletUpgradableMetaData.ABI, walletupgradable.WalletUpgradableMetaData.Bin)
	WalletGuestModule = artifact("WALLET_GUEST", walletguest.WalletGuestMetaData.ABI, walletguest.WalletGuestMetaData.Bin)
	WalletUtils = artifact("WALLET_UTILS", walletutils.WalletUtilsMetaData.ABI, walletutils.WalletUtilsMetaData.Bin)
	WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", walletgasestimator.WalletGasEstimatorMetaData.ABI, walletgasestimator.WalletGasEstimatorMetaData.Bin, walletgasestimator.WalletGasEstimatorDeployedBin)
	GasEstimator = artifact("GAS_ESTIMATOR", gasestimator.GasEstimatorMetaData.ABI, gasestimator.GasEstimatorMetaData.Bin, gasestimator.GasEstimatorDeployedBin)

	IERC1271 = artifact("IERC1271", ierc1271.IERC1271ABI, "")

	IERC20 = artifact("IERC20", tokens.IERC20ABI, "")
	IERC721 = artifact("IERC721", tokens.IERC721ABI, "")
	IERC1155 = artifact("IERC1155", tokens.IERC1155ABI, "")

	NiftyswapExchange = artifact("NIFTYSWAP_EXCHANGE", niftyswap.NiftyswapFactoryMetaData.ABI, niftyswap.NiftyswapFactoryMetaData.Bin)
	NiftyswapFactory = artifact("NIFTYSWAP_FACTORY", niftyswap.NiftyswapFactoryMetaData.ABI, niftyswap.NiftyswapFactoryMetaData.Bin)
	WrapAndNiftyswap = artifact("WRAP_AND_NIFTYSWAP", niftyswap.WrapAndNiftyswapMetaData.ABI, niftyswap.WrapAndNiftyswapMetaData.Bin)

	ERC20Mock = ethartifact.MustParseArtifactJSON(artifact_erc20mock)
}

func artifact(contractName, abiJSON, bytecodeHex string, deployedBytecodeHex ...string) ethartifact.Artifact {
	var deployedBin []byte
	if len(deployedBytecodeHex) > 0 {
		deployedBin = common.FromHex(deployedBytecodeHex[0])
	}
	return ethartifact.Artifact{
		ContractName: contractName,
		ABI:          ethcontract.MustParseABI(abiJSON),
		Bin:          common.FromHex(bytecodeHex),
		DeployedBin:  deployedBin,
	}
}

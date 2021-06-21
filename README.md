go-sequence
===========

Sequence Wallet client written in Go.


## Usage

TODO: add docs.. etc. examples, etc.

## Developing the go-sequence library

1. `make boostrap` -- will install node modules of ./testutil/chain
2. `make start-test-chain` -- starts the test ethereum chain (id 4337)
3. (in a separate terminal) `make go-test` -- runs test suite


## Testing

Testing is super important, to run the tests just call `make test`. As well, you can
run the test-chain separately with `make start-test-chain` then in another terminal run `make go-test`.

**NOTE:** Go by default will execute tests in parallel if you run `go test -v ./...`, so ensure to pass `-p 1`
to set parallelization to just 1 (so it runs serially). The `make go-test` command is already set to do this.

## LICENSE

Apache 2.0

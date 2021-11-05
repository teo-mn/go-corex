# Corex 

EVM-compatible chain secured by the aBFT consensus algorithm.

## Building the source

Building `corex` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make corex
```
The build output is ```build/corex``` executable.

## Running `corex`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `corex` instance.

### Launching a network

Launching `corex` for a network:

```shell
$ corex --genesis /path/to/genesis.g
```

### Configuration

As an alternative to passing the numerous flags to the `corex` binary, you can also pass a
configuration file via:

```shell
$ corex --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ corex --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `corex validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ corex --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`corex` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ corex --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ corex --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ corex --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ corex --datadir /path/to/datadir account new # create new account
$ corex --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Lachesis has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/corex-mn/go-corex/app	0.033s
?   	github.com/corex-mn/go-corex/cmd/cmdtest	[no test files]
ok  	github.com/corex-mn/go-corex/cmd/corex	13.890s
?   	github.com/corex-mn/go-corex/cmd/corex/metrics	[no test files]
?   	github.com/corex-mn/go-corex/cmd/corex/tracing	[no test files]
?   	github.com/corex-mn/go-corex/crypto	[no test files]
?   	github.com/corex-mn/go-corex/debug	[no test files]
?   	github.com/corex-mn/go-corex/ethapi	[no test files]
?   	github.com/corex-mn/go-corex/eventcheck	[no test files]
?   	github.com/corex-mn/go-corex/eventcheck/basiccheck	[no test files]
?   	github.com/corex-mn/go-corex/eventcheck/gaspowercheck	[no test files]
?   	github.com/corex-mn/go-corex/eventcheck/heavycheck	[no test files]
?   	github.com/corex-mn/go-corex/eventcheck/parentscheck	[no test files]
ok  	github.com/corex-mn/go-corex/evmcore	6.322s
?   	github.com/corex-mn/go-corex/gossip	[no test files]
?   	github.com/corex-mn/go-corex/gossip/emitter	[no test files]
ok  	github.com/corex-mn/go-corex/gossip/filters	1.250s
?   	github.com/corex-mn/go-corex/gossip/gasprice	[no test files]
?   	github.com/corex-mn/go-corex/gossip/occuredtxs	[no test files]
?   	github.com/corex-mn/go-corex/gossip/piecefunc	[no test files]
ok  	github.com/corex-mn/go-corex/integration	21.640s
```

Also it is tested with [fuzzing](./FUZZING.md).



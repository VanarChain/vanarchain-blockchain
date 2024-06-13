# Vanar Chain

Welcome to the Vanar Chain repository! Vanar Chain is a blockchain platform designed to provide a robust, scalable, and secure environment for decentralized applications (DApps). This README file will guide you through the setup, usage, and contribution process for this project. Go to mainnet documentation to see details regarding Vanar Mainnet, please visit [Mainnet webpage](https://docs.vanarchain.com/).

## Introduction
Vanar Chain is a next-generation blockchain platform that aims to address the limitations of existing blockchain systems by offering high throughput, low latency, and enhanced security. It is built with a modular architecture that allows for flexibility and ease of integration with various applications. For more details, you can go to [VanarChain website](https://docs.vanarchain.com/).

### Green Vanar
Vanar is very mindful about running all nodes and infra on carbon-free and green data centers as per [Google Cloud Platform](https://cloud.google.com/sustainability/region-carbon) guidelines. A Vanar Validator Node with a less than 90% score won't be accepted. We encourage node validators to run in a region with high CFE% (>90). To get more information you can read more about Green Vanar chain vision [here](https://vanarchain.com/en/blog/empowering-the-future-with-vanar-a-ceos-vision-for-a-green-fast-and-ai-driven-ecosystem.html).


## Features
1. High Throughput: It can process a high volume of transactions within a very short time frame.
2. Low Latency: Ensures quick transaction confirmation times.
3. Secure: Implements advanced cryptographic algorithms and consensus mechanisms.
4. Modular Architecture: Easily extendable and customizable.
5. Smart Contracts: Supports the creation and deployment of smart contracts.
6. Decentralized: Fully decentralized with no single point of failure.
7. Fee: Fixed [Gas Fee](https://docs.vanarchain.com/getting-started/vanar-architecture/fixed-fees/gas-fees-tiers) for all transactions

## Building the source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth).

Building `geth` requires both a Go (version 1.21 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd`
directory.

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| :--------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`geth`** | Our main Ethereum CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ethereum network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `geth --help` and the [CLI page](https://geth.ethereum.org/docs/fundamentals/command-line-options) for command line options. |
|   `clef`   | Stand-alone signing tool, which can be used as a backend signer for `geth`.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
|  `devp2p`  | Utilities to interact with nodes on the networking layer, without running a full blockchain.                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|  `abigen`  | Source code generator to convert Ethereum contract definitions into easy-to-use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) page for details.                                  |
| `bootnode` | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                                                                                                               |
|   `evm`    | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                               |
| `rlpdump`  | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                |

## Running `geth`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://geth.ethereum.org/docs/fundamentals/command-line-options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance.

### Hardware Requirements

Minimum:

* CPU with 8 cores
* 32GB RAM
* 500GB free storage space to sync the Mainnet
* A stable and high-speed 5 Gbps internet connection.

Recommended:

* Fast CPU with 16 cores
* 64GB RAM
* High-performance SSD with at least 1TB of free space
* A stable and high-speed 10 Gbps internet connection.

### Steps to Run a Fullnode on the Vanar network

By far the most common scenario is people wanting to simply interact with the Vanar network: create accounts; transfer funds; deploy and interact with contracts. We can sync quickly to the current state of the network. For that steps are listed below: -

#### Step 1: Choose a blockchain network
1. The blockchain network you are participating in is Vanarchain.
2. Ensure you understand the network's architecture, consensus algorithm, and validators' roles.
3. Familiarize yourself with the network's documentation, technical specifications, and community resources.
#### Step 2: Set up the environment
1. Install a Linux-based operating system Ubuntu 22.04 with a compatible architecture (x86_64).
2. Configure the network settings, DNS, and firewall rules to allow port 30311 tcp/udp incoming connections.
#### Step 3: Install dependencies and Source Code
1. Install the required dependencies, such as: -
* git (for cloning the blockchain's repository)
* make (for building the blockchain's binary)
* gcc (for compiling the blockchain's code)

2. Clone the blockchain's repository using git clone:
[Git Clone](https://github.com/VanarChain/vanarchain-blockchain.git).
This is the URL of the GitHub repository. It points to a specific directory within the repository.

3. Change into the cloned repository directory:
```shell
cd vanarchain-blockchain
```
cd: This is a command that changes the current directory to a new location.
vanarchain-blockchain: This is the name of the repository, which is also the directory name.

4. Run the make command to build the blockchain's binary:
```shell
make all
```
5. Copies the geth file from the build directory to the parent directory. Copies the bootnode file from the build directory to the parent directory. Changes the current directory to the parent directory.
```shell
cp ./build/bin/geth  /usr/bin/geth
```

This command will:
* cp ./build/bin/geth  /usr/bin/geth: This is a command that copies a file from one location to another. It is executed as follows:
* cp: This is the command that stands for "copy".
* ./build/bin/geth: This is the source file to be copied. The ./ before build indicates that the directory is relative to the current directory.
* /usr/bin/geth: This is the destination directory and filename where the file will be copied to. /usr/bin is a standard directory on Unix-like systems where executable programs are typically stored. geth is the name of the file to which the geth file will be copied.

### Configuration

Configuration files are located in the [config directory](https://github.com/VanarChain/vanarchain-blockchain/blob/20c8ff475de523f0753772caf48659651d73bef1/eth/ethconfig/config.go). You can modify these files to change various settings such as network parameters, consensus algorithms, and more.

### Programmatically interfacing `geth` nodes

As a developer, sooner rather than later you'll want to start interacting with `geth` and the
Vanar network via your own programs and not manually through the console. To aid
this, `geth` has built-in support for a JSON-RPC based APIs ([standard APIs](https://ethereum.github.io/execution-apis/api-documentation/)
and [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `geth`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept WebSocket requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,txpool,web3`)
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `geth` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

*Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
Ethereum nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!*

### Interact with fullnode
Start up geth's built-in interactive [JavaScript console](https://geth.ethereum.org/docs/interacting-with-geth/javascript-console), (via the trailing console subcommand) through which you can interact using [web3 methods](https://web3js.readthedocs.io/en/v1.10.0/) (note: the web3 version bundled within geth is very old, and not up to date with official docs), as well as geth's own [management APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc). This tool is optional and if you leave it out you can always attach to an already running geth instance with geth attach.

### More
More details about running an [RPC node](https://docs.google.com/document/d/1FE-4FxaD-YMfLeT-FUwh7oONsf7i_XJz/edit#heading=h.1r9awflj7s3g) and [becoming a validator](https://docs.google.com/document/d/1TJQNycVlsDHd-HleazPJhJgpxqPayh6_/edit)

*Note: Although some internal protective measures prevent transactions from crossing over between the main network and the test network, you should always use separate accounts for play and real money. Unless you manually move accounts, geth will by default correctly separate the two networks and will not make any accounts available between them.*

## Forks
Forks are when major technical upgrades or changes need to be made to the network – they typically stem from Ethereum Improvement Proposals (EIPs) and change the "rules" of the protocol.

When upgrades are needed in traditional, centrally-controlled software, the company will just publish a new version for the end-user. Blockchains work differently because there is no central ownership. Ethereum clients must update their software to implement the new fork rules. Plus block creators (miners in a proof-of-work world, validators in a proof-of-stake world) and nodes must create blocks and validate against the new rules. More on consensus mechanisms

These rule changes may create a temporary split in the network. New blocks could be produced according to the new rules or the old ones. Forks are usually agreed upon ahead of time so that clients adopt the changes in unison and the fork with the upgrades becomes the main chain. 

### Vanarchain Forks
Vanarchain draws significant inspiration from the pioneering Ethereum blockchain. Vanarchain is the fork of the Ethereum blockchain. Vanarchain leverages Ethereum's established foundation and core functionalities.

1. **Lahore Fork**
* Implementation Block# 1_068_650 
* Fee tiers were introduced for the first time. Below are the details: -
* Tiers
  - tier1 := 21000 - 500000
  - tier2 := 500001 - 600000
  - tier3 := 600001 - 800000
  - tier4 := 800001 - 1,200,000
  - tier5 := 1,200,001 - 2,000,000
  - tier6 := 2,000,001 - 5,000,000
  - tier7 := 5,000,001 - 15,000,000
  - tier8 := 15,000,001 - 30,000,000

* Multiplier factor for tiers
  - tier1 = 1
  - tier2 = 1
  - tier3 = 500
  - tier4 = 1000
  - tier5 = 1500
  - tier6 = 2000
  - tier7 = 4000
  - Default = 8000

2. **Karachi Fork**
* Implementation Block# 2_244_020
* Fee tiers were relaxed. Details are below: -
* Tiers
  - tier1 = 21000 - 12,000,000
  - tier2 = 12,000,001 - 15,000,000
  - tier3 = 15,000,001 - 20,000,000
  - tier4 = 20,000,001 - 25,000,000
  - tier5 = 25,000,001 - 30,000,000
* Multiplier factor for tiers
  - tier1 = 1
  - tier1 = 1500
  - tier2 = 2000
  - tier3 = 3000
  - tier4 = 4000
  - Default = 8000

* Global gas cap was updated from 50,000,000 to 12,000,000
* Tx max size was updated from 128 to 64 kb

3. **Sialkot Fork**
* Implementation Block# 1_678_900
* Fee tiers were revised. The latest fork name is Sialkot with build version v1.1.0. Details are below: -
* Tiers
  - tier1 = 21000 - 12,000,000
  - tier2 = 12,000,001 - 15,000,000
  - tier3 = 15,000,001 - 20,000,000
  - tier4 = 20,000,001 - 25,000,000
  - tier5 = 25,000,001 - 30,000,000
* Multiplier factor for tiers
  - tier1 = 1
  - tier2 = 3000
  - tier3 = 6000
  - tier4 = 15000
  - tier5 = 30000
  - Default = 60000

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) are licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.

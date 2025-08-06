# Modularis

Modularis is a simple implementation of a blockchain network inspired by Ethereum, written in the Go programming language. This project is designed with a modular and easy-to-understand structure to demonstrate the core concepts of a blockchain.

## Features

* **Modular Architecture:** The codebase is divided into distinct packages such as `core`, `crypto`, `network`, and `types`.
* **P2P Networking:** Capable of establishing connections between nodes via different `Transport` layers. Currently, a `LocalTransport` is implemented to simulate peer-to-peer communication on a single machine.
* **Digital Signatures:** Utilizes the ECDSA algorithm for signing transactions and blocks to ensure security and integrity.
* **Blocks and Transactions:** Implements data structures for blocks and transactions, complete with signing and verification mechanisms.
* **Hashing:** Uses SHA256 for hashing block data to ensure immutability.

## Getting Started

### Prerequisites

* Go version 1.22.10 or higher must be installed.

### Installation

1.  First, clone the project from its Git repository:

    ```sh
    git clone <YOUR_REPOSITORY_URL>
    cd Modularis
    ```

2.  Install the project dependencies:

    ```sh
    go mod tidy
    ```

### Running the Project

To build and run the project, use the following command:

```sh
make run
```

This command first compiles the project using `go build` and then executes the binary located at `./bin/modularis`.

### Running Tests

To execute all the tests written for the project, use this command:

```sh
make test
```

This will run all test files within the project.

## Project Structure

```
.
├── core/               # Core blockchain logic (blocks, transactions)
│   ├── block.go
│   ├── block_test.go
│   ├── transaction.go
│   └── transaction_test.go
├── crypto/             # Cryptographic functions
│   ├── keypair.go
│   └── keypair_test.go
├── network/            # Network management and node communication
│   ├── local_transport.go
│   ├── local_transport_test.go
│   ├── server.go
│   └── transport.go
├── types/              # Custom data types (Address, Hash)
│   ├── address.go
│   └── hash.go
├── main.go             # Application entry point
├── go.mod              # Go module dependencies
├── Makefile            # Build and execution commands
└── LICENSE             # Project license
```

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/Mostbesep/Modularis/blob/main/LICENSE) file for more details.
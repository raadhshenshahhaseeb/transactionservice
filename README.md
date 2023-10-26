# Transaction Service

This repository contains the implementation of a transaction service designed to handle Ethereum transactions. The service is structured to manage gas prices, gas limits, and nonce for transactions. This service can also allow interactions to and with the blockchain.

## Key Components

### 1. [Bootstrapper](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/cmd/bootstrapper/bootstrapper.go)
Initializes the service and sets up necessary configurations.

### 2. [Main](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/cmd/main.go)
The main entry point for the service.

### 3. [Configuration](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/configuration/config.go)
Handles the configuration settings for the service.

### 4. [Logger](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/pkg/logger/logger.go)
Provides logging capabilities for the service.

### 5. [Node](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/pkg/node/node.go)
Manages the Ethereum node interactions.

### 6. [Signer](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/pkg/signer/signer.go)
Handles the signing of transactions.

### 7. [Transaction](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/pkg/transaction/transaction.go)
Core component that manages Ethereum transactions. It provides functionalities like sending transactions, waiting for receipts, simulating transactions, and more.

### 8. [Mock Transaction](https://github.com/raadhshenshahhaseeb/transactionservice/blob/master/pkg/transaction/txMock/mock_transaction.go)
A mock implementation of the transaction service for testing purposes.
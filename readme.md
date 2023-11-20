# Go KVDB

This is an implementation of a hash based key-value database that is written in Go.

In addition, I also plan on implementing Raft Consensus to make it a fault tolerant distributed key-value database

## Goals
1. Create a key value store in Go
    - Hashmap implementation
    - Persist data in binary JSON

2. implement Raft Consensus to make it a distributed database

3. Optional Goals
    - Concurrent/prevent race conditions
    - Transaction based Atomicity
## How to run
### client mode
This will start the CLI client and server. The CLI client will take commands to send to the database.
```
go run main.go client
```
#### CLI client commands
will retrieve the value paired with the key.
If no value is paired with the key, server will notify client of a failure
```
get key1
```

Will post the key and it's paired value into the database. If the key already exists, the key value pair will update to the already sent key
```
set key1 value1
```

Will delete the key value pair in the database. If the key does not exist in the database, the databse will no op and continue as normal.
```
delete key1
```

## How to test
### test everything
```
go test ./... cover
```
### individual unit tests
```
go test ./kvStore/
```

## Todo
- create a test system for the client to do large amounts of operations to the database.
    - This is for eventually testing fault tolerance of a distributed kv-store.
- Create an API for the KVstore for CRUD interoperability
- implement snapshots, for writing data in memory to storage.
- implement atomic transactions
- implement Raft Consensus Algorithm
    - Dockerize the project

### Operations, initializations and destruction
- Initialize, open file when application starts
- load BSON file into memory. 
    - Will result in string keys, any datatype.
    - While file is in memory, do operations
- Close file when application ends
    - Serialize data into bson

## Changelog
### 11/19/2023
- Created dockerized distributed prototype in javascript. Is in another repo.
- Implemented a server that will listen for GET, SET, and DELETE requests.
    - will only take requests whose data is in the format of data.KVpair
    - implemented a custom datatype for sending requests to the kvStore server (data.KVpair)
- Implemented a CLI client for GET, SET, and DELETE operations
    - set can only set 1 word values with no spaces, and will always be interpretted as a string.
    - will fix at a later date.
- project can now only be run in server mode or client mode
    - client mode will start the server and the CLI tool to take in operations and send them to the database.
    - server mode is not yet implemented.
### 11/7/2023
- Created unit tests for operations.go
- Also created test folder "tests" for future general testing
### 11/5/2023
- Simplified operations by changing KVmap to map[string]interface{}
    - This means that dev won't have to add a use case for each data type serializing and deserializing
- deleted test functions and methods made obsoleted by changing KVmap's structure.
- renamed data.db to data.bson
### 11/3/2023
- Added operation for user to delete key value pairs.
### 10/29/2023
- Added hash based scheme for getting/setting data
    - keys are queried through a hash and Value is stored as a byte slice with a tag that states the byte slice's data type.
    - Data that is in storage is loaded into memory and accessed through a hash scheme.
- Added operations to get and set integers and strings.
- Added conversion functions for serializing and deserializing int, int arrays, strings, string arrays from and to binary 
### 10/24/2023
- Set up project, preliminary prototyping
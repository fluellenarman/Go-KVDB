# Go KVDB

This is an implementation of a hash based key-value database that is written in Go.

In addition, I also plan on implementing Raft Consensus to make it a fault tolerant distributed key-value database

## Goals
1. Create a key value store in Go
    - Hashmap implementation
    - Persist data in binary JSON

2. implement Raft Consensus to make it a distributed database

3. Optional Goals
    - create CRUD API's for IPC
    - Concurrent/prevent race conditions
    - Transaction based Atomicity

## Todo
- create a test system for the client to do large amounts of operations to the database.
    - This is for eventually testing fault tolerance of a distributed kv-store.
- Create an API for the KVstore for CRUD interoperability
- implement delete, for deleting key-value pairs
- implement snapshots, for writing data in memory to storage.
- implement atomic transactions
- implement Raft Consensus Algorithm
    - Dockerize the project

### Operations, initializations and destruction
- Initialize, open file when application starts
- load file into memory. 
    - Will result in string keys, binary json values.
    - While file is in memory, do operations
- Close file when application ends
    - Serialize data into bson

## Changelog
### 10/29/2023
- Added hash based scheme for getting/setting data
    - keys are queried through a hash and Value is stored as a byte slice with a tag that states the byte slice's data type.
    - Data that is in storage is loaded into memory and accessed through a hash scheme.
- Added functionality to get and set integers and strings.
- Added conversion functions for serializing and deserializing int, int arrays, strings, string arrays from and to binary 
### 10/24/2023
- Set up project, preliminary prototyping
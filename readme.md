# Go KVDB

This is an implementation of a key-value database that is written in Go.

In addition, I also plan on implementing Raft Consensus into Go KVDB

## Goals
1. Create a key value store in Go
    - Hashmap implementation
    - Persist data in binary JSON
    
2. Transition Go kvdb a DBMS 
    - create CRUD API's for IPC
    - create a TKINTER GUI for local management
    - Concurrenct/prevent race conditions
    - Transaction based Atomicity

3. implement Raft Consensus to make it a distributed database

## Todo
- Implement file opening and closing
- implement the kvStore to load data into memory/a map
    - data will be a map with string keys, and []uint8 binary array as values. The binary array is a json
- Implement serializing/deserializing data for writing and loading

- (Reach goal)
    - Implement a simple paging mechanism

### Operations, initializations and destruction
- Initialize, open file when application starts
- load file into memory. 
    - Will result in string keys, binary json values.
    - While file is in memory, do operations
- Close file when application ends
    - Serialize data into bson

## Changelog
- 10/24/2023

Set up project, preliminary prototyping
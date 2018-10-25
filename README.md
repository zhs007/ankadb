# AnkaDB

[![Build Status](https://travis-ci.org/zhs007/ankadb.svg?branch=master)](https://travis-ci.org/zhs007/ankadb)

``AnkaDB`` is a scalable embedded database with ``golang``.  
``AnkaDB`` supports multiple database engines, now ``LevelDB``, and will soon support ``RocksDB``.  
``AnkaDB`` is now single-node and will soon support multiple nodes, using the ``raft`` protocol.  
``AnkaDB`` queries and modifies data with ``GraphQL``.  
``AnkaDB`` used ``GRPC`` & ``ProtoBuf v3``.  

[``TradingDB``](https://github.com/zhs007/tradingdb) is an implementation of ``AnkaDB``.

---
### Update

##### **v0.1**
- Complete basic single node.
- Support ``graphql`` query.
- Support ``leveldb``.
- Support http service.
- Support grpc service.
- Support local queries.

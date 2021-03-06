# AnkaDB

[![Build Status](https://travis-ci.org/zhs007/ankadb.svg?branch=master)](https://travis-ci.org/zhs007/ankadb)

``AnkaDB`` is a scalable embedded database with ``golang``.  
``AnkaDB`` supports multiple database engines, ``LevelDB`` and ``RocksDB``.  
``AnkaDB`` is now single-node and will soon support multiple nodes, using the ``raft`` protocol.  
``AnkaDB`` queries and modifies data with ``GraphQL``.  
``AnkaDB`` used ``GRPC`` & ``ProtoBuf v3``.  

[``TradingDB``](https://github.com/zhs007/tradingdb) is an implementation of ``AnkaDB``.

Key-Value sample is [here](https://github.com/zhs007/ankadb/blob/master/ankadb_test.go).  
GraphQL sample is [here](https://github.com/zhs007/ankadb/blob/master/graphql_test.go).

---
### Update

##### **v0.6**
- Support ``raft`` protocol.

##### **v0.5**
- Support ``RocksDB``.

##### **v0.3**
- Refactor AnkaDB & DBLogic.
- Add Key-Value interface.
- Add GraphQL sample.
- Add test.
- Add Development Log.

##### **v0.2**
- Refactor the error module and delete the CODE in the protobuf.
- Support ``graphiql``
- Replace ``chan`` with ``context``

##### **v0.1**
- Complete basic single node.
- Support ``graphql`` query.
- Support ``leveldb``.
- Support http service.
- Support grpc service.
- Support local queries.

---
### AnkaDB Development Log

[``Come Here``](https://github.com/zhs007/ankadb/blob/master/blog.md)

---
### How To Upgrade v2 To v3

- change ``*AnkaDB`` to ``AnkaDB``.
- change ``Config`` to ``*Config``.
- change ``MgrDB`` to ``GetDBMgr()``.
- remove your ``dblogic``, and use ``ankadb.NewBaseDBLogic`` new a DBLogic.
- ``Config.ListDB.PathDB`` doesn't require join ``Config.PathDBRoot``.
- you can remove all the query result definitions, and use ``MakeMsgFromResultEx`` to resolve the query result.
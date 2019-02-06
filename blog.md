# AnkaDB Development Log

### 2019-02-06

今天开始优化``GraphQL``部分。  
增加了``QueryTemplate``，减少了后期``Query``的字符串解析。  

Start optimizing the ``GraphQL`` section today.  
Added ``QueryTemplate`` to reduce string parsing.  

### 2019-01-27

今天开始增加``AnkaDB``的一组更底层的``key-value``接口。  
主要是感觉``GraphQL``的多了一层语法转换，对于简单逻辑来说，``key-value``足够了。  
这次的``key-value``并没有加在``DBLogic``层，而是直接加在底层的，毕竟都``key-value``了。  

关于分支管理，``master``是最新的release，所有的tag都会打在``master``上。  

Today I started adding a set of lower-level ``key-value`` interfaces to ``AnkaDB``.
For simple logic, ``key-value`` is sufficient.
The ``key-value`` interfaces directly added to the ``AnkaDB``, are not in ``DBLogic``.

About branch, the ``master`` is the latest release, and all the tags will be on the ``master``.
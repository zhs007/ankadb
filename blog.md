# AnkaDB Development Log

### 2019-02-28

今天发现``GraphQL``里的逻辑代码其实很难进行单元测试，所以调整了代码结构，将参数解析以后的代码移到basefunc里，这样可以更方便的做测试。  
然后还增加了v0.2升级到v0.3的注意事项。

I found that the logic code in ``GraphQL`` is difficult to unit test, so I  moved the code outside the parameter parsing to basefunc, so that these functions can be tested.  

Then I added the precautions for v0.2 upgrade to v0.3.

### 2019-02-10

今天增加了Dockerfile，主要是为了``RocksDB``编译加的，还调整了``travis``配置，直接用``Docker``。  
后面项目如果要Docker部署，可以直接基于这个Dockerfile来做。

I added the Dockerfile today, for ``RocksDB`` compilation, and adjusted the ``travis`` configuration.
If your project is to be deployed by Docker, it can be done based on this Dockerfile.

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
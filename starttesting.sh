rm -rf test/database_test
rm -rf test/database_testasync
rm -rf test/test001
rm -rf test/test001-001
rm -rf test/test001-002
rm -rf test/test_msg
rm -rf test/test_user

CGO_CFLAGS="-I/usr/local/rocksdb/include" \
CGO_LDFLAGS="-L/usr/local/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
    go test ./... -v
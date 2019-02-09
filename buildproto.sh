protoc -I proto/ proto/ankadb.proto --go_out=plugins=grpc:proto
protoc -I test/ test/test.proto --go_out=plugins=grpc:test
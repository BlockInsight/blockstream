protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../probe probe.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../types types.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../cached cached.proto
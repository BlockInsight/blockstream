protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../probe probe.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../types types.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../cached cached.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../worker worker.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../jobs jobs.proto
protoc --proto_path=./ --go_out=plugins=grpc,paths=source_relative:../storage storage.proto
rem protoc "greet/greetpb/greet.proto" --go_out=plugins=grpc:.

rem protoc --go_out=greet --go_opt=paths=source_relative --go-grpc_out==plugins=grpc:greet --go-grpc_opt=paths=source_relative greet/greetpb/*.proto

rem protoc greet/greetpb/greet.proto --go_out=. --go-grpc_out=.

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

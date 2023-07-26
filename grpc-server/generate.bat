del /q /S .\api\product\v1
protoc --proto_path=api/proto --go_out=api/product/v1 --go_opt=paths=source_relative --go-grpc_out=api/product/v1 --go-grpc_opt=paths=source_relative  api/proto/*.proto
# grpc-client


## Getting started

This microservice has simple server which connects to grpc api from another microservice.

## Protobuf

Protocol Buffers are a language and platform neutral mechanism for serializing structured data. It generates source code from proto files to a variety of languages.

Protobuf installation https://github.com/protocolbuffers/protobuf/releases

Alternative protobuf installation with choco. As an administrator call on console below commands:

```
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

choco install protoc
```


## Calling gRPC methods with evans

Evans installation https://github.com/ktr0731/evans/releases

Calling grpc endpoint from evans

```
evans --host localhost --port 9001 -r repl

show services
call Hello
```


### Reference Documentation
For further reference, please consider the following sections:

* [Grpc protocol](https://grpc.io/)

# grpc的例子

```shell
# init hello project
go mod tidy
# generate proto file
make pb 
# start server
make srv
# another terminal ,start  client send to server
make call
```
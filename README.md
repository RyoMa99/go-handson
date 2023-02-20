- [資料](https://itnext.io/bff-pattern-with-go-microservices-using-rest-grpc-87d269bc2434)

# mongo dbコンテナの起動
```bash
docker run --rm --name user-db -d -p 27017:27017  --mount "type=bind,src=$(pwd)/mongo-init.js,dst=/docker-entrypoint-initdb.d/mongo-init.js,readonly" mongo
```
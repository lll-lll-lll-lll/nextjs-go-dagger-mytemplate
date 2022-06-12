# dockerコマンド
- docker build --tag docker-go-server .
- docker run --publish 8080:8080  docker-go-server
- curl http://localhost:8080/albums
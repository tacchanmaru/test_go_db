### ERROR と 解決策
$ docker-compose up
エラー文：docker.errors.DockerException: Error while fetching server API version: HTTPConnection.request() got an unexpected keyword argument 'chunked'
↓
$ docker compose up
参照：https://github.com/itzg/docker-minecraft-server/issues/2578
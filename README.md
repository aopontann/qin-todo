# Qin Todo

## フロントエンド

## バックエンド

### 技術
- [Gin](https://github.com/gin-gonic/gin)
- [MySQL](https://www.mysql.com/jp/)
- [database/sql](https://pkg.go.dev/database/sql)
- [goose](https://github.com/pressly/goose)
- [air](https://github.com/cosmtrek/air)
- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)


## セットアップ
### API
1. リポジトリのクローン
```
$ git clone https://github.com/qin-todo-team-l/qin-todo-backend
```
2. コンテナイメージの作成と起動（M1チップ使用している場合、docker-compose.ymlの8, 25行目のコメントを外してから実行してください）
```
$ docker-compose up -d --build
```
3. 起動した後、このコマンドを実行して、`{"message":"pong"}`が返って来ればOK（初回起動の場合、起動するのに少し時間かかります）
```
$ curl http://localhost:18080/ping
```

### DB
1. DBマイグレーション
```
$ goose -dir=tools/database/migrations mysql "user1:pass@tcp(mysql:3306)/qin-todo" up
```
2. デモデータの作成
```
$ go run tools/database/seed.go
```

### モック
このコマンドを実行するだけで、API, DB, モックコンテナが起動します。
```
docker-compose up -d --build
```
もし、モックコンテナだけ起動したい場合、このコマンドを実行してください
```
docker-compose up -d prism
```

### 注意事項
toml, cnfなどの設定ファイルを変更した場合は、コンテナイメージの再構築をした方が良いため、変更後はこのコマンドを実行してください。
```
docker-compose up -d --build
```

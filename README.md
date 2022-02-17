# Qin Todo バックエンド

### セットアップ
1. リポジトリのクローン
```
$ git clone https://github.com/qin-todo-team-l/qin-todo-backend
```
2. コンテナイメージの作成と起動（M1チップ使用している場合、docker-compose.ymlの8, 25行目のコメントを外してから実行してください）
```
$ docker-compose up -d
```
3. 起動した後、このコマンドを実行して、`{"message":"pong"}`が返って来ればOK（初回起動の場合、起動するのに少し時間かかります）
```
$ curl http://localhost:18080/ping
```

DBの接続部分は後で書き足します...

# go_proj_backend
go言語のバックエンド開発(TEST)

- 初期化
    - go mod init my-go-webserver
- 依存関係の整理
    - go mod tidy
- サーバーの実行
    - go run main.go

- 再初期化を行うとき
    - go.modをtidyを削除する
    - 初期化するモジュール名を変更する
    - 初期化と依存関係の整理を行う

- to_doリスト
    - テスト処理の追加
    - dockerを使ってDB作成を行う

- APIの叩き方
    - Create
        - curl -X POST http://localhost:8080 -d '{"id":"3", "name":"Charlie"}' -H "Content-Type: application/json"
    - Get(create時もしくはupdate時に実施する)
        - curl http://localhost:8080?id=3
    - Update
        - curl -X PUT http://localhost:8080 -d '{"id":"3", "name":"Charlie Updated"}' -H "Content-Type: application/json"
    - Delete
        - curl -X DELETE http://localhost:8080?id=3




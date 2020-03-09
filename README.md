【色々いじってみてる】

- 開発環境を実行

`docker-compose up`

- ビルド 入って build

``` bash
docker-compose up -d
docker exec -it go_app bash
go build
```

- バイナリを実行(Linux環境用)

`./api`

実行に関してはワンバイナリで済む実装になりそうだけど運用次第で検討

たぶんないだろうけどMACでバイナリ実行したい時はクロスコンパイルがいるよ

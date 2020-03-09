【色々いじってみてる中】

//開発環境を実行
docker-compose up

// ビルド 入って build
docker-compose up -d
docker exec -it go_app bash
go build

// バイナリを実行(Linux環境用)
./api

たぶんないだろうけどMACでバイナリ実行したい時はクロスコンパイルがいるよ

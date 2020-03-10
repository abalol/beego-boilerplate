【色々いじってみてる】

## VSCode環境構築編
1. goをローカルにインストールしてパスを通す。  
[参考: goenvでgoをインストール 〜初心者向け〜](https://qiita.com/yut-kt/items/9f5ac1e788df61f64290)

1. echo $GOPATH してそのパスの配下にcd
1. `git clone git@github.com:abalol/beego-boilerplate.git app` する
1. `cd app`
1. beeとbeegoをインストールする
```
 go get -u github.com/astaxie/beego
 go get -u github.com/beego/bee
```
## 開発環境を実行

`docker-compose up`

## ビルド

``` bash
docker-compose up -d
docker exec -it go_app bash
go build
```

もしくは

```
docker-compose build
```

## バイナリを実行(Linux環境用)

```
docker-compose up -d
docker exec -it go_app bash
./api
```

## Linux以外の環境でバイナリの実行
一旦クロスコンパイルは無視してます。

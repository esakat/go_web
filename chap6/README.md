# 外部環境の構築メモ

外部環境である`NSQ`と`MongoDB`はローカルに入れるのなんかやなので
Dockerを使う

## NSQ

起動
```
$ docker pull nsqio/nsq:v1.0.0-compat
$ docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq:v1.0.0-compat /nsqlookupd
```

バックグラウンド実行
```
$ docker run -d --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq:v1.0.0-compat  /nsqlookupd
```

確認
```
$ curl localhost:4161
{"message":"NOT_FOUND"}
```

## MongoDB

```
$ docker pull mongo:3.6
$ 
```

起動
`mongodb`という名前で起動する
```
$ docker run --name mongodb -p 27017:27017 -d mongo:3.6
```

mongoコマンドの実行方法
コンテナのため、mongoコマンドがローカルにインストールされていない
mongoコマンドを打つ際は別のmongodbコンテナを起動して打つ必要あり

(例) testというDBに繋げている様子
```
$ docker run -it --link mongodb:mongo --rm mongo sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/test"'
```

# 必要なパッケージダウンロード

```
$ go get github.com/bitly/go-nsq // NSQドライバ
$ go get gopkg.in/mgo.v2 // MongoDBドライバ
$ go get github.com/garyburd/go-oauth/oauth
$ go get github.com/joeshaw/envdecode
```


# テスト用の作業

## mongoにテストデータ投入

```
$ docker exec -it mongodb bash
(mongodb) $ mongo
> use ballots
> db.pools.insert(
  {
    "title":"今の気分は?",
    "options":["happy","sad","fail","win"]
  }
)
```

## nsqにトピック作成

```
$ docker run -it --name lookupd_commander nsqio/nsq:v1.0.0-compat  
$ docker exec -it lookupd bash
$ docker run -it --link lookupd:testnsq --rm testnsq sh -c 'exec nsq_tail --topic="votes" --lookupd-http-address=localhost:4161'
```
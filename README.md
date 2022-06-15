# Go(Echo) Go Modules × Dockerで開発環境構築

### 環境
Go 1.18.3
Echo

### 一応、GO言語をインストールすること
##### ※ 下記不要
~~https://golang.org/dl/~~

### 使用方法
`docker-compose up -d` で開発環境構築完了

#### echo backendコンテナに移動する
※ docker ps で起動中コンテナ確認

docker exec -i -t echo-backend_back_1 sh

#### goモジュール参考
https://blog.framinal.life/entry/2021/04/11/013819
https://nishinatoshiharu.com/go-modules-overview/

#### goモジュールインストール go.modに記載されているライブラリーをインストールする
###### ※ go.mod ： モジュールを管理するファイル
###### ※ go.sum ： 依存モジュールのチェックサムの管理をしてるファイル

go mod tidy

もしくは

go get

### DB接続情報はconf/config.goを確認
### Usersテーブルを作成する

### migration実施
##### 参考URL https://qiita.com/tanden/items/7b4fb1686a61dd5f580d#golang-migratemigrate%E3%81%A8%E3%81%AF

#### dbフォルダに移動(echo backendコンテナ内)
cd /go/src/app/db

#### up (マイグレーション実行)
go run migrate.go -exec up

#### down
go run migrate.go -exec down

#### -f (force option)
go run migrate.go -exec up -f

### エディターデバッグ設定

#### Golandの設定方法
#### 参考 https://qiita.com/keitakn/items/f46347f871083356149b
#### Run → Debug → 0. Edit Configurations... から設定を作成
#### Go Remote のTemplateを使って作成し、Nameは、任意、Hostはlocalhost
#### Portは、delveのデフォルトポートの2345、On disconnectをAskにし設定する

#### Visual Studio Codeの設定方法
#### 参考 https://qiita.com/masataka715/items/f87afa3e7f2c4e640ba7
#### https://qiita.com/f_sugar/items/49871913fb4baf1ad8d2
#### .vscode/launch.jsonを読み込ませる

### 起動確認

### ホーム画面
### localhost:5566

### JWT TOKEN認証
### token取得
### jwt token 参考 https://qiita.com/unvavo/items/b344a3ded2df8fa65c58
### curl -X POST -d 'username=test' -d 'password=test' localhost:5555/login

#### token認証確認
#### curl -H "Authorization: Bearer トークン" localhost:5555/restricted/welcome

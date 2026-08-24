# legalmatch_api

🇯🇵 日本語 | [🇺🇸 English](README.en.md)

法律相談を「利用者」と「弁護士」でマッチングするサービスを想定した、Go製のREST APIです。JWTベースの認証、リフレッシュトークン、相談案件(agenda)管理を備えています。

## 技術スタック

| 分類 | 技術 |
|---|---|
| 言語 / フレームワーク | Go 1.24, [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io/) |
| DB | MySQL 8.0 |
| 認証 | JWT (access token + refresh token), bcrypt によるパスワードハッシュ化 |
| マイグレーション | [goose](https://github.com/pressly/goose) |
| 実行環境 | Docker Compose |

## 主な機能

- ユーザー登録 / ログイン / ログアウト
- JWTアクセストークン + リフレッシュトークンによる認証
- ロールベースのアクセス制御(user / lawyer 等)
- ニックネーム重複チェック
- 相談案件(agenda)のデータモデル・マイグレーション

## ディレクトリ構成

```
auth/         ログイン・トークン発行・リフレッシュ処理
config/       DB接続・JWTシークレット読み込み
controllers/  ユーザー関連のリクエストハンドラ
middlewares/  CORS・ロガー・JWT認証ミドルウェア
migrations/   goose によるDBマイグレーション
models/       GORMモデル(User, RefreshToken, Agenda)
requests/     リクエストのバリデーション用構造体
routes/       ルーティング定義
utils/        JWTパース等のユーティリティ
```

## セットアップ

### 前提

- Go 1.24 以上
- Docker / Docker Compose
- [goose](https://github.com/pressly/goose) CLI(マイグレーション実行用)

### 手順

1. リポジトリを取得し、環境変数ファイルを用意する

   ```bash
   git clone https://github.com/Twilightcross/legalmatch_api.git
   cd legalmatch_api
   cp .env.example .env
   ```

   `.env` の値(DB認証情報、`JWT_SECRET` など)は各自の環境に合わせて書き換えてください。

2. MySQLをDocker Composeで起動する

   ```bash
   docker compose up -d
   ```

3. マイグレーションを実行する

   ```bash
   make migrate-up
   ```

4. APIサーバーを起動する

   ```bash
   go run main.go
   ```

   デフォルトで `http://localhost:8080` で起動します。

## APIエンドポイント

| メソッド | パス | 説明 | 認証 |
|---|---|---|---|
| POST | `/api/v1/users/register` | ユーザー登録 | 不要 |
| GET | `/api/v1/users/check-nickname?nickname=xxx` | ニックネーム重複確認 | 不要 |
| GET | `/api/v1/users/myinfo` | 自分の情報を取得 | 必要 |
| POST | `/api/v1/auth/login` | ログイン(access/refresh token発行) | 不要 |
| POST | `/api/v1/auth/refresh-token` | access tokenの再発行 | refresh token |
| POST | `/api/v1/auth/logout` | refresh tokenの失効 | refresh token |

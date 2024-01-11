# Scheduraphy
## 画像を送ると文字を読み取ってGoogleカレンダーに予定を登録してくれるLINEBotです。
![Fy9KErbagAE6X2P](https://github.com/claustra01/scheduraphy/assets/108509532/ce6ff638-caa5-4a11-a65c-a2c6c162c0a8)

## 使用技術
- Golang
- React / Vite / Next.js
- PostgreSQL
- LINE Messaging API
- LIFF
- Azure Computer Vision
- Google OAuth
- Google Calndar API

## 開発環境での実行
```shell
docker compose up -d --build
```

## DBのマイグレーション
```shell
go run main.go migrate
```

## ACRへのpush
```shell
docker login scheduraphycontainer.azurecr.io
docker compose build --no-cache
docker compose push
```
- loginに必要なusernameとpasswordはポータルからACRのアクセスキーを開くと確認できる

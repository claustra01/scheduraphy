# Scheduraphy

## 開発環境での実行
```shell
docker compose up -d --build
```

## ACRへのpush
```shell
docker login scheduraphycontainer.azurecr.io
docker compose build --no-cache
docker compose push
```
- loginに必要なusernameとpasswordはポータルからACRのアクセスキーを開くと確認できる

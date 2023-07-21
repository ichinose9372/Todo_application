# ベースイメージを指定
FROM golang:1.17-alpine

# 以降のコマンドはGolang 1.17を使用してビルドおよび実行

# ワーキングディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o myapp .

# アプリケーションを実行
CMD ["./myapp"]

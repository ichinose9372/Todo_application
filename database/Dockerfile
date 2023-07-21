# ベースイメージを指定
FROM mysql:latest

# 環境変数の設定
ENV MYSQL_ROOT_PASSWORD=root_password
ENV MYSQL_DATABASE=your_database
ENV MYSQL_USER=your_user
ENV MYSQL_PASSWORD=your_password

# ポートの公開
EXPOSE 3306

# 初期データベースのセットアップ
COPY ./init.sql /docker-entrypoint-initdb.d/

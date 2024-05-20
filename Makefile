# バイナリの名前
SERVER_BIN = server
CLIENT_BIN = client

# ソースディレクトリとバイナリディレクトリ
SERVER_SRC_DIR = server
CLIENT_SRC_DIR = client

# サーバーとクライアントのソースファイル
SERVER_SRC = $(SERVER_SRC_DIR)/server.go
CLIENT_SRC = $(CLIENT_SRC_DIR)/client.go

# デフォルトのターゲットはすべてビルド
all: build

# ビルドターゲット
build: $(SERVER_BIN) $(CLIENT_BIN)

$(SERVER_BIN): $(SERVER_SRC)
    go build -o $(SERVER_BIN) $(SERVER_SRC)

$(CLIENT_BIN): $(CLIENT_SRC)
    go build -o $(CLIENT_BIN) $(CLIENT_SRC)

# クリーンターゲット
clean:
    rm -f $(SERVER_BIN) $(CLIENT_BIN)

# サーバーの実行ターゲット
run-server: $(SERVER_BIN)
    ./$(SERVER_BIN)

# クライアントの実行ターゲット
run-client: $(CLIENT_BIN)
    ./$(CLIENT_BIN)

.PHONY: all build clean run-server run-client

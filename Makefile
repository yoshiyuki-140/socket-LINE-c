# コンパイラ
CC = gcc

# コンパイルオプション
CFLAGS = -Wall -Wextra -O2

# ターゲットバイナリ
TARGETS = client server

# 依存ファイル
CLIENT_OBJS = client.o
SERVER_OBJS = server.o

all: $(TARGETS)

client: $(CLIENT_OBJS)
	$(CC) $(CFLAGS) -o client $(CLIENT_OBJS)

server: $(SERVER_OBJS)
	$(CC) $(CFLAGS) -o server $(SERVER_OBJS)

%.o: %.c
	$(CC) $(CFLAGS) -c $<

clean:
	rm -f $(TARGETS) *.o

.PHONY: all clean

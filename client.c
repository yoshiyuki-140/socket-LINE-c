// client.c

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <pthread.h>

#define PORT 8000
#define BUFSIZE 1000

void *receive_messages(void *arg)
{
    int sockfd = *(int *)arg;
    char buf[BUFSIZE];
    int len;

    while ((len = read(sockfd, buf, sizeof(buf) - 1)) > 0)
    {
        buf[len] = '\0';
        printf("%s", buf);
    }

    return NULL;
}

int main(int argc, char const *argv[])
{
    int sockfd;
    struct sockaddr_in server_addr;
    char buf[BUFSIZE];
    pthread_t tid;

    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0)
    {
        perror("ERR: socket");
        exit(1);
    }

    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = inet_addr("192.168.3.10");
    server_addr.sin_port = htons(PORT);

    if (connect(sockfd, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0)
    {
        perror("ERR: connect");
        close(sockfd);
        exit(1);
    }

    pthread_create(&tid, NULL, receive_messages, &sockfd);
    pthread_detach(tid);

    while (fgets(buf, sizeof(buf), stdin) != NULL)
    {
        write(sockfd, buf, strlen(buf));
    }

    close(sockfd);
    return 0;
}

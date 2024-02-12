#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main() {
    char *moisture_pipe_path = "../../moisturepipe";

    int pipe_fd = open(moisture_pipe_path, O_WRONLY);
    if (pipe_fd < 0) {
        fprintf(stderr, "Error opening named pipe\n");
        exit(1);
    }
    puts("here");
    char *msg = "Hello I am the senate!\n";

    ssize_t bytesWritten = write(pipe_fd, msg, strlen(msg));
    if (bytesWritten < 0) {
        fprintf(stderr, "Error writing string to named pipe");
        close(pipe_fd);
        exit(2);
    }

    return 0;
}
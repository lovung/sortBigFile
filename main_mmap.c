#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <sys/mman.h>
#include <unistd.h>
#include <fcntl.h>

#define LENGTH 1000
#define MAX_LENGTH_OF_ONE_LINE 15
#define DEBUG 0
#define INPUT_FILENAME "small_list.txt"
#define OUTPUT_FILENAME "out.txt"

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

int write2File(const int *arr, int length)
{
    int fd = open(OUTPUT_FILENAME, O_RDWR | O_CREAT, (mode_t)0600);
    size_t textsize = LENGTH * MAX_LENGTH_OF_ONE_LINE;
    lseek(fd, textsize, SEEK_SET);
    write(fd, "", 1);
    char *map = mmap(0, textsize + 1, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    memcpy(map, arr, textsize);
    msync(map, textsize + 1, MS_SYNC);
    munmap(map, textsize + 1);
    close(fd);
    return 0;
}

int readFromFile(int *arr)
{
    int fd = open(INPUT_FILENAME, O_RDONLY);
    size_t pagesize = getpagesize();
    char *region = mmap(
        (void *)(pagesize * (1 << 20)), pagesize,
        PROT_READ, MAP_FILE | MAP_PRIVATE,
        fd, 0);
    fwrite(region, 1, pagesize, stdout);
    int unmap_result = munmap(region, pagesize);
    close(fd);
    return 0;
}

int main(void)
{
    int i = 0;
    int *arr;
    arr = malloc(LENGTH * sizeof(int));

    readFromFile(arr);
#if DEBUG
    printf("Before:\n");
    for (int j = 0; j < LENGTH; j++)
    {
        printf("%d\n", *(arr + j)); /* write the line */
    }
#endif
    qsort(arr, LENGTH, sizeof(int), cmpfunc);
    FILE *fp;
    fp = fopen("out.txt", "w");
    for (int j = 0; j < LENGTH; j++)
    {
        fprintf(fp, "%d\n", *(arr + j));
#if DEBUG
        printf("%d\n", *(arr + j)); /* write the line */
#endif
        fclose(fp);
    }
    else
    {
        perror(filename); /* why didn't the file open? */
    }

    free(arr);

    return 0;
}
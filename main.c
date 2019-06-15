#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 1000000
#define DEBUG 0

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

int main(void)
{
    int i = 0;
    int *arr;
    arr = malloc(LENGTH * sizeof(int));
    static const char filename[] = "list.txt";
    FILE *file = fopen(filename, "r");
    if (file != NULL)
    {
        char line[20]; /* or other suitable maximum line size */
        while (fgets(line, sizeof line, file) != NULL)
        { /* read a line */
            *(arr + i++) = atoi(line);
        }

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
        }
        fclose(fp);
        fclose(file);
    }
    else
    {
        perror(filename); /* why didn't the file open? */
    }

    free(arr);

    return 0;
}

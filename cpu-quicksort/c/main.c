#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define LENGTH 1000000
#define DEBUG 0

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

void quicksort(int *x, int first, int last)
{
    int pivot, j, i;
    float temp;

    if (first < last)
    {
        pivot = first;
        i = first;
        j = last;

        while (i < j)
        {
            while (x[i] <= x[pivot] && i < last)
                i++;
            while (x[j] > x[pivot])
                j--;
            if (i < j)
            {
                temp = x[i];
                x[i] = x[j];
                x[j] = temp;
            }
        }

        temp = x[pivot];
        x[pivot] = x[j];
        x[j] = temp;
        quicksort(x, first, j - 1);
        quicksort(x, j + 1, last);
    }
}

int *arr;
int main(void)
{
    int i = 0;
    arr = malloc(LENGTH * sizeof(int));
    static const char filename[] = "resources/list_2.txt";
    clock_t t;
    t = clock();
    FILE *file = fopen(filename, "r");

    if (file != NULL)
    {
        char line[20]; /* or other suitable maximum line size */
        while (fgets(line, sizeof line, file) != NULL)
        { /* read a line */
            *(arr + i++) = atoi(line);
        }

        t = clock() - t;
        printf("Read: %fms\n", ((double)t) * 1000 / CLOCKS_PER_SEC);
        t = clock();
#if DEBUG
        printf("Before:\n");
        for (int j = 0; j < LENGTH; j++)
        {
            printf("%d\n", *(arr + j)); /* write the line */
        }
#endif
        // qsort(arr, LENGTH, sizeof(int), cmpfunc);
        quicksort(arr, 0, LENGTH);
        t = clock() - t;
        printf("Sort: %fs\n", ((double)t) * 1000 / CLOCKS_PER_SEC);
        t = clock();
        FILE *fp;
        fp = fopen("out/out.txt", "w");
        for (int j = 0; j < LENGTH; j++)
        {
            fprintf(fp, "%d\n", *(arr + j));
#if DEBUG
            printf("%d\n", *(arr + j)); /* write the line */
#endif
        }
        t = clock() - t;
        printf("Write: %fs\n", ((double)t) * 1000 / CLOCKS_PER_SEC);
        t = clock();
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

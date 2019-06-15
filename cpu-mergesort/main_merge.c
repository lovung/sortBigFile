#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 1000000
#define DEBUG 0

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
void merge(int arr[], int l, int m, int r)
{
    int i, j, k;
    int n1 = m - l + 1;
    int n2 = r - m;

    /* create temp arrays */
    int L[n1], R[n2];

    /* Copy data to temp arrays L[] and R[] */
    for (i = 0; i < n1; i++)
        L[i] = arr[l + i];
    for (j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    /* Merge the temp arrays back into arr[l..r]*/
    i = 0; // Initial index of first subarray
    j = 0; // Initial index of second subarray
    k = l; // Initial index of merged subarray
    while (i < n1 && j < n2)
    {
        if (L[i] <= R[j])
        {
            arr[k] = L[i];
            i++;
        }
        else
        {
            arr[k] = R[j];
            j++;
        }
        k++;
    }

    /* Copy the remaining elements of L[], if there 
       are any */
    while (i < n1)
    {
        arr[k] = L[i];
        i++;
        k++;
    }

    /* Copy the remaining elements of R[], if there 
       are any */
    while (j < n2)
    {
        arr[k] = R[j];
        j++;
        k++;
    }
}

/* l is for left index and r is right index of the 
   sub-array of arr to be sorted */
void mergeSort(int arr[], int l, int r)
{
    if (l < r)
    {
        // Same as (l+r)/2, but avoids overflow for
        // large l and h
        int m = l + (r - l) / 2;

        // Sort first and second halves
        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);

        merge(arr, l, m, r);
    }
}

int main(void)
{
    int i = 0;
    int *arr;
    int arr_size;
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
        // qsort(arr, LENGTH, sizeof(int), cmpfunc);
        arr_size = sizeof(arr) / sizeof(arr[0]);
        mergeSort(arr, 0, arr_size - 1);
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

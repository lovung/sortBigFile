#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 1000000
#define DEBUG 0

int cmpfunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}

// Function to Merge Arrays L and R into A.
// lefCount = number of elements in L
// rightCount = number of elements in R.
void Merge(int *A, int *L, int leftCount, int *R, int rightCount)
{
    int i, j, k;

    // i - to mark the index of left aubarray (L)
    // j - to mark the index of right sub-raay (R)
    // k - to mark the index of merged subarray (A)
    i = 0;
    j = 0;
    k = 0;

    while (i < leftCount && j < rightCount)
    {
        if (L[i] < R[j])
            A[k++] = L[i++];
        else
            A[k++] = R[j++];
    }
    while (i < leftCount)
        A[k++] = L[i++];
    while (j < rightCount)
        A[k++] = R[j++];
}

// Recursive function to sort an array of integers.
void MergeSort(int *A, int n)
{
    int mid, i, *L, *R;
    if (n < 2)
        return; // base condition. If the array has less than two element, do nothing.

    mid = n / 2; // find the mid index.

    // create left and right subarrays
    // mid elements (from index 0 till mid-1) should be part of left sub-array
    // and (n-mid) elements (from mid to n-1) will be part of right sub-array
    L = (int *)malloc(mid * sizeof(int));
    R = (int *)malloc((n - mid) * sizeof(int));

    for (i = 0; i < mid; i++)
        L[i] = A[i]; // creating left subarray
    for (i = mid; i < n; i++)
        R[i - mid] = A[i]; // creating right subarray

    MergeSort(L, mid);            // sorting the left subarray
    MergeSort(R, n - mid);        // sorting the right subarray
    Merge(A, L, mid, R, n - mid); // Merging L and R into A as sorted list.
    free(L);
    free(R);
}

int main(void)
{
    int i = 0;
    int *arr;
    int arr_size;
    arr = malloc(LENGTH * sizeof(int));
    static const char filename[] = "resources/list.txt";
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
        MergeSort(arr, LENGTH);
        FILE *fp;
        fp = fopen("out/out.txt", "w");
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

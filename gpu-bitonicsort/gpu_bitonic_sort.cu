/*
 * Parallel bitonic sort using CUDA.
 * Compile with
 * nvcc -arch=sm_11 bitonic_sort.cu
 * Based on http://www.tools-of-computing.com/tc/CS/Sorts/bitonic_sort.htm
 * License: BSD 3
 */

#include <stdlib.h>
#include <stdio.h>
#include <time.h>

/* Every thread gets exactly one value in the unsorted array. */
#define THREADS 128 // 2^8
#define BLOCKS 8192 // 2^14
#define NUM_VALS THREADS *BLOCKS

void print_elapsed(clock_t start, clock_t stop)
{
    double elapsed = ((double)(stop - start)) / CLOCKS_PER_SEC;
    printf("Elapsed time: %.3fs\n", elapsed);
}

float random_float()
{
    return (float)rand() / (float)RAND_MAX;
}

void array_print(float *arr, int length)
{
    int i;
    for (i = 0; i < length; ++i)
    {
        printf("%1.3f ", arr[i]);
    }
    printf("\n");
}

void array_fill(float *arr, int length)
{
    srand(time(NULL));
    int i;
    for (i = 0; i < length; ++i)
    {
        arr[i] = random_float();
    }
}

__global__ void bitonic_sort_step(float *dev_values, int j, int k)
{
    unsigned int i, ixj; /* Sorting partners: i and ixj */
    i = threadIdx.x + blockDim.x * blockIdx.x;
    ixj = i ^ j;

    /* The threads with the lowest ids sort the array. */
    if ((ixj) > i)
    {
        if ((i & k) == 0)
        {
            /* Sort ascending */
            if (dev_values[i] > dev_values[ixj])
            {
                /* exchange(i,ixj); */
                float temp = dev_values[i];
                dev_values[i] = dev_values[ixj];
                dev_values[ixj] = temp;
            }
        }
        if ((i & k) != 0)
        {
            /* Sort descending */
            if (dev_values[i] < dev_values[ixj])
            {
                /* exchange(i,ixj); */
                float temp = dev_values[i];
                dev_values[i] = dev_values[ixj];
                dev_values[ixj] = temp;
            }
        }
    }
}

/**
 * Inplace bitonic sort using CUDA.
 */
void bitonic_sort(float *values)
{
    float *dev_values;
    size_t size = NUM_VALS * sizeof(float);

    cudaMalloc((void **)&dev_values, size);
    cudaMemcpy(dev_values, values, size, cudaMemcpyHostToDevice);

    dim3 blocks(BLOCKS, 1);   /* Number of blocks   */
    dim3 threads(THREADS, 1); /* Number of threads  */

    int j, k;
    /* Major step */
    for (k = 2; k <= NUM_VALS; k <<= 1)
    {
        /* Minor step */
        for (j = k >> 1; j > 0; j = j >> 1)
        {
            bitonic_sort_step<<<blocks, threads>>>(dev_values, j, k);
        }
    }
    cudaMemcpy(values, dev_values, size, cudaMemcpyDeviceToHost);
    cudaFree(dev_values);
}

void writefile(char *filename, float *buffer, int num)
{
    FILE *fp;
    fp = fopen(filename, "w");
    for (int j = 0; j < num; j++)
    {
        fprintf(fp, "%0.0f\n", *(buffer + j));
    }
    fclose(fp);
}

int main(int argc, char *argv[])
{
    clock_t start, stop;

    if (argc != 3)
    {
        printf("Invalid argument count.  %s accepts 1-4 arguments, %d given\n",
               argv[0], argc);
        return -1;
    }

    float *values = (float *)malloc(NUM_VALS * sizeof(float));
    // array_fill(values, NUM_VALS);
    if (values == NULL)
    {
        printf("Insufficient host memory to allocate at %d", __LINE__);
        return -3;
    }

    start = clock();
    FILE *fin = fopen(argv[1], "r");
    for (int i = 0; i < NUM_VALS; i++)
    {
        if (EOF == fscanf(fin, "%f ", &values[i]))
        {
            break;
        }
    }

    bitonic_sort(values); /* Inplace */

    writefile(argv[2], values, NUM_VALS);

    stop = clock();
    print_elapsed(start, stop);
    free(values);
}
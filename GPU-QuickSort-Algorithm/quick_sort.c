#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
int DEBUG = 1;
int *data;

void readfile(char *filename, int *buffer, int num){
    FILE *fh;
    int i = 0;
    fh = fopen(filename, "r");
    if (fh != NULL)
    {
        char line[20]; /* or other suitable maximum line size */
        while (fgets(line, sizeof line, fh) != NULL)
        { /* read a line */
            *(buffer + i++) = atoi(line);
        }
        fclose(fh);
    }
}

void writefile(char *filename, int *buffer, int num) {
    FILE *fp;
    fp = fopen(filename, "w");
    for (int j = 0; j < num; j++)
    {
        fprintf(fp, "%d\n", *(buffer + j));
    }
    fclose(fp);
}

int * get_list(int len){

    int *suffix_list = (int *) malloc(len*sizeof(int));
    int i;

    for(i=0; i<len; i++){
        suffix_list[i] = i;         
    }
    return suffix_list;
}

void quicksort(int* x, int first, int last){
    int pivot,j,i;
	float temp;

    if (first<last) {
        pivot=first;
        i=first;
        j=last;

        while(i<j){
            while(x[i]<=x[pivot]&&i<last)
                i++;
            while(x[j]>x[pivot])
                j--;
            if(i<j){
                temp=x[i];
                x[i]=x[j];
                x[j]=temp;
            }
        }

        temp=x[pivot];
        x[pivot]=x[j];
        x[j]=temp;
        quicksort(x,first,j-1);
        quicksort(x,j+1,last);
    }
}
  
void print_suffix_list(int *list, int len){
    int i=0;
    for(i=0; i<len; i++){
        printf("%d", list[i]);
        if(i != (len - 1)) printf(" ");
    }
    printf("\n");
}

int main(int argc, char *argv[]){
	clock_t start, end;
	double runTime;

    if(argc != 4){
        printf("Usage: ./quicksort -num -filename -outfile \n");
        exit(-1);
    }
    
    int num = atoi(argv[1]);
    char *filename = argv[2];
    char *out = argv[3];

	start = clock();
    data = (int *)malloc((num)*sizeof(int));
    readfile(filename, data, num);

    // int *suffix_list = get_list(strlen(data));
    quicksort(data, 0, num);
    //print_suffix_list(suffix_list, data_len);
    
    writefile(out, data, num);
    
	end = clock();
	free(data);

	runTime = (end - start) / (double) CLOCKS_PER_SEC ;
	printf("%d %f\n", num, runTime);
}



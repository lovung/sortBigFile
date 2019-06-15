#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 1000000
#define DEBUG 1
#define MAX_LENGTH_OF_ONE_LINE 15
#define INPUT_FILENAME "small_list.txt"
#define OUTPUT_FILENAME "out.txt"

/* Link list node */
struct Node
{
    int data;
    struct Node *next;
};

void insert(struct Node *base, struct Node *node)
{
    node->next = base->next;
    base->next = node;
}

/* Function to print linked list */
void printList(struct Node *head)
{
    struct Node *temp = head;
    while (temp != NULL)
    {
        printf("%d  ", temp->data);
        temp = temp->next;
    }
}

struct Node *newNode(int data)
{
    struct Node *newNode = malloc(sizeof(struct Node));
    newNode->data = data;
    newNode->next = NULL;
    return newNode;
}

/* A utility function to insert a node at the beginning of linked list */
void push(struct Node **head_ref, int new_data)
{
    /* allocate node */
    struct Node *new_node = malloc(sizeof(struct Node));

    /* put in the data  */
    new_node->data = new_data;

    /* link the old list off the new node */
    new_node->next = (*head_ref);

    /* move the head to point to the new node */
    (*head_ref) = new_node;
}

void readAndSort(struct Node **head_ref)
{
    struct Node *tmp = NULL;

    FILE *file = fopen(INPUT_FILENAME, "r");
    char line[20]; /* or other suitable maximum line size */
    if (file != NULL)
    {
        fgets(line, sizeof line, file);
        tmp = newNode(atoi(line));
        (*head_ref) = tmp;
        while (fgets(line, sizeof line, file) != NULL)
        { /* read a line */
            int i = atoi(line);
            tmp = (*head_ref);
            if (i <= tmp->data)
            {
                (*head_ref) = newNode(i);
                (*head_ref)->next = tmp;
                continue;
            }

            while (tmp != NULL)
            {
                if (tmp->next != NULL)
                {
                    if (i <= tmp->next->data)
                    {
                        insert(tmp, newNode(i));
                    }
                }
                else
                {
                    insert(tmp, newNode(i));
                }
                tmp = tmp->next;
            }
        }
    }
    else
    {
        perror(INPUT_FILENAME); /* why didn't the file open? */
    }
    fclose(file);
}

void writeListToFile(struct Node *head)
{
    FILE *fp;
    fp = fopen(OUTPUT_FILENAME, "w");
    struct Node *temp = head;
    while (temp != NULL)
    {
        fprintf(fp, "%d\n", temp->data);
        temp = temp->next;
    }
    fclose(fp);
}

int main(void)
{
    struct Node *a = NULL;
#if 0 
    readAndSort(&a);
#if DEBUG
    printf("Linked List before sorting \n");
    printList(a);
#endif
    writeListToFile(a);
#if DEBUG
    printf("\nLinked List after sorting \n");
    printList(a);
#endif
#else
    push(&a, 5);
    push(&a, 20);
    push(&a, 4);
    push(&a, 3);
    push(&a, 30);

    printf("Linked List before sorting \n");
    printList(a);

    insertionSort(&a);

    printf("\nLinked List after sorting \n");
    printList(a);

#endif
    return 0;

    return 0;
}

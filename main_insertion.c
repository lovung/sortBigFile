#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LENGTH 1000000
#define DEBUG 0
#define MAX_LENGTH_OF_ONE_LINE 15
#define INPUT_FILENAME "small_list.txt"
#define OUTPUT_FILENAME "out.txt"

/* Link list node */
struct Node
{
    int data;
    struct Node *next;
};

// Function to insert a given node in a sorted linked list
void sortedInsert(struct Node **, struct Node *);

// function to sort a singly linked list using insertion sort
void insertionSort(struct Node **head_ref)
{
    // Initialize sorted linked list
    struct Node *sorted = NULL;

    // Traverse the given linked list and insert every
    // node to sorted
    struct Node *current = *head_ref;
    while (current != NULL)
    {
        // Store next for next iteration
        struct Node *next = current->next;

        // insert current in sorted linked list
        sortedInsert(&sorted, current);

        // Update current
        current = next;
    }

    // Update head_ref to point to sorted linked list
    *head_ref = sorted;
}

/* function to insert a new_node in a list. Note that this 
  function expects a pointer to head_ref as this can modify the 
  head of the input linked list (similar to push())*/
void sortedInsert(struct Node **head_ref, struct Node *new_node)
{
    struct Node *current;
    /* Special case for the head end */
    if (*head_ref == NULL || (*head_ref)->data >= new_node->data)
    {
        new_node->next = *head_ref;
        *head_ref = new_node;
    }
    else
    {
        /* Locate the node before the point of insertion */
        current = *head_ref;
        while (current->next != NULL &&
               current->next->data < new_node->data)
        {
            current = current->next;
        }
        new_node->next = current->next;
        current->next = new_node;
    }
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

void readListFromFile(struct Node **head_ref)
{
    FILE *file = fopen(INPUT_FILENAME, "r");
    if (file != NULL)
    {
        char line[20]; /* or other suitable maximum line size */
        while (fgets(line, sizeof line, file) != NULL)
        { /* read a line */
            push(head_ref, atoi(line));
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
    readListFromFile(&a);
#if DEBUG
    printf("Linked List before sorting \n");
    printList(a);
#endif
    insertionSort(&a);
    writeListToFile(a);
#if DEBUG
    printf("\nLinked List after sorting \n");
    printList(a);
#endif
    return 0;
}
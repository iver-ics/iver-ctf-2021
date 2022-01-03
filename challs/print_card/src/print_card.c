#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main()
{
  FILE *file;
  char flag[50];
  char name[50];

  setbuf(stdout, NULL);
  setbuf(stdin, NULL);
  setbuf(stderr, NULL);

  file = fopen("flag.txt", "r");
  fgets(flag, sizeof(flag), file);

  printf("What is your elf name?\n");

  fgets(name, sizeof(name), stdin);

  printf("Wellcome ");
  printf(name);
  printf("Let me print a card for you :-)\n");

  return 0;
}
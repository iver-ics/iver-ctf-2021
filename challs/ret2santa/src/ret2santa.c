#include <stdio.h>
#include <stdlib.h>

// Kompilerad med: clang ret2santa.c -o ret2santa

void main()
{
    char idea[30];

    setbuf(stdout, NULL);
    setbuf(stdin, NULL);
    setbuf(stderr, NULL);

    puts("Christmas is getting cold; we need a win NOW! Please help us!: ");
    gets(idea);
    puts("\nLet's try that!");
}

int win()
{
    puts("You're awesome! Here's a flag for you :-)\n");
    system("cat flag.txt");
    exit(0);
}
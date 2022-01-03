#include <stdio.h>
#include <stdlib.h>

// Compiled with:
//  clang pwn_santa_2.c -o pwn_santa_2

void vuln()
{
    int beutiful = 0;
    char poem[1024];

    puts("Hi, please read me a christmas poem?");
    gets(poem);

    if (beutiful)
    {
        puts("That's beutiful, here! I'm gonna use this flag to wipe my tears  :')");
        system("/bin/cat flag.txt");
    }
    else
    {
        puts("Eum. Okey...");
    }
}

int main()
{
    setbuf(stdout, NULL);
    setbuf(stdin, NULL);
    setbuf(stderr, NULL);

    vuln();
}

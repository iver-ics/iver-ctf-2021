#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

unsigned long long int values[] =
    {
        // REPLACETHIS
};

int main()
{
  setbuf(stdout, NULL);
  setbuf(stdin, NULL);
  setbuf(stderr, NULL);

  unsigned long long int value = 0ULL;

  for (int i = 0; i < sizeof(values) / 8; i++)
  {
    value ^= values[i];
    unsigned long left = value >> 32;
    unsigned long right = value & 0xFFFFFFFF;

    if (left == 0x13371337)
    {
      printf("Whoop! Found one :-)", right);
    }
  }
}

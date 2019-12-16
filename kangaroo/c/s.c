#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

const int UPPERBOUND = 10000;

char* kangaroo(int k1s, int k1d, int k2s, int k2d)
{
  // Complete this function
  //printf("k1s and k2s are respectively: %i and %i\n", k1s, k1d);
  //printf("k2s and k2d are respectively: %i and %i\n\n", k2s, k2d);

  //printf("((<hop n> * <k1d>) + <k1s>) == ((<hop n> * <k2d>) + <k2s>): | k1j == kj2 |\n");

  for ( int j = 1; j <= UPPERBOUND; j++ )
  {
    //printf("(%-8i * %8i + %-8i) == (%-8i * %8i + %-8i): | %8i == %8i |\n",
    //       j,   k1d,    k1s,      j,   k2d,    k2s, (k1d * j + k1s), (k2d * j + k2s));

    //if ( (k1d * j + k1s) == (k2d * j + k2s) )
    if ( ((k1d - k2d) > 0) && ((k2s - k1s) >= 0) && ((k2s - k1s)%(k1d - k2d) == 0) )
    {
      //printf("They meet at count %i\n", j);
      return("YES");
    }
  }

  return("NO");
}

int main()
{
  int k1s, k2s, k1d, k2d;
  scanf("%i %i %i %i", &k1s, &k1d, &k2s, &k2d);

  char* result = kangaroo(k1s, k1d, k2s, k2d);

  printf("%s\n", result);
  return 0;
}

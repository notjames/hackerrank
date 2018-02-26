#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

const int upperbound = 10e5;
const int lowerbound = -10e5;

int checkBounds(int apples_size, int* apples, int oranges_size, int* oranges)
{
  int sane_bounds = 0;

  for ( int i = 0; i < oranges_size; i++ )
  {
    if ( oranges[i] >= upperbound )
    {
      fprintf(stderr, "The orange thrown %i is not within the upper boundary.\n", oranges[i]);
      sane_bounds++;
    }

    if ( oranges[i] <= lowerbound )
    {
      fprintf(stderr, "The orange thrown %i is not within the lower boundary.\n", oranges[i]);
      sane_bounds++;
    }
  }

  for ( int i = 0; i <= apples_size; i++ )
  {
    if ( apples[i] > upperbound )
    {
      fprintf(stderr, "The apple thrown %i is not within the upper boundary.\n", apples[i]);
      sane_bounds++;
    }

    if ( apples[i] <= lowerbound )
    {
      fprintf(stderr, "The apple thrown %i is not within the lower boundary.\n", apples[i]);
      sane_bounds++;
    }
  }

  if ( sane_bounds > 0 )
  {
    return 0;
  }

  return 1;
}

void countApplesAndOranges(int s, int t, int a, int b, int apples_size, int* apples, int oranges_size, int* oranges)
{
  if ( checkBounds(apples_size, apples, oranges_size, oranges) )
  {
    for ( int i = 0; i < apples_size; i++ )
    {
      printf("apples distance thrown: %i\n", apples[i]);
    }
  }
  else
  {
    printf("Not continuing...\n");
  }
}

int main()
{
    int s, t, a, b, m, n;

    scanf("%i %i", &s, &t);
    scanf("%i %i", &a, &b);
    scanf("%i %i", &m, &n);

    int *apples  = malloc(sizeof(int) * m);
    int *oranges = malloc(sizeof(int) * n);

    for (int apple_i = 0; apple_i < m; apple_i++)
    {
      scanf("%i",&apples[apple_i]);
    }

    for (int orange_i = 0; orange_i < n; orange_i++)
    {
      scanf("%i",&oranges[orange_i]);
    }

    countApplesAndOranges(s, t, a, b, m, apples, n, oranges);
    return 0;
}

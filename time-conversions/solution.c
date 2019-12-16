#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* readline();
char* timeConversion(char* s);

int split (char *str, char c, char ***arr)
{
  int count = 1;
  int token_len = 1;
  int i = 0;
  char *p;
  char *t;

  p = str;
  while (*p != '\0')
  {
    if (*p == c)
      count++;
    p++;
  }

  *arr = (char**) malloc(sizeof(char*) * count);
  if (*arr == NULL)
    exit(1);

  p = str;
  while (*p != '\0')
  {
    if (*p == c)
    {
      (*arr)[i] = (char*) malloc( sizeof(char) * token_len );
      if ((*arr)[i] == NULL)
        exit(1);

      token_len = 0;
      i++;
    }
    p++;
    token_len++;
  }
  (*arr)[i] = (char*) malloc( sizeof(char) * token_len );
  if ((*arr)[i] == NULL)
    exit(1);

  i = 0;
  p = str;
  t = ((*arr)[i]);
  while (*p != '\0')
  {
    if (*p != c && *p != '\0')
    {
      *t = *p;
      t++;
    }
    else
    {
      *t = '\0';
      i++;
      t = ((*arr)[i]);
    }
    p++;
  }

  return count;
}

char* timeConversion(char* s)
{
  char* ret;
  char seconds[2];
  char *meridian;
  int h;
  char** arr;

  arr      = NULL;
  split(s, ':', &arr);
  meridian = arr[2] + (strlen(arr[2]) - 2);
  h        = atoi(arr[0]);
  ret      = malloc(10 * sizeof(char));

  strncpy(seconds, arr[2], 2);

  //printf("meridian: %s\n", meridian);
  if ( strcmp(meridian, "PM") == 0 )
  {
    if ( h != 12 )
      h += 12;
  }
  else
  {
    if ( h == 12 )
    {
      h = 0;
    }
  }

  ( h < 10 ) ?  sprintf(ret, "0%i:%s:%s", h, arr[1], seconds) :
                sprintf(ret, "%i:%s:%s", h, arr[1], seconds);

  return ret;
}

int main()
{
  FILE* fptr;

  fptr         = fopen(getenv("OUTPUT_PATH"), "w");
  char* s      = readline();
  char* result = timeConversion(s);

  fprintf(fptr, "%s\n", result);

  return fclose(fptr);
}

char* readline()
{
  size_t alloc_length = 1024;
  size_t data_length  = 0;
  char* data          = malloc(alloc_length);

  while (true)
  {
    char* cursor = data + data_length;
    char* line   = fgets(cursor, alloc_length - data_length, stdin);

    if (!line) { break; }

    data_length += strlen(cursor);

    if (data_length < alloc_length - 1 || data[data_length - 1] == '\n') { break; }

    size_t new_length = alloc_length << 1;
    data = realloc(data, new_length);

    if (!data) { break; }

    alloc_length = new_length;
  }

  if (data[data_length - 1] == '\n')
  {
      data[data_length - 1] = '\0';
  }

  data = realloc(data, data_length);

  return data;
}


#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include"../include/function.h"
#define SIZE 50

char expression[10][20] = { 0 };
char op[10] = { 0 };
float numArr[SIZE] = { 0 };

int count = 0;

void checkForValidOprand(char *p)
{
	int c=0;
	char *temp=strtok(p,".");
	while(temp!=NULL)
	{
		temp=strtok(NULL,".");
		if(temp!=NULL){
			c++;
		}
	}
	if(c>=2)
	{
		printf("invalid expression");
		exit(0);
	}	
}
void
seperate_tok (char *expression)
{
  char p[20] = { 0 };
  int k = 0;
  int i = 0;
  int st = 0, end = 0;
  for (i = 0; i < strlen (expression); i++)
    {
      if (!ispunct (expression[i])||expression[i]=='.')
	{
	  end++;
	}
      else
	{
    if(st==end)
    {
      p[k] = '0';
    }
	  for (int j = st, k = 0; j < end; j++, k++)
	    {
	      p[k] = expression[j];
	    }
		checkForValidOprand(p);
	  numArr[count] = atof (p);
	  if (op[count - 1] == '*' || op[count - 1] == '/'||op[count - 1] == '%')
	    {
      if (op[count - 1] == '%')
	{
	  count--;
	  numArr[count] = modFunction (numArr[count], numArr[count + 1]);
	}
	      if (op[count - 1] == '*')
		{
		  count--;
		  numArr[count] =
		    multFunction (numArr[count], numArr[count + 1]);
		}
	      if (op[count - 1] == '/')
		{
		  count--;
		  numArr[count] =
		    divFunction (numArr[count], numArr[count + 1]);
		}
	    }
	  op[count] = expression[end];
	  end++;
	  st = end;
	  count++;
	  for (int j = 0; j < 20; j++)
	    {
	      p[j] = 0;
	    }
	}
    }
  for (int j = st, k = 0; j < end; j++, k++)
    {
      p[k] = expression[j];
    }
		checkForValidOprand(p);
  numArr[count] = atof (p);
  if (op[count - 1] == '*' || op[count - 1] == '/'||op[count - 1] == '%')
    {
      if (op[count - 1] == '*')
	{
	  count--;
	  numArr[count] = multFunction (numArr[count], numArr[count + 1]);
	}
      if (op[count - 1] == '%')
	{
	  count--;
	  numArr[count] = modFunction (numArr[count], numArr[count + 1]);
	}
      if (op[count - 1] == '/')
	{
	  count--;
	  numArr[count] = divFunction (numArr[count], numArr[count + 1]);
	}
    }
}

void
solve ()
{
  int st = 0;
  count;
  while (st != count)
    {
      if (op[st] == '+')
	{
	  numArr[st+1] = addFunction (numArr[st], numArr[st+1]);
	  st++;
	}
      if (op[st] == '-')
	{
	  numArr[st+1] = subFunction (numArr[st], numArr[st+1]);
	  st++;
	}

    }
  printf ("Ans: %0.4f", numArr[st]);
}

void
main ()
{
  char expression[50];
  printf ("Input the expression:");
  gets (expression);
  seperate_tok (expression);
  solve ();
}

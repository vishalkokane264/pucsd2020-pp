#include<stdio.h>
float divFunction(float op1,float op2)
{
    if(op2==0)
    {
        printf("can't divide by 0\n");
        return 0;
    }
    else
    {
        return (op1/op2);
    }
    
}
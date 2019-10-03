#include <cs50.h>
#include <stdio.h>
int main(void)
{
    int n = get_int("Height: ");
    while(n>8 || n<1)
    {
        n=get_int("Height: ");
    }
    int i,j;
    for (i=0;i<n;i++)
    {
       for (j=0;j<=n;j++)
       {
           if((i+j)<n)
           {
               printf(" ");
           }
           else
           {
               printf("#");
           }
       }
       printf("\n");
    }
}

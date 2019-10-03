#include <cs50.h>
#include <stdio.h>
int main(void)
{
    float n = get_float("Enter amount owed : ");
    int amt=n*100;
    int sum=0;
    int x,dolr=0,qtr=0,dm=0,nckl=0,pnn=0;
    dolr=amt/100;
    x=dolr*100;
    qtr=(amt-x)/25;
    dm=(amt-x-(qtr*25))/10;
    nckl=(amt-x-(qtr*25)-(dm*10))/5;
    pnn=(amt-(x*100)-(qtr*25)-(dm*10)-(nckl*5));
    if(dolr>0)
        sum=sum+dolr;
    if(qtr>0)
        sum=sum+qtr;
    if(dm>0)
        sum=sum+dm;
    if(nckl>0)
        sum=sum+nckl;
    if(pnn>0)
        sum=sum+pnn;
    printf("%d \n",sum);
}

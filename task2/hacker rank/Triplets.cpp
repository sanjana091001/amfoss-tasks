#include<iostream>  //comparing the triplets pgm
#include<conio.h>

using namespace std;

int main()
{
    int a[4],b[4],counta=0,countb=0;
    cout<<"  Enter marks of Alice  ";
    cin>>a[0]>>a[1]>>a[2];  //inputting scores of Alice
    cout<<"  Enter marks of Bob  ";
    cin>>b[0]>>b[1]>>b[2];  //inputting scores of Bob
    for(int i=0;i<3;i++)
    {
        if(a[i]>b[i])
        counta++;
        else if(a[i]<b[i])
        countb++;
    }
    cout<<"  Score of Alice and Bob are  ";
    cout<<counta<<" "<<countb;
    return 0;
    getch();
}

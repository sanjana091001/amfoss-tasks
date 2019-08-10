#include <iostream>
#include<conio.h>

using namespace std;

int main()
{
    int i,j,n;
    cout<<"enter the number of rows/columns in your matrix  "; //accepting values of rows/columns
    cin>>n;
    for(i=0;i<n+1;i++) 
    {
        cout<<endl;
        for(j=0;j<n+1;j++)
        {
            if((i+j)<n+1)
            cout<<" ";
            else
            cout<<"#";
            
        }
    }

    return 0;
    getch();
}

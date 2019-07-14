#include <iostream>  //program to find absolute diagonal difference of a matrix
#include <conio.h>

using namespace std;

int main()
{
    int arr[50][50],n,i,j,sum1=0,sum2=0,diagonaldifference;
    cout<<"Enter the number of rows/columns in your matrix :";
    cin>>n;
    for(i=0;i<n;i++)  //for loop to input values of matrix
    {
        for(j=0;j<n;j++)
        {
            cout<<"Enter the element at position:("<<i+1<<","<<j+1<<") ";
            cin>>arr[i][j];
        }
    }
    cout<<endl<<"the array you entered is ";
    for(i=0;i<n;i++)   //for loop to output matrix
    {
        cout<<endl;
        for(j=0;j<n;j++)
        {
            cout<<arr[i][j];
            cout<<" ";
        }
    }
    for(i=0;i<n;i++)  //loop for getting sum of primary diagnol 
    {
        for(j=0;j<n;j++)
        {
            if(i==j)
            sum1=sum1+arr[i][j];
        }
    }
    for(i=0;i<n;i++)  //loop for getting sum of secondary diagnol
    {
        for(j=0;j<n;j++)
        {
            if(i+j==(n-1))
            sum2=sum2+arr[i][j];
        }
    }
    if(sum1>sum2)
    diagonaldifference=sum1-sum2;
    else if(sum1<sum2)
    diagonaldifference=sum2-sum1;
    else
    diagonaldifference=0;
    cout<<endl<<"The diagonal difference is : "<<diagonaldifference;
    return 0;
    getch();
}

#include <iostream>

using namespace std;

int main()
{  
    int l,n[50],k,i,count=0;  
    cin>>l;  //l=total number of participants
    cin>>k;  //k=score above which participants will proceed to next round
    for(i=0;i<l;i++)  //loop accepting scores into array
        cin>>n[i];  
    for(i=0;i<l;i++)
        cout<<n[i]<<" ";
    for(i=0;i<l;i++)
    {
        if(n[i]>k)
        count=count+1;
    }
    cout<<endl<<count;

    return 0;
}

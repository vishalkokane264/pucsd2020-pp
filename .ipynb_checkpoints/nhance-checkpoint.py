def bestTimeToParty(celebTimes):
    st_time=[]
    end_time=[]
    temp_st=0
    maxelt=1

    for i in range(0,len(celebTimes)):
        st_time.append(int(celebTimes[i][0]))
        end_time.append(int(celebTimes[i][1]))
    for i in range(0,len(st_time)):
        for j in range(0,len(st_time)):
            if(st_time[i]<st_time[j]):
                temp_st=st_time[i]
                st_time[i]=st_time[j]
                st_time[j]=temp_st
                temp_st=end_time[i]
                end_time[i]=end_time[j]
                end_time[j]=temp_st
    count=0
    mintime=st_time[0]

    for i in range(0,len(st_time)):
        count=0
        for j in range(0,len(st_time)):
            if(st_time[i]==st_time[j]):
                count+=1
        if(count>maxelt):
            maxelt=count
            mintime=st_time[i]

    # Add your code here

    output = 'Best time to attend the party is at '+str(mintime)+' o\'clock : '+str(maxelt)+' celebrities will be attending!'
    return output


if __name__ == '__main__':
    celebTimes = []
    data=''
    data2=''
    no_of_cele=0
# Accept inputs
    no_of_cele=int(input())
    for i in range(0,no_of_cele):
        (data,data2)=input().split(',')
        celebTimes.append(data)
    print(celebTimes)
                
print(bestTimeToParty(celebTimes))

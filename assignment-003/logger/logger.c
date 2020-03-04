#include<stdio.h>
#include<string.h>
#include<time.h>
char filename[20];
FILE *logFile;
char myFile[10]="logger.c";
char* printMessages(int line, char *msg)
{
    time_t t;
    t=time(NULL);
    char *timestr = asctime( localtime(&t));
    timestr[strlen(timestr)-1]=0;
    fprintf(logFile,"%s\tFile:%s\tFunc:%s\tLine:%d\t%s",timestr,__FILE__,__func__,line,msg);
}

int getMessage()
{
    return rand()%10000;
}
void logMessages()
{
    FILE *sampleFile;
    int count=0,line=0;
    int i=0;
    char str[80];
    while(count!=10)
    {
        i=0;
        line=getMessage();
        sampleFile=fopen("example","r");
        while( fgets(str, 80, sampleFile) != NULL &&i!=line)
        {
            i++;
        }
        printMessages(line,str);
        count++;
        fclose(sampleFile);
    }
}
void openFile(char *filename)
{
    logFile=fopen(filename,"a+");
//    printf("File open");    
}
void closeFile()
{
    fclose(logFile);
}
int main(int argc, char const *argv[])
{
    if(argc==1)
    {
        printf("Give log file name.");
        fgets(filename,15,stdin);
        openFile(filename);
        logMessages();
        closeFile();
    }
    else
    {
        strcpy(filename,argv[1]);
        openFile(filename);
        printf("%s",filename);
        logMessages();
        closeFile();
    }
}

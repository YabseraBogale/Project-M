#include<iostream>
#include<filesystem>
#include "dirent.h"
int main(){
    std::string filename="Hello World.txt";
    for(auto &i: std::filesystem::directory_iterator(".")){
        if(i.is_directory()){
            DIR *k;
            k=opendir(i.path().c_str());
            while(k!=NULL){
                struct dirent *name;
                while((name=readdir(k))!=NULL){
                    std::cout<<name->d_name<<std::endl;
                }
                closedir(k);
            }
            
        }
    }
    return 0;
}
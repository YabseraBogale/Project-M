#include<iostream>
#include<filesystem>
#include "dirent.h"
int main(){
    std::string filename="Hello World.txt";
    for(auto &i: std::filesystem::directory_iterator(".")){
        if(i.is_directory()){
            //DIR *k;
            std::cout<<i.path()<<std::endl;
            //k=opendir();
            
        }
    }
    return 0;
}
#include<iostream>
#include<filesystem>
int main(){
    std::string filename="Readme.md";
    //std::getline(std::cin,filename);
    bool exist=std::filesystem::exists(filename);
    if(exist==true){
        
    } else{
        for(auto &dir: std::filesystem::directory_iterator("..")){
                std::cout<<dir.path()<<std::endl;
        }
    }
    return 0;
}
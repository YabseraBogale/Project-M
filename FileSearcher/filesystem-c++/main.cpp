#include<iostream>
#include<filesystem>
#include<fstream>
int main(){
    std::string filename="Readme.md";
    //std::getline(std::cin,filename);
    bool exist=std::filesystem::exists(filename);
    if(exist==true){
        std::cout<<std::filesystem::current_path()<<std::endl;
    } else{
        
    }
    return 0;
}
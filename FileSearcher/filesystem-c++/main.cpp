#include<iostream>
#include<filesystem>
int main(){
    std::string filename="Hello World.txt";
    for(auto &i: std::filesystem::directory_iterator(".")){
        std::cout<<i.is_directory()<<std::endl;
    }
    return 0;
}
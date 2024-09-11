#include<iostream>
#include<filesystem>

int main(){
    auto pwd=std::filesystem::current_path();     
    std::cout<<pwd<<std::endl;
    return 0;
}


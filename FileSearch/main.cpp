#include<iostream>
#include<array>
int main(){
    auto p=popen("locate ~/*.txt","r");
    std::array<char, 128> buffer;
    std::string result;
    if(!p){
        throw std::runtime_error("falied");
    } else {
        while(!feof(p)){
            if(fgets(buffer.data(),buffer.size(),p)!=nullptr){
                result+=buffer.data();
            }
        }
        auto end=pclose(p);
        if(end==EXIT_SUCCESS){

            std::cout<<"ok\n"<<result;
        }
        else{
            std::cout<<"not ok\n"<<result<<std::endl;
        }
    } 
    return 0;
}
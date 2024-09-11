x86_64-w64-mingw32-g++ main.cpp -static-libstdc++ -static-libgcc -o main && wine ./main.exe
git add . && git commit -m "ok" && git push
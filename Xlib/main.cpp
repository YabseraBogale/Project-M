#include<iostream>
#include <stdlib.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xos.h>

Display *display;
int main()
{
  display = XOpenDisplay(NULL);
  if (display == NULL) {
    std::cout<< "Failure connecting to X Server";
    std::exit(0);
  }
  return 0;
}
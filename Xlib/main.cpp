#include <iostream>
#include <cstdlib>
#include <cstring>
 
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/XKBlib.h>
#include <X11/Xatom.h>
#include <X11/keysym.h>
 
int main()
{
    XVisualInfo* visual = 0;
    Window window = 0;   
    XSetWindowAttributes attributes;
     
    memset(&attributes, 0, sizeof(attributes));
     
    Display* display;
    int screennr;
     
    display = XOpenDisplay(0);
    if (!display)
        return -1;
     
    screennr = DefaultScreen(display);
     
    XVisualInfo visTemp;
    int visNumber;
     
    visTemp.screen = screennr;
    visTemp.depth = 24;
    while (!visual && visTemp.depth >= 16)
    {
        visual = XGetVisualInfo(display, VisualScreenMask | VisualDepthMask, &visTemp, &visNumber);
        visTemp.depth -= 8;
    }
     
    Colormap colormap;
    colormap = XCreateColormap(display,
            RootWindow(display, visual->screen),
            visual->visual, AllocNone);
 
    attributes.colormap = colormap;
    attributes.border_pixel = 0;
    attributes.event_mask = StructureNotifyMask | FocusChangeMask | ExposureMask;
    attributes.event_mask |= PointerMotionMask |
                ButtonPressMask | KeyPressMask |
                ButtonReleaseMask | KeyReleaseMask;
    attributes.override_redirect = 0;
     
    window = XCreateWindow(display, RootWindow(display, visual->screen), 
                           0, 0, 800, 600, 0, visual->depth, 
                           InputOutput, visual->visual,
                           CWBorderPixel | CWColormap | CWEventMask | CWOverrideRedirect,
                           &attributes);
    XMapRaised(display, window);
     
    for (;;)
    {
        XEvent event;
        XNextEvent(display, &event);
             
        switch (event.type)
        {
            case KeyPress:
            {
                long unsigned int X11Key;
                char buf[8]={0};
                XLookupString(&event.xkey, buf, sizeof(buf), &X11Key, NULL);
     
                if (X11Key == XK_w)
                    std::cout << "Button Pressed!\n";
            }
            break;
            case KeyRelease:
            {
                long unsigned int X11Key;
                char buf[8]={0};
                XLookupString(&event.xkey, buf, sizeof(buf), &X11Key, NULL);
     
                if (X11Key == XK_w)
                    std::cout << "Button Released!\n";
            }
            break;
        }
    }
     
    return 0;
}
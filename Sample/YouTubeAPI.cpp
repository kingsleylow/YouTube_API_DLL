#include <iostream>
#include <windows.h>
#include "YouTubeAPI.h"

using namespace std;

typedef GoInt32(*getSubDigit)(void);
typedef GoInt32(*getViewsDigit)(void);
typedef GoInt32(*getAvgDigit)(void);

int main() 
{
    char s;
    HINSTANCE API = LoadLibrary("YouTubeAPI.dll");
    getSubDigit subNo = (getSubDigit)GetProcAddress(API, "getSubscriber");
    if (subNo)
    {
        GoInt32 subscribe = subNo();
        cout << "I have " << subscribe << " subscriber" << endl;
    }
    getViewsDigit viewNo = (getViewsDigit)GetProcAddress(API, "getViews");
    if (viewNo)
    {
        GoInt32 views = viewNo();
        cout << "My videos' have viewed " << views << " times" << endl;
    }
    getAvgDigit avg = (getAvgDigit)GetProcAddress(API, "getAvgViewsVideos");
    if (avg)
    {
        GoInt32 vpv = avg();
        cout << "Each videos have viewed " << avg << " time" << s << " on average" << endl;
    }
    cout << "My YouTube channel: rk0_d" << endl;
    cout << "Website: https://rkocyrus.github.io" << endl;
    return 0;
}

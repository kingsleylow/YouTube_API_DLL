# YouTube_API_DLL
 Get YouTube Data API via DLL

(By rk0_d a.k.a rkoCyrus)

Before use
=======
You **MUST** have your channel id and key from YouTube Data API that you can use it
When you found your channel id and the key, please modify line 40 on YouTubeAPI.go:
`yourYouTubeApiSite string = "https://www.googleapis.com/youtube/v3/channels?part=statistics&id=(Your channel ID)&key=(Your YouTube Data API key)"`
In additions, you need to install them before build:  
    1. Mingw-w64 (MSVC maybe compliable)  
    2. go (at lease 1.13.6 or above)  

Build
=======
After you modified your API site, you can run `build.bat`.  
When `build.bat` was executed, it will generate:  
    1. .h (header)  
    2. .dll (dll library)  
    3. .def  
    4. .a  
    5. lib (library, mostly use in visual studio)  
And you can use dll to build what you want!

Function
=======
It will get  
    * No. of subscriber  
    * Total views  
    * Views per videos  
You can add another function in `YouTubeAPI.go` if needed

Demo
=======
You will find a sample console program of my channel's status which can run normally.
To execute it, please open in cmd and type `YouTubeAPI.exe` on it's current path instead of just double-click it since **IT DOES'T HAVE SLEEP FUNCTION**

Lincense
=======
**NO (HOWEVER, DO NOT USE IT AS MAKING PROFIT!)**

Profile
=======
Website: https://rkocyrus.github.io  
Youtube: https://www.youtube.com/user/long1212cn  
Facebook: https://www.facebook.com/rk0dd/  
Twitter: https://twitter.com/rk0cc  
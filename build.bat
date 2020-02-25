@echo off
rem this is building dll script
rem DO NOT MODIFY IF NOT NECESSARY
if exist .\YouTubeAPI.go (
    call go build -buildmode=c-archive .\YouTubeAPI.go
    @echo EXPORTS> YouTubeAPI.def
    @echo ^ getSubscriber>> YouTubeAPI.def
    @echo ^ getViews>> YouTubeAPI.def
    @echo ^ getAvgViewsVideos>> YouTubeAPI.def
    call gcc YouTubeAPI.def YouTubeAPI.a -shared -lwinmm -lWs2_32 -o YouTubeAPI.dll -Wl,--out-implib,YouTubeAPI.lib
    echo "Finished"
    timeout /t 5
) else (
    echo "Error"
    timeout /t 5
)
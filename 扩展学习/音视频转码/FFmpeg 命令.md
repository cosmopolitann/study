# FFmpeg 命令

#### 1.视频转码

```go
ffmpeg -i E:/sample1.mp4 -threads 2 -strict -2 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 4 -hls_list_size 0 -f hls E:/ok/sampele.m3u8
```



#### 2.推流 rtsp

```go
ffmpeg -re -stream_loop -1 -i C:\Users\Administrator\Desktop\samp;e2.mp4 -c copy -rtsp_transport tcp  -f rtsp rtsp://localhost:8554/mystream

//C:\Users\Administrator\Desktop

ffmpeg -re -stream_loop -1 -i file.ts -c copy -f rtsp rtsp://localhost:8554/mystream

//
C:\Users\Administrator\Desktop>ffmpeg -re -stream_loop -1 -i C:/Users/Administrator/Desktop/"sample 2.mp4" -c copy -rtsp_transport tcp  -f rtsp rtsp://localhost
:8554/mystream

```



#### 3.拉流 rtsp



```go
ffmpeg -i rtsp://localhost:8554/mystrea
c copy -rtsp_transport tcp -f segment -segment_time 30 output_%03d.mp4

//C:\Users\Administrator\Desktop\1

ffmpeg -i rtsp://localhost:8554/mystream -c copy -rtsp_transport tcp -f segment -segment_time 5 C:\Users\Administrator\Desktop\1\output_%03d.mp4

```

```
ffmpeg -i rtsp://localhost:8554/mystream -c copy output.mp4
```



#### 4.启动

```go
rtsp-simple-server.exe --protocols="udp,tcp"
```


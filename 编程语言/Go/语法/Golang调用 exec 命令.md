### Golang 调用 exec 命令

##### 1.Linux

```go
package main
   
import (
    "bytes"
    "fmt"
    "os/exec"
)
   
func main() {
    in := bytes.NewBuffer(nil)
    cmd := exec.Command("sh")
    cmd.Stdin = in
    go func() {
        in.WriteString("echo hello world > test.txt\n")
        in.WriteString("exit\n")
    }()
    if err := cmd.Run(); err != nil {
        fmt.Println(err)
        return
    }
}

```

#### 2.windows

```go
func main() {

	cmd := exec.Command("sh")
	// cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	s.Add(1)
	go func() {
		// start stop restart
		
		in.WriteString("mkdir -p E:\\pro\\sss1111\n")
		fmt.Println("执行第二个")
		time.Sleep(time.Second * 10)
		in.WriteString("mkdir -p E:\\pro\\sss2222\n")

		//in.WriteString("ffmpeg -i rtsp://localhost:8554/mystream -c copy -rtsp_transport tcp -f segment -segment_time 10 C:/Users/Administrator/Desktop/1/output_%d.mp4\n") //写入你的命令，可以有多行，"\n"表示回车
		s.Done()
	}()

	s.Wait()
	//ffmpeg -i rtsp://localhost:8554/mystream -c copy -rtsp_transport tcp -f segment -segment_time 10 C:\Users\Administrator\Desktop\1\output_%d.mp4
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	rt := out.String() //mahonia.NewDecoder("gbk").ConvertString(out.String()) //
	fmt.Println(rt)

}

```


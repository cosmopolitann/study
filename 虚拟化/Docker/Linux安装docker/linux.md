## Linux安装docker

### Ubuntu系统

#### 1.使用官方安装脚本自动安装

安装命令如下

```
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

#### 2.手动安装

也可以使用国内 daocloud 一键安装命令：

```
curl -sSL https://get.daocloud.io/docker | sh
```

安装流程

```
server@server-gpu-97:~$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
# Executing docker install script, commit: 7cae5f8b0decc17d6571f9f52eb840fbc13b2737
+ sudo -E sh -c 'apt-get update -qq >/dev/null'
[sudo] password for server:
+ sudo -E sh -c 'DEBIAN_FRONTEND=noninteractive apt-get install -y -qq apt-transport-https ca-certificates curl >/dev/null'
+ sudo -E sh -c 'curl -fsSL "https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg" | apt-key add -qq - >/dev/null'
Warning: apt-key output should not be parsed (stdout is not a terminal)
+ sudo -E sh -c 'echo "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu focal stable" > /etc/apt/sources.list.d/docker.list'
+ sudo -E sh -c 'apt-get update -qq >/dev/null'
+ '[' -n '' ']'
+ sudo -E sh -c 'apt-get install -y -qq --no-install-recommends docker-ce >/dev/null'
+ '[' -n 1 ']'
+ sudo -E sh -c 'DEBIAN_FRONTEND=noninteractive apt-get install -y -qq docker-ce-rootless-extras >/dev/null'
+ sudo -E sh -c 'docker version'
Client: Docker Engine - Community
 Version:           20.10.6
 API version:       1.41
 Go version:        go1.13.15
 Git commit:        370c289
 Built:             Fri Apr  9 22:47:17 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.6
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.15
  Git commit:       8728dd2
  Built:            Fri Apr  9 22:45:28 2021
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.4.4
  GitCommit:        05f951a3781f4f2c1911b05e61c160e9c30eaa8e
 runc:
  Version:          1.0.0-rc93
  GitCommit:        12644e614e25b05da6fd08a38ffa0cfe1903fdec
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0

================================================================================

To run Docker as a non-privileged user, consider setting up the
Docker daemon in rootless mode for your user:

    dockerd-rootless-setuptool.sh install

Visit https://docs.docker.com/go/rootless/ to learn about rootless mode.


To run the Docker daemon as a fully privileged service, but granting non-root
users access, refer to https://docs.docker.com/go/daemon-access/

WARNING: Access to the remote API on a privileged Docker daemon is equivalent
         to root access on the host. Refer to the 'Docker daemon attack surface'
         documentation for details: https://docs.docker.com/go/attack-surface/

================================================================================
```

查看docker 版本

```
server@server-gpu-97:~$ docker --version
Docker version 20.10.6, build 370c289
server@server-gpu-97:~$
```



### 安装docker-compose

```go
sudo apt install docker-compose
```


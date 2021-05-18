p2p 聊天  应用  libp2p
p2p chat app with libp2p

这个程序演示了一个简单的p2p聊天应用程序。它可以在两个对等的网络之中。
This program demonstrates a simple p2p chat application. It can work between two peers if



两者都有一个私有IP地址(相同的网络)。
Both have a private IP address (same network).

其中至少有一个有公共IP地址。
At least one of them has a public IP address.


假设“A”和“B”在不同的网络上，主机“A”可能有也可能没有公共IP地址，但主机“B”有。
Assume if 'A' and 'B' are on different networks host 'A' may or may not have a public IP address but host 'B' has one.

Usage: Run ./chat -sp <SOURCE_PORT> on host 'B' where <SOURCE_PORT> can be any port number. Now run ./chat -d <MULTIADDR_B> on node 'A' [<MULTIADDR_B> is multiaddress of host 'B' which can be obtained from host 'B' console].

构建这个聊天软件
从这个目录运行一下命令


Build
From the go-libp2p/examples directory run the following:

> cd chat/
> go build
Usage
On node 'B'

> ./chat -sp 3001
Run ./chat -d /ip4/127.0.0.1/tcp/3001/p2p/QmdXGaeGiVA745XorV1jr11RHxB9z4fqykm6xCUPX1aTJo

2018/02/27 01:21:32 Got a new stream!
> hi (received messages in green colour)
> hello (sent messages in white colour)
> no
On node 'A'. Replace 127.0.0.1 with <PUBLIC_IP> if node 'B' has one.

> ./chat -d /ip4/127.0.0.1/tcp/3001/p2p/QmdXGaeGiVA745XorV1jr11RHxB9z4fqykm6xCUPX1aTJo
Run ./chat -d /ip4/127.0.0.1/tcp/3001/p2p/QmdXGaeGiVA745XorV1jr11RHxB9z4fqykm6xCUPX1aTJo

This node's multiaddress:
/ip4/0.0.0.0/tcp/0/p2p/QmWVx9NwsgaVWMRHNCpesq1WQAw2T3JurjGDNeVNWifPS7
> hi
> hello
NOTE: debug mode is enabled by default, debug mode will always generate the same node id (on each node) on every execution. Disable debug using --debug false flag while running your executable.


如果您正在寻找具有对等发现的实现，则使用会合点支持对等发现。
Note: If you are looking for an implementation with peer discovery, chat-with-rendezvous, supports peer discovery using a rendezvous point.

Authors
Abhishek Upperwal
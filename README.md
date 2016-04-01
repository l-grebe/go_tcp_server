## go_tcp_server

net包中文API：[http://studygolang.com/static/pkgdoc/pkg/net.htm#TCPAddr](http://studygolang.com/static/pkgdoc/pkg/net.htm#TCPAddr)

#### type TCPAddr
    type TCPAddr struct {
        IP   IP
        Port int
        Zone string // IPv6范围寻址域
    }
* TCPAddr代表一个TCP终端地址。

#### func ResolveTCPAddr
    func ResolveTCPAddr(net, addr string) (*TCPAddr, error)

* ResolveTCPAddr将addr作为TCP地址解析并返回。
* 参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；
* net必须是"tcp"、"tcp4"或"tcp6"。
* IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

####  func ListenTCP
    func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)

* ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener。
* net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。

#### func (*TCPListener) Accept
    func (l *TCPListener) Accept() (Conn, error)
* Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。

#### func (*IPConn) RemoteAddr
    func (c *IPConn) RemoteAddr() Addr
* RemoteAddr返回远端网络地址

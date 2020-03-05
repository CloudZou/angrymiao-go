# GRPC/RPC 注册中心使用说明

. grpc服务实现代码在github.com/CloudZou/punk/pkg/net/rpc/warden包，相应的服务发现客户端实现在/net/rpc/warden/resolver 
. rpc服务实现代码在github.com/CloudZou/punk/pkg/net/rpc,服务发现实现在/net/rpc/client2.go文件,NewDiscoveryCli函数
. 服务注册中心地址为discovery.angrymiao.com,目前单节点部署在k8s集群上

## 项目简介
+ rpc 服务代码参考:
*rpc方法里面的第一个context.Context 必须是 "github.com/CloudZou/punk/pkg/net/rpc/context"这个路径下的，*
*另外第二个参数和第三个参数必须是指针，分表代表入参和返回参数*
```cassandraql
// New new rpc server.
func New(s *service.Service) (svr *rpc.Server) {
	r := &RPC{s: s}
	svr = rpc.NewServer(conf.Conf.RPCServer)
	if err := svr.Register(r); err != nil {
		panic(err)
	}
	return
}

func (r *RPC) SayHello(c context.Context, arg *string, res *struct{}) (err error) {
	log.Error("Say Hello From RPC Client")
	_, err = r.s.SayHello(c, &api.HelloReq{Name: *arg})
	return
}

func (r *RPC) SayHelloURL(c context.Context, name *string, res *struct{}) (err error) {
	_, err = r.s.SayHelloURL(c, &api.HelloReq{Name: *name})
	return
}
```
+ rpc 客户端代码调用参考，使用rpc.NewDiscoveryCli(_appid,c)来获取远程服务实例，不需要指定ip端口等信息,**(需要指定方法，是按照RPC+服务端接口方法名来匹配的，如RPC.Ping,Ping是指服务端方法的签名,RPC是固定的）**:
```
type Service struct {
	client     *rpc.Client2
}

// New new rpc service.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	s.client = rpc.NewDiscoveryCli(_appid, c)
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	err = s.client.Call(ctx, "RPC.Ping", nil, nil)
	return
}

func (s *Service) SayHello(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHello", name, nil)
	return
}

func (s *Service) SayHelloURL(ctx context.Context, name string) (err error) {
	err = s.client.Call(ctx, "RPC.SayHelloURL", name, nil)
	return
}

```

+ grpc服务代码参考:
```cassandraql
// New new a grpc rpc.
func New(s *service.Service) (ws *warden.Server, err error) {
	ws = warden.NewServer(conf.Conf.GRPCServer)
	pb.RegisterDemoServer(ws.Server(), &rpc{as: s})
	if ws, err = ws.Start(); err != nil {
		return
	}
	return
}

type rpc struct {
	as *service.Service
}

var _ pb.DemoServer = &rpc{}

func (s *rpc) Ping(ctx context.Context, req *empty.Empty) (reply *empty.Empty, err error) {
	reply, err = s.as.Ping(ctx, req)
	return nil, nil
}

func (s *rpc) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply, err = s.as.SayHello(ctx, req)
	return
}

func (s *rpc) SayHelloURL(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	resp, err = s.as.SayHelloURL(ctx, req)
	return
}

```

+ grpc客户端代码调用参考, 直接使用v1.NewClient来从discovery获取远程实例，其中grpc的服务ip端口获取封装在github.com/CloudZou/punk/pkg/net/rpc/warden/resolver包:
```cassandraql
type Service struct {
	demoClient v1.DemoClient
}


// New new rpc service.
func New(c *rpc.ClientConfig) (s *Service) {
	s = &Service{}
	//GRPC调用模式，从服务注册中discovery.angrymiao.com找到grpc server 地址
	s.demoClient, err = v1.NewClient(nil)
	if err != nil {
		panic(err)
	}
	return
}

func (s *Service) Ping(ctx context.Context) (err error) {
	_, err = s.demoClient.Ping(ctx, nil)
	return
}

func (s *Service) SayHello(ctx context.Context, name string) (err error) {
	_, err = s.demoClient.SayHello(ctx, &v1.HelloReq{Name:name})
	return
}

func (s *Service) SayHelloURL(ctx context.Context, name string) (err error) {
	_, err = s.demoClient.SayHelloURL(ctx, &v1.HelloReq{Name:name})
	return
}

```
*必须使用pb协议，实现DemoServer的相应方法,注意此时的context.Context是系统路径下的context包，跟使用rpc方式的context包路径不一致，服务注册方式同后面的服务注册逻辑一致*


*注册完成后必须调用discovery的Register把服务注册给服务注册中心discovery,主要是包含区域，zone，服务唯一标识，以及包括http,grpc,rpc服务的ip端口等信息.见如下代码*
```cassandraql
env.Region = "region01"
env.Zone = "zone01"
dis := discovery.New(nil)
ins := &naming.Instance{
    AppID:    "user.service",
    Hostname: hn,
    Addrs: []string{
        "http://" + ip + ":" + env.HTTPPort,
        "gorpc://" + ip + ":" + env.GORPCPort,
        "grpc://" + ip + ":" + env.GRPCPort,
        "http://test.angrymiao.com",
    },
}
if cancel, err = dis.Register(context.Background(), ins); err != nil {
    panic(err)
}
```
***当然这部分逻辑可以放在k8s启动docker的时候由k8s来完成具体的注册信息，相应的ip端口可以写在k8s yaml脚本里面***
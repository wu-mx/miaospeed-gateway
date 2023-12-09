# Miaospeed-Gateway

一个Miaospeed中间件，通过修改请求以支持通过Miaoko官方主端连接到修改过预设变量的Miaospeed。

通常可以和Miaoko部署在同一服务器。Miaoko产品的超级管理员可以通过Gateway连接到自构建的增强安全性的Miaospeed，以防止他人使用官方弱安全性的Miaospeed BuildToken来伪造请求数据，同时可以用于实现一些神奇功能。

## 配置
所有配置均存放在config.yaml内，可在readme.config.yaml内找到示例。

如选择自定义文件，请在启动时使用`-c`参数指定配置文件路径。

```yaml
slaves: # 所有测试点列表，你需要从Miaoko主端拷贝Address,Token过来用于验证和连接。
  GuangZhouCM: # 按照实例从主端拷贝需要的字段即可
    address: ws://127.0.0.1:9876
    token: EXECTOKEN
    # 上面两个字段必填，后端的英文名须和主端配置一致，Miaospeed Gateway通过主端发起的后端请求中
    # 的英文名称来查找对应的后端，从而将请求转发至后端
    disable: false # 禁用后端，禁用后主端发起请求会拒绝连接
    buildToken: MSGATE0|114514|geNeral|1X571R930|T0kEN # 自定义后端的Build Token，构建专属后端有利于防止其他后端托管者伪造结果
    tlsPubKey: "./pub.crt" # 自定义后端公钥，该选项在使用mwss协议时无效
    skipTokenVerify: false # 跳过Gateway对主端的执行Token验证
    skipTLSVerify: false # 跳过Gateway对后端的TLS证书验证，某些自签证书可能需要开启
    invoker : 1234567890 # 神奇设置，不怕被揍请自动忽略本行
serverTLS: true # 开启Gateway服务的TLS,效果同miaospeed中的 -mtls
listen: :8080 # 服务监听端口，可监听unix socket
whiteList: # 白名单，只有在白名单中的主端才能连接Gateway，留空或不填则为关闭
  - 93372553
  - 19198100

```
## 使用
将Gateway启动完成后，在Miaoko主端的配置中，将绑定到Gateway的后端的adrress字段统一改为`mwss://<Gateway地址>:<Gateway端口>`即可。
> 所有绑定的后端在主端的地址须一致，英文名称、ExecToken也须与主端一致

 ## 构建专属Miaospeed后端
基本编译流程可在[MiaoSpeed官方文档](https://github.com/miaokobot/miaospeed)中找到，但是直接编译出的后端安全性极低(包括官方构建)。构建密钥，TLS证书等都以明文形式存放在后端中，任何具有基础反编译知识的人都可以读取这些字面量。同时Miaoko官方也早在前段时间在其官方频道公布了构建密钥(由于频道已设为私有本仓库暂不公开该密钥)，这使得任何人都可以构建出一个伪造数据的Miaospeed后端。

我们提供两种方式以让您使用自定义内设变量。

1. 使用[Garble](https://github.com/burrowers/garble)构建自定义Miaospeed，来加强自定义后端的安全性。
2. 使用[Miaospeed Community Build](https://github.com/Paimonhub/miaospeed_community/)，可在预先构建好的二进制文件中直接设置有关信息。

如果您将Miaospeed后端提供给他人部署，我们推荐使用Garble，Garble会直接将字面量写死在程序并加强混淆。

### 使用Garble构建Miaospeed
1. 安装Go环境(>=1.20，使用Debian系包管理器安装的不要使用apt，使用snap，apt的go过于老旧)，clone Miaospeed仓库(可使用[官方版](https://github.com/miaokobot/miaospeed)或[墨少离版](https://github.com/miaokobot/miaospeed))，需要构建Meta版本的请将go.mod中最后一行的注释去掉。
2. 安装Garble. `go install mvdan.cc/garble@latest`
3. 使用Garble构建Miaospeed
   ```bash
   garble build -literals -tiny -seed 114514 -ldflags "-s -w -X \"main.COMMIT=0a0089d5d78a171ef6defc1d17adf23f55e6c680\" -X \"main.BUILDCOUNT=33\" -X \"main.BRAND=MiaoSpeed\" -X \"main.COMPILATIONTIME=1678112713\""
   ```
   seed可自定义，commit,buildcount,等可自行修改。GOOS，GOARCH等编译目标设置请自行谷歌。
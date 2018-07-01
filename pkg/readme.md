
## pkg
```
go env

//set GOPATH=D:\Go\gopath
```

> 自定义包放在$GOROOT/$GOPATH下的src目录下
> :star: GOPATH 可设置多个 

* 安装自定义包
```
// export GOPATH=$home/pkg
set GOPATH=$home/pkg

go build <pkg name>
go install <pkg name>
```

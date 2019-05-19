
## pkg
```
go env

//set GOPATH=D:\GoSpace
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


* add $GOPATH/bin to $PATH
> [How to Write Go Code - The Go Programming Language](https://golang.org/doc/code.html#GOPATH)
```
vim ~/.bash_profile
// add a line
export PATH=$PATH:$(go env GOPATH)/bin
source  ~/.bash_profile
```

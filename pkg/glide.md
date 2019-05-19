
## Glide

## golang.org/x/net
> $(go env GOPATH)

```
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git

git clone https://github.com/golang/tools.git 
git clone https://github.com/golang/lint.git 

go get golang.org/x/lint/golint 

// Shift+Alt+F
gofmt -w .
```

## glide mirror

```
glide get golang.org/x/sys/unix
```

## dep
> [Installation · dep](https://golang.github.io/dep/docs/installation.html)

```
dep init
dep ensure -add github.com/foo/bar 

```

## Refer

1. [glide/README_CN.md xkeyideal/glide](https://github.com/xkeyideal/glide/blob/master/README_CN.md)

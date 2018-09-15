
## Go Tutorial

## wiki
1. [LearnServerProgramming · golang/go Wiki](https://github.com/golang/go/wiki/LearnServerProgramming)
2. [Tinywan/golang-tutorial: Golang 系列教程(译)、学习笔记](https://github.com/Tinywan/golang-tutorial)
3. [Unknwon/go-rock-libraries-showcases](https://github.com/Unknwon/go-rock-libraries-showcases)

## :tada:docker on Windows | tips
* First, list all machines: 
`docker-machine ls`
* If you don't have one, you can create one with: 
`docker-machine create default`

```
docker-machine create --driver virtualbox --engine-registry-mirror https://1rqn7spi.mirror.aliyuncs.com test
```
* Start the machine, if stopped: 
`docker-machine start default`
* Connect with machine via SSH: 
`docker-machine ssh default`

```
ssh docker@192.168.99.100
docker/tcuser
sudo -i
```

* install docker-compose in `default`
[Install Docker Compose | Docker Documentation](https://docs.docker.com/v17.09/compose/install/#install-compose)

:warning: VirtualBox Host-Only Ethernet Adapter 


```
sudo sed -i "s|EXTRA_ARGS='|EXTRA_ARGS='--registry-mirror=http://9zkjjecg.mirror.aliyuncs.com/ |g" /var/lib/boot2docker/profile 
exit 
docker-machine restart default
```

```
yum -y install epel-release
yum -y install htop
```

## Refer
1. [Go Tutorial](https://www.tutorialspoint.com/go/index.htm)
2. [Go by Example](https://gobyexample.com/)
3. [Unknwon / Repositories](https://github.com/Unknwon?tab=repositories)
4. [gogs/gogs: Gogs is a painless self-hosted Git service.](https://github.com/gogs/gogs)
5. [go-macaron/macaron: Package macaron is a high productive and modular web framework in Go.](https://github.com/go-macaron/macaron)
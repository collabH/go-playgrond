# go-playgroup
go Quick Start

# build env
## go lib install
```shell
# brew查询go是否存在
brew search go 
# 安装go
brew install go 
# 配置go环境变量 
vim ~/.bash
 vim ~/.zshrc 
# GOROOT
export GOROOT=/opt/homebrew/Cellar/go/1.21.0/libexec
#GOPATH root bin
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
#GOPATH
export GOPATH=$HOME/go
#GOPATH bin
export PATH=$PATH:$GOPATH/bin
```

## go依赖管理工具安装

```shell
# 启动go默认版本管理工具
export GO111MODULE=on
# 设置goproxy
export GOPROXY=https://goproxy.cn
```
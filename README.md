# ihash
命令行，带进度，支持通配符的hash校验工具
# 快速安装
go get -u github.com/xiaoqidun/ihash
# 编译安装
```
git clone https://github.com/xiaoqidun/ihash.git
cd ihash
go build ihash.go
```
# 校验类型
- md5sum
- sha1sum
- sha256sum
- sha512sum
- sha3sum224
- sha3sum256
- sha3sum384
- sha3sum512
# 快捷命令
自行将./bin/替换成处于PATH环境变量中的目录路径
```
ihash -install ./bin/
```
# 字符校验
```
echo admin | md5sum
echo admin | ihash md5sum
```
# 文件校验
```
md5sum *
ihash md5sum *
```
# 传参说明
```
快捷命令 文件1 文件2 文件3 ...
原始程序 校验类型 文件1 文件2 文件3 ...
不传任何文件时从stdin（命令行）读取字符
```
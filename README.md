# ihash
命令行，全平台，支持通配符的hash校验工具
# 快速安装
go get -u github.com/xiaoqidun/ihash
# 编译安装
```
git clone https://github.com/xiaoqidun/ihash.git
cd ihash
go build -mod=vendor ihash.go
```
# 手动安装
1. 根据系统架构下载为你编译好的[二进制文件](https://github.com/xiaoqidun/ihash/releases)
2. 将下载好的二进制文件重命名为ihash并保留后缀
3. 把ihash文件移动到系统PATH环境变量中的目录下
4. windows外的系统需使用chmod命令赋予可执行权限
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

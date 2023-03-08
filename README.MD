# LSQ CLI TOOL


`lsq` is a powerful CLI tool for Do Easy Something.

## 1. 安装

### 1). 手动安装

> You might need setting the goproxy to make through building.
> Please make sure your Golang version > v1.15.

1. 最新版本安装
    ```
    go install github.com/iAsuma/lsq-cli/lsq@master
    ```

2. 指定版本安装
    ```
    go install github.com/iAsuma/lsq-cli/lsq@v1.0.2
    ```

3. Check installation
   ```
   lsq -v
   ```
### 2). PreBuilt Binary

    TODO

## 2. Commands
```html
$ lsq
USAGE
lsq COMMAND [OPTION]

COMMAND
gen        automatically generate GoFrameV2 files for api/controller/logic/model
version    show version information of current binary

OPTION
-v, --version   show version information of current binary
-h, --help      more information about this command

ADDITIONAL
Use "lsq COMMAND -h" for details about a command.

```

## 3. 生成标准化开发文件
查看命令说明
```
lsq gen mc --help(-h)
```
goframe 2.0版本
```
lsq gen mc -n xxx --last
```
goframe 2.1及以上版本
```
lsq gen mc -n xxx
```

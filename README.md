[![Build Status](https://travis-ci.org/Jockey66666/fx2preset-lite.svg?branch=master)](https://travis-ci.org/Jockey66666/fx2preset-lite)

作者：饶曉文
链接：https://www.jianshu.com/p/155b3f18abae
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
# fx2preset-lite
## 簡介
內部使用工具，使用前要先安裝[go](https://golang.org/) 1.14以後的版本

## 安裝
```
go get -v -u github.com/Jockey66666/fx2preset-lite
```

# 功能介紹
## 還原文件
```
fx2preset-lite restore
```

## 列出所有preset
```
fx2preset-lite list
```

## 打開preset目錄
```
fx2preset-lite find "American Dream"
```

## 登出tonecloud
```
fx2preset-lite logout
```

## 切換tonecloud環境
### prodction
```
fx2preset-lite env prod
```

### staging
```
fx2preset-lite env stage
```

[![Travis CI](https://travis-ci.org/Jockey66666/fx2preset-lite.svg?branch=master)](https://travis-ci.org/Jockey66666/fx2preset-lite)
![Github Actions](https://github.com/Jockey66666/fx2preset-lite/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jockey66666/fx2preset-lite)](https://goreportcard.com/report/github.com/jockey66666/fx2preset-lite)


# fx2preset-lite
## 簡介
內部使用工具，使用前要先安裝 [go](https://golang.org/) 1.14以後的版本

## 安裝
```
go get -v -u github.com/Jockey66666/fx2preset-lite
```

# 功能介紹
## 還原文件 (Windows版只支援刪除)
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

[![Travis CI](https://travis-ci.org/Jockey66666/fx2-tool.svg?branch=master)](https://travis-ci.org/Jockey66666/fx2-tool)
![macOS](https://github.com/Jockey66666/fx2-tool/workflows/macOS/badge.svg)
![Windows](https://github.com/Jockey66666/fx2-tool/workflows/Windows/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jockey66666/fx2-tool)](https://goreportcard.com/report/github.com/jockey66666/fx2-tool)


# fx2-tool
## 簡介
內部使用工具，使用前要先安裝 [go](https://golang.org/) 1.14以後的版本

## 安裝
```
go get -v -u github.com/Jockey66666/fx2-tool
```

# 功能介紹
## 還原文件 (Windows版只支援刪除)
```
fx2-tool restore
```

## 列出所有preset
```
fx2-tool list
```

## 打開preset目錄
```
fx2-tool find "American Dream"
```

## 登出tonecloud
```
fx2-tool logout
```

## 切換tonecloud環境
### prodction
```
fx2-tool env prod
```

### staging
```
fx2-tool env stage
```

[![Travis CI](https://travis-ci.org/Jockey66666/fx2tool.svg?branch=master)](https://travis-ci.org/Jockey66666/fx2tool)
![macOS](https://github.com/Jockey66666/fx2tool/workflows/macOS/badge.svg)
![Windows](https://github.com/Jockey66666/fx2tool/workflows/Windows/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/jockey66666/fx2tool)](https://goreportcard.com/report/github.com/jockey66666/fx2tool)


# fx2tool
## 簡介
內部使用工具，使用前要先安裝 [go](https://golang.org/) 1.14以後的版本

## 安裝
### command-line版本

```
go get -v -u github.com/Jockey66666/fx2tool
```

### gui版本

```
go get -v -u github.com/Jockey66666/fx2tool/fx2tool-ui
```

# 功能介紹
## 還原文件 (Windows版只支援刪除)
```
fx2tool restore
```

## 列出所有preset
```
fx2tool list
```

## 打開preset目錄
```
fx2tool find "American Dream"
```

## 登出tonecloud
```
fx2tool logout
```

## 切換tonecloud環境
### prodction
```
fx2tool env prod
```

### staging
```
fx2tool env stage
```

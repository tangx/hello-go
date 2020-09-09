# hello-go

![Go](https://github.com/tangx/hello-go/workflows/Go/badge.svg)

hello-go 测试

### golang 库

+ `highLight` ( github.com/alecthomas/chroma/quick ) : 彩色输出
+ `goconvey` ( github.com/smartystreets/goconvey/convey ) : 代码测试
+ `cronjob` ( github.com/robfig/cron/v3 ) : 定时任务

## Commit Guide

+ [CommitGuide.README.md](CommitGuide.README.md)

## template

在 `golang` 中， 有很多 template。

+ `html/template` 出于安全考虑， 会自动转义 `<` 之类的符号。
+ `text/template` 有什么是什么。

## NoteBook

1. 直接在 `github/workflows` 下是用 `run` 模块， 无法读取环境变量。 无法使用 `$()` 获取执行结果。 但可以在预设值的 `build.sh` 中完成。



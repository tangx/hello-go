# hello-go

![Go](https://github.com/tangx/hello-go/workflows/Go/badge.svg)

hello-go 测试

## NoteBook

1. 直接在 `github/workflows` 下是用 `run` 模块， 无法读取环境变量。 无法使用 `$()` 获取执行结果。 但可以在预设值的 `build.sh` 中完成。



## Commit Guide

+ example: `fix: correct the Keybase Build Notifications marketplace URL`

```ts
enum ConventionalCommitTypes {
  feat = 'Features', // 功能点
  fix = 'Bug Fixes', // bug 修复
  docs = 'Documentation', // 文档改动
  style = 'Styles',  // 格式化, 缺失分号等; 不包括生产代码变动
  refactor = 'Code Refactoring',  // 重构代码
  perf = 'Performance Improvements', // 性能提升
  test = 'Tests',  // 添加缺失的测试, 重构测试, 不包括生产代码变动
  build = 'Builds', // 编译
  ci = 'Continuous Integration',  // CI
  chore = 'Chores',  // 杂事, 更新grunt任务等; 不包括生产代码变动
  revert = 'Reverts',  // 回滚
}
```


## template

在 `golang` 中， 有很多 template。
+ `html/template` 出于安全考虑， 会自动转义 `<` 之类的符号。
+ `text/template` 有什么是什么。

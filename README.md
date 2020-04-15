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
  docs = 'Documentation', // 文档
  style = 'Styles',  // 样式
  refactor = 'Code Refactoring',  // 重构
  perf = 'Performance Improvements', // 性能提升
  test = 'Tests',  // 测试
  build = 'Builds', // 编译
  ci = 'Continuous Integration',  // CI
  chore = 'Chores',  // 杂事
  revert = 'Reverts',  // 回滚
}
```

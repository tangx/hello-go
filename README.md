# hello-go

![Go](https://github.com/tangx/hello-go/workflows/Go/badge.svg)

hello-go 测试

## NoteBook

1. 直接在 `github/workflows` 下是用 `run` 模块， 无法读取环境变量。 无法使用 `$()` 获取执行结果。 但可以在预设值的 `build.sh` 中完成。



## Commit Guide

+ example: `fix: correct the Keybase Build Notifications marketplace URL`

```ts
enum ConventionalCommitTypes {
  feat = 'Features',
  fix = 'Bug Fixes',
  docs = 'Documentation',
  style = 'Styles',
  refactor = 'Code Refactoring',
  perf = 'Performance Improvements',
  test = 'Tests',
  build = 'Builds',
  ci = 'Continuous Integration',
  chore = 'Chores',
  revert = 'Reverts',
}
```


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

# Log 日志包（不推荐使用，目前存在很多问题）

基于`slog`的二次封装

## 支持功能

> 基础功能
+ 支持基础的日志信息 [☑️]
+ 支持不同的日志级别 [☑️]
+ 支持自定义配置
+ 支持输出到标准输出和保存到文件
> 高级功能
+ 支持多种日志格式
+ 支持按级别分类输出
+ 支持结构化日志
+ 支持日志轮转
+ 具备Hook能力
> 额外可选功能
+ 支持颜色输出
+ 兼容标准库log
+ 支持输出到不同位置

## Inspiration
- [江湖十年:万字解析slog](https://jianghushinian.cn/2024/06/24/go-s-official-structured-log-package-slog/#slog-%E5%BF%AB%E9%80%9F%E5%85%A5%E9%97%A8)
- [江湖十年:使用go第三方日志库zap](https://jianghushinian.cn/2023/03/19/use-of-zap-in-go-third-party-log-library/)
- [江湖十年:如何优雅地封装一个更友好的日志库](https://jianghushinian.cn/2023/04/16/how-to-wrap-a-more-user-friendly-logging-package-based-on-zap/]https://jianghushinian.cn/2023/04/16/how-to-wrap-a-more-user-friendly-logging-package-based-on-zap/)
- [zap](https://pkg.go.dev/go.uber.org/zap)
- [marmotedu/log](https://github.com/marmotedu/iam/pkg/log)
- [slog](https://pkg.go.dev/log/slog))
- [log](https://pkg.go.dev/log)
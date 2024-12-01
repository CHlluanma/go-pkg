# Config 包

Config 包是对`spf13/Viper`的二次封装，目的是提供一个简单的、默认的配置文件读取逻辑，支持`json`、`yaml`、`toml`等格式.

在默认情况下，将根据`环境变量`->`配置文件`的顺序读取配置，

默认的配置文件为`./config/${MODE_ENV}/config.yaml`，默认的`MODE_ENV`为`dev`，使用`LoadConfig`方法读取配置.

同时支持读取`.env`文件，默认的配置文件为`${ROOT}/dev.env`，使用`LoadDotEnv`方法读取配置.


## TODO:

- 支持配置的优先级
- 支持多配置文件
- 支持远程配置（etcd，consul...)
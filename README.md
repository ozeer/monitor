## Prometheus监控Demo

## 遇到的问题
1、服务启动报错：error getting credentials - err: exec: "docker-credential-desktop": executable file not found in $PATH, out: ``
解决办法：
```
vi ~/.docker/config.json
# 将credsStore 改成 credStore
```

2、启动报错：mysql_exporter caller=config.go:150 level=error msg="failed to validate config"
解决办法：


## 参考文章或代码
- https://github.com/xhyonline/xutil
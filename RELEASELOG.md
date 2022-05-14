## Release Log

### `v0.1.7 `

发布于 2022-05-14 👏

🎉特性

1. 升级 [Meta Kennel](https://github.com/MetaCubeX/Clash.Meta) 为 Alpha
2. 域名嗅探，

   - 用于嗅探TCP请求中实际的域名，而非clash的dns映射获取

   ```yaml
   sniffer:
      enable: true #控制开关
      sniffig:
        - tls
        - http
      port-whitelist: #目的端口白名单，嗅探器只会嗅探白名单中的端口，默认0-65535，推荐设置成常见端口
        - 80
        - 443
        - 8000-9000
      skip-domain: # 嗅探的域名结果如果在此名单则不会生效
        - baidu.com
        - google.com
      force-domain: # 需要嗅探的域名，这里域名是clash原有逻辑获取的域名，如为空则只会嗅探IP请求，填写'+'则嗅探所有请求
        - +.qq.com
   ```
3. TCP并发连接
   ```yaml
   tcp-concurrent: true #默认为false
   ```
4. Relay策略组

   - Relay策略可以利用udo over tcp的协议支持UDP

5. 策略组过滤节点优化

   - 优化节点过滤逻辑，当前将不会每次请求进行一次过滤匹配

🎇其他
- ipv6: `false` 将完全关闭IPv6请求，不允许IPv6请求连接，包括纯IPv6
- DOQ环流问题优化

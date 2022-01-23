## Release Log

### `v0.1.6 `

发布于 2022-01-23 19:30 👏

🎉特性

1. 升级 [Meta Kennel](https://github.com/Clash-Mini/Clash.Mini/clash) 为 1.9.0 Dev
2. 支持GeoSite延迟加载
3. 允许策略组为空，默认DIRECT
4. 新增 `Network` 规则, 支持匹配网络类型 ( TCP / UDP )
5. 新增多条件规则 ( `NOT` `OR` `AND` )
    ```yaml  
    -AND,((DOMAIN,baidu.com),(NETWORK,UDP)),REJECT
    -OR,((DOMAIN,baidu.com),(NETWORK,UDP)),REJECT
    -NOT,(DOMAIN,baidu.com),REJECT 
    ```
6. Linux AutoRoute模块支持 ip route

🎇修复

1. 数不清的问题

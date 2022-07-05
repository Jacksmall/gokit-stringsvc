# gokit-stringsvc
gokit stringsvc 示例

>go run .
```
```

新开终端:
```
curl -XPOST -d'{"s":"hello,world!"}' localhost:8082/uppercase
{"v":"HELLO,WORLD!","err":""}

curl -XPOST -d'{"s":"hello,world!"}' localhost:8082/count
{"v":12}
```

未波及到数据库,使用net/http 官方库

# gokit-stringsvc
gokit stringsvc 示例

>go run .
```
msg=HTTP addr=:8082
```

新开终端:
```
curl -XPOST -d'{"s":"hello,world!"}' localhost:8082/uppercase
{"v":"HELLO,WORLD!","err":""}

server:
method=uppercase input=hello,world! output=HELLO,WORLD! took=832ns

curl -XPOST -d'{"s":"hello,world!"}' localhost:8082/count
{"v":12}

server:
method=count input=hello,world! output=12 took=345ns
```

新增log和instrumention中间件
页面访问:http://localhost:8082/metrics
可以看到请求的总数，请求的延迟情况，countResult
```
...
# HELP my_group_string_service_count_result The result of each count method.
# TYPE my_group_string_service_count_result summary
my_group_string_service_count_result_sum 12
my_group_string_service_count_result_count 1
# HELP my_group_string_service_request_count Number of requests received.
# TYPE my_group_string_service_request_count counter
my_group_string_service_request_count{error="false",method="count"} 1
my_group_string_service_request_count{error="false",method="uppercase"} 1
# HELP my_group_string_service_request_latency_microseconds Total duration of requests in microseconds.
# TYPE my_group_string_service_request_latency_microseconds summary
my_group_string_service_request_latency_microseconds_sum{error="false",method="count"} 3.7647e-05
my_group_string_service_request_latency_microseconds_count{error="false",method="count"} 1
my_group_string_service_request_latency_microseconds_sum{error="false",method="uppercase"} 3.7717e-05
my_group_string_service_request_latency_microseconds_count{error="false",method="uppercase"} 1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 4
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

未波及到数据库,使用net/http 官方库

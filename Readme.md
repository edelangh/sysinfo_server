
# Description

This program create a server called sysinfo_server which allow user to retriev informations about their system boot time

# Requirements

 * go 1.17
 * systemd-analyze

# Examples
```
$ ./sysinfo_server &
[1] 31579
$ 2021/09/18 13:37:52 Server ready, endpoints: /version and /duration

$ curl http://localhost:8080/
Start by browsing /version or /duration
Try also with Header "Accept: application/json"
$ curl http://localhost:8080/version
v1.1.0
$ curl http://localhost:8080/duration
26.030s
$ curl -H "Accept: application/json" localhost:8080/version
{"version":"v1.1.0"}
$ curl -H "Accept: application/json" localhost:8080/duration
{"duration":"26.030s"}
```

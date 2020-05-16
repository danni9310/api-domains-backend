# Backend API domains

API aimed to search information about different domains

![Logo](https://miro.medium.com/max/920/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg)

- Set GOPATH

`export GOPATH=/home/user/backend-domains`

- Validate GOPATH

`echo $GOPATH`

- Run project

`go run src/main.go`

## Endpoints

One domain

`curl --request GET 'http://192.168.0.13:1206/domain/facebook.com'`

Response

```json{ "Servers_changed": false, "Ssl_grade": "B", "Previous_ssl_grade": "B", "Logo": "https://static.xx.fbcdn.net/rsrc.php/yz/r/KFyVIAWzntM.ico", "Title": "Facebook - Inicia sesión o regístrate", "Is_down": false, "Servers": [ { "Address": "157.240.11.35", "Ssl_grade": "B", "Country": "US", "Owner": "Facebook, Inc. (THEFA-3)" }, { "Address": "2a03:2880:f127:283:face:b00c:0:25de", "Ssl_grade": "B", "Country": "IE", "Owner": "" } ] }```

History domains

`curl --request GET 'http://192.168.0.13:1206/domain'`

Response

```json{ "Item": [ "facebook.com" ] }```

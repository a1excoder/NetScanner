# NetScanner

<h3>Simple console application for check status port</h3>
<hr>

build and run
```bash
git clone https://github.com/a1excoder/NetScanner
cd NetScanner/
go build

./NetScanner # output - [warning] after bin name write params, <--help>
```

how to use
```bash
./NetScanner --help
# output
: '
check only one port <./NetScanner -s -ip 192.168.0.1 -port 443>
check many(from to) ports <./NetScanner -sm -ip 192.168.0.1 -pr 22 444>
check many(from to) ports <./NetScanner -sm -ip 192.168.0.1 -pr 22 444 -t 2>
check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444>
check many(from to) ports and show only opened <./NetScanner -smo -ip 192.168.0.1 -pr 22 444 -t 2>
send an http request via Get and get a response <./NetScanner -rg https://jsonplaceholder.typicode.com/posts/1>
'
```

```bash
# parameters
# --help : help
# -s : scan
# -ip : IPv4 host address
# -port : what port do you want to check
# -sm : scan many
# -pr : port radius
# -smo : scan many opened
# -pg : send an http Get request
```

```bash
# check port 80
./NetScanner -s -ip 192.168.0.1 -port 80
# my output
# 192.168.0.1 host is scanning
# port 80 is opened

# if close
# 192.168.0.1 host is scanning
# port 80 is closed
```

```bash
# check many ports
./NetScanner -sm -ip 192.168.0.1 -pr 78 83 # check from 78 to 83
# my output output
# 192.168.0.1 host is scanning
# port 78 is closed
# port 79 is closed
# port 80 is opened (HTTP)
# port 81 is opened (HOSTS2-NS)
# port 82 is closed
# port 83 is closed
'
```

```bash
# check many opened ports
./NetScanner -smo -ip 192.168.0.1 -pr 78 83 # check from 78 to 83
# my output output
# 192.168.0.1 host is scanning
# port 80 is opened (HTTP)
# port 81 is opened (HOSTS2-NS)
```

```bash
# send an http request via Get and get a response
# output
: '
Code status: 200
{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}
'
```

# namify

Replaces IP addresses with according domain names.

## Install

```
$ go get -u github.com/simonuvarov/namify
```

## Basic Usage

```
$ head recon/example/nmap.xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE nmaprun>
```

```
$ cat recon/example/nmap.xml | httpull | tee httpull.txt
http://1.1.1.1
http://1.1.1.1:8080
https://4.4.4.4:8443
```

```
$ cat resolved.txt
1.1.1.1 example.com,www.example.com
4.4.4.4 admin.example.com
```

```
$ namify -u httpull.txt -d resolved.txt
http://example.com
http://example.com:8080
http://wwww.example.com
http://www.example.com:8080
https://admin.example.com:8443
```

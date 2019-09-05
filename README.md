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
http://2.2.2.2:8080
https://4.4.4.4:8443
```

```
$ namify -u httpull.txt -d resolved.txt
http://example.com
http://www.example.com:8080
https://admin.example.com:8443
```

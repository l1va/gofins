# gofins

[![Build Status](https://travis-ci.org/l1va/gofins.svg?branch=master)](https://travis-ci.org/l1va/gofins)


This is fins command client written by Go.

The library support communication to omron PLC from Go application.

Ideas were taken from https://github.com/hiroeorz/omron-fins-go and https://github.com/patrick--/node-omron-fins

Library was tested with Omron PLC NJ501-1300.

<b>PS.</b>Build is ok but Travis is failing to run tests with go 1.10
(with earlier go versions tests are passed too):
```
 === RUN   TestFinsClient
 2018/03/11 18:45:18 read udp 127.0.0.1:33746->127.0.0.1:9600: use of closed network connection
 FAIL   github.com/l1va/gofins/fins	0.004s
 The command "go test -v ./..." exited with 1.
 ```
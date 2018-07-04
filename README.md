# GoFINS

[![Build Status](https://travis-ci.org/l1va/gofins.svg?branch=master)](https://travis-ci.org/l1va/gofins)


This is fins command client written by Go.

The library support communication to omron PLC from Go application.

Ideas were taken from https://github.com/hiroeorz/omron-fins-go and https://github.com/patrick--/node-omron-fins

Library was tested with <b>Omron PLC NJ501-1300</b>. Mean time of the cycle request-response is 4ms.

Feel free to ask questions, raise issues and make pull requests!


<b>PS. Build is ok</b> but Travis is randomly failing to run tests with <b>go 1.10</b>
(with earlier go versions tests are passed too): [travis](https://travis-ci.org/l1va/gofins)
```
 === RUN   TestFinsClient
 2018/03/11 18:45:18 read udp 127.0.0.1:33746->127.0.0.1:9600: use of closed network connection
 FAIL   github.com/l1va/gofins/fins	0.004s
 The command "go test -v ./..." exited with 1.
 ```
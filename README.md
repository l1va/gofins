# GoFINS

[![Build Status](https://travis-ci.org/l1va/gofins.svg?branch=master)](https://travis-ci.org/l1va/gofins)

This is fins command client written by Go.

The library support communication to omron PLC from Go application.

Ideas were taken from https://github.com/hiroeorz/omron-fins-go and https://github.com/patrick--/node-omron-fins

Library was tested with <b>Omron PLC NJ501-1300</b>. Mean time of the cycle request-response is 4ms.
Additional work in the siyka-au repository was tested against a <b>CP1L-EM</b>.

There is simple Omron FINS Server (PLC emulator) in the fins/server.go 

Feel free to ask questions, raise issues and make pull requests!

 ### Thanks
 [malleblas](https://github.com/malleblas) for his PR: https://github.com/l1va/gofins/pull/1

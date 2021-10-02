#!/bin/sh
echo '*1\r\n$4\r\nping\r\n' | nc localhost 6379
echo '*2\r\n$4\r\necho\r\n$3\r\nhey\r\n' | nc localhost 6379
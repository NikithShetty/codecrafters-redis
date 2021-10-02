#!/bin/sh
echo '*1\r\n$4\r\nping\r\n' | nc localhost 6379
echo '*2\r\n$4\r\necho\r\n$3\r\nhey\r\n' | nc localhost 6379
echo '*3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nhello\r\n' | nc localhost 6379
echo '*2\r\n$3\r\nget\r\n$3\r\nkey\r\n' | nc localhost 6379
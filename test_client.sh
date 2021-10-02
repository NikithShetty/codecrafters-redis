#!/bin/sh
echo '*1\r\n$4\r\nping\r\n' | nc localhost 6379
echo '*2\r\n$4\r\necho\r\n$3\r\nhey\r\n' | nc localhost 6379
echo '*3\r\n$3\r\nset\r\n$4\r\nheya\r\n$5\r\nhello\r\n' | nc localhost 6379
echo '*2\r\n$3\r\nget\r\n$4\r\nheya\r\n' | nc localhost 6379
echo '*2\r\n$3\r\nget\r\n$4\r\nkey1\r\n' | nc localhost 6379

echo 'Test set px'

echo '*5\r\n$3\r\nset\r\n$4\r\nkey1\r\n$5\r\nhello\r\n$2\r\npx\r\n$4\r\n1000\r\n' | nc localhost 6379
echo '*2\r\n$3\r\nget\r\n$4\r\nkey1\r\n' | nc localhost 6379
sleep 1.1s
echo '*2\r\n$3\r\nget\r\n$4\r\nkey1\r\n' | nc localhost 6379
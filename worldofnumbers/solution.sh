#!/bin/bash

v="$(< /dev/stdin)"
a="$(echo "$v" | tr '\n' ' ' | awk '{print $1}')"
b="$(echo "$v" | tr '\n' ' ' | awk '{print $2}')"

echo "$((a + b))"
echo "$((a - b))"
echo "$((a * b))"
echo "$((a / b))"

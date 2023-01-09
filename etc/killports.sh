#!/bin/bash

YELLOW="\033[0;30m\033[33m"
GREEN="\033[0;30m\033[32m"
RED="\033[0;30m\033[31m"
STREND="\033[0m\n"


counter=0
process=$(lsof -ti:1323)
if [ ! -z "$process" ]; then
  kill -9 $process
  counter=$((counter + 1))
fi


case $counter in
1|0) printf "${GREEN}$counter processes killed ${STREND}";;
2) printf "${YELLOW}$counter processes killed ${STREND}";;
3) printf "${RED}$counter processes killed ${STREND}";;
esac
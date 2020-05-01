#!/bin/sh

HOST=chloelights
FILE=project1042

set -e

cat ${FILE} | xz -z | pv | ssh ${HOST} "dd | xzcat > ${FILE} && chmod +x ${FILE}" 

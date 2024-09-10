#!/bin/bash
set -e
set -x

echo "`date` Installing throughput on this machine"

mkdir -p ~/throughput
EXE_DIR=~/throughput
CURR_DIR=$(pwd)
go build -o build/throughputd ./cmd/throughputd
ln -s ${CURR_DIR}/build/throughputd ${EXE_DIR}/throughputd
ln -s ${CURR_DIR}/start.sh ${EXE_DIR}/start.sh
cp -f val[23]*.toml ${EXE_DIR}/

echo "`date` Installation completed"
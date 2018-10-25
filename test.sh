#!/bin/bash

mkdir -p $TEST_RESULTS/gomoku
mkdir -p $TEST_RESULTS/pepito

go test -v gomoku/gomoku 2>&1 | go-junit-report > $TEST_RESULTS/gomoku/report.xml
go test -v gomoku/pepito 2>&1 | go-junit-report > $TEST_RESULTS/pepito/report.xml

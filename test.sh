#!/bin/bash

go test -v gomoku/gomoku 2>&1 | go-junit-report > $TEST_RESULTS/report.xml
go test -v gomoku/pepito 2>&1 | go-junit-report > $TEST_RESULTS/report.xml

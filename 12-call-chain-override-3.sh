#!/usr/bin/env bash
source 00-common.sh

curlie localhost:8080 "x-override: service-3:v2"

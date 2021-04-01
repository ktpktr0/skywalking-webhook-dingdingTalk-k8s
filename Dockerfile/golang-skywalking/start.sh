#!/bin/sh

exec ${SKYWALKING_WORK_SPACE}/webhook/${APP_NAME} --config configs/conf.yml

wait $currentPID
printOK $?

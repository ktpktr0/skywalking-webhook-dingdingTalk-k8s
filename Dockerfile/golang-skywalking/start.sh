#!/bin/sh

exec ${SKYWALKING_WORK_SPACE}/webhook/${APP_NAME} --config ${SKYWALKING_WEBHOOK_CONFIG_DIR}/conf.yml

wait $currentPID
printOK $?

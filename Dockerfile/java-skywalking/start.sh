#!/bin/sh

exec java -jar ${SKYWALKING_WORK_SPACE}/webhook/${APP_NAME} --spring.config.location=${SKYWALKING_WEBHOOK_CONFIG_DIR}/application.properties

wait $currentPID
printOK $?

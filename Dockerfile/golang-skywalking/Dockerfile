#FROM openjdk:8u92-alpine
FROM alpine:3.8
MAINTAINER web-xiaxia<xiongchuang@vip.qq.com>

#设置环境变量
ENV SKYWALKING_WORK_SPACE=/skywalking \
    APP_NAME=golang-skywalking-webhook \
    SKYWALKING_WEBHOOK_CONFIG_DIR=/skywalking/webhook/configs

#执行需要的内容
RUN mkdir -p ${SKYWALKING_WORK_SPACE}/webhook && \
    mkdir ${SKYWALKING_WEBHOOK_CONFIG_DIR} 

#复制内容
COPY ${APP_NAME} ${SKYWALKING_WORK_SPACE}/webhook

#复制启动脚本
COPY start.sh ${SKYWALKING_WORK_SPACE}

#执行权限
RUN chmod 775 ${SKYWALKING_WORK_SPACE}/start.sh && \
    chmod 775 ${SKYWALKING_WORK_SPACE}/webhook/${APP_NAME}

#工作空间
WORKDIR ${SKYWALKING_WORK_SPACE}
EXPOSE 8080
CMD ["/skywalking/start.sh"]

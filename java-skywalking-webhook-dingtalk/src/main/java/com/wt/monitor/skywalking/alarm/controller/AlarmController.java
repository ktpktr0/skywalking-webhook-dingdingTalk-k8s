package com.wt.monitor.skywalking.alarm.controller;

import com.alibaba.fastjson.JSON;
import com.wt.monitor.skywalking.alarm.domain.AlarmMessageDTO;
import com.wt.monitor.skywalking.alarm.service.DingTalkAlarmService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.text.MessageFormat;
import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.List;

/**
 * @author ChengJianSheng
 * @date 2020/12/1
 */
@Slf4j
@RestController
@RequestMapping("/skywalking")
public class AlarmController {

    @Autowired
    private DingTalkAlarmService dingTalkAlarmService;

    @PostMapping("/alarm")
    public void alarm(@RequestBody List<AlarmMessageDTO> alarmMessageDTOList) {
       log.info("收到告警信息: {}", JSON.toJSONString(alarmMessageDTOList));
       if (null != alarmMessageDTOList) {
           alarmMessageDTOList.forEach(e -> {
               LocalDateTime startDateTime = LocalDateTime.ofEpochSecond(e.getStartTime()/1000, 0, ZoneOffset.ofHours(8));
               String msg = MessageFormat.format("-----来自SkyWalking的告警-----\n【time】: {0}\n【scope】: {1}\n【name】: {2}\n【message】: {3}", startDateTime.toString(), e.getScope(), e.getName(), e.getAlarmMessage());
               dingTalkAlarmService.sendMessage(msg);
           });
       }
    }
}

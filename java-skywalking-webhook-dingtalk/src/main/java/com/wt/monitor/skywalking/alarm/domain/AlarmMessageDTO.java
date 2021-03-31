package com.wt.monitor.skywalking.alarm.domain;

import lombok.Data;

import java.io.Serializable;

/**
 * @author ChengJianSheng
 * @data 2020/12/1
 */
@Data
public class AlarmMessageDTO implements Serializable {

    private int scopeId;

    private String scope;

    /**
     * Target scope entity name
     */
    private String name;

    private String id0;

    private String id1;

    private String ruleName;

    /**
     * Alarm text message
     */
    private String alarmMessage;

    /**
     * Alarm time measured in milliseconds
     */
    private long startTime;

}

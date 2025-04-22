package com.hjc.todolist.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.ser.std.ToStringSerializer;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class Task {
    private Long taskId;

    private String taskName;

    private String taskDescription;

    private Long userId;

    private int status;

    private int isDeleted;

    private LocalDateTime createTime;

    private Integer createUserId;

    private Integer updateUserId;

    private LocalDateTime updateTime;
}

package com.hjc.todolist.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.ser.std.ToStringSerializer;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class Task {
    @JsonSerialize(using = ToStringSerializer.class)
    private Long taskId;

    private String taskName;

    private String taskDescription;

    private Long userId;

    private int status;

    private int isDeleted;

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private Date createTime;

    private Integer createUserId;

    private Integer updateUserId;

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private Date updateTime;
}

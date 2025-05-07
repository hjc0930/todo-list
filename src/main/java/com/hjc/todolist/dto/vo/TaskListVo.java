package com.hjc.todolist.dto.vo;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.ser.std.ToStringSerializer;
import lombok.Data;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;

@Data
public class TaskListVo {
    @JsonSerialize(using = ToStringSerializer.class)
    private Long taskId;

    private String taskName;

    private String taskDescription;

    private String userName;

    private int status;

    private int isDeleted;

    private OffsetDateTime createTime;

    private Integer createUserId;

    private Integer updateUserId;

    private OffsetDateTime updateTime;
}

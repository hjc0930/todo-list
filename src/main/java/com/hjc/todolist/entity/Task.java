package com.hjc.todolist.entity;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class Task {
    private Long taskId;

    private String taskName;

    private String taskDescription;

    private Long userId;

    private String status;

    private Byte isDeleted;

    private LocalDateTime createTime;

    private Integer createUserId;

    private Integer updateUserId;

    private LocalDateTime updateTime;

}

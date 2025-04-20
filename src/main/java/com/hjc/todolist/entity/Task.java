package com.hjc.todolist.entity;

import lombok.Data;
import java.util.Date;

@Data
public class Task {
    private Long taskId;

    private String taskName;

    private String taskDescription;

    private Long userId;

    private int status;

    private Byte isDeleted;

    private Date createTime;

    private Integer createUserId;

    private Integer updateUserId;

    private Date updateTime;

}

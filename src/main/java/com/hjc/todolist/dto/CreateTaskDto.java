package com.hjc.todolist.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

import java.time.LocalDateTime;
import java.util.Date;

@Data
public class CreateTaskDto {
    @NotBlank(message = "任务名称不能为空")
    private String taskName;
    @NotBlank(message = "任务描述不能为空")
    private String taskDescription;

    private int staus;

    private int isDeleted;

    private LocalDateTime createTime;

    private Integer createUserId;

    private Integer updateUserId;

    private LocalDateTime updateTime;
}

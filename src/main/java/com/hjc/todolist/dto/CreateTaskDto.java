package com.hjc.todolist.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class CreateTaskDto {
    @NotBlank(message = "任务名称不能为空")
    private String taskName;
    @NotBlank(message = "任务描述不能为空")
    private String taskDesc;
}

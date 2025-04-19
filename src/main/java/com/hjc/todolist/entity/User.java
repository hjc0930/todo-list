package com.hjc.todolist.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class User {
    private Long userId;
    private String userName;
    private String password;
    private String phoneNumber;
    private Byte locked;

    private Integer createUserId;
    private LocalDateTime createTime;
    private Integer updateUserId;
    private LocalDateTime updateTime;
}

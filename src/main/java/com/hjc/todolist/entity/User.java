package com.hjc.todolist.entity;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.ser.std.ToStringSerializer;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class User {
    @JsonSerialize(using = ToStringSerializer.class)
    private Long userId;

    private String userName;

    @JsonIgnore
    private String password;

    private String phoneNumber;

    private Byte locked;

    private Integer createUserId;

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private LocalDateTime createTime;

    private Integer updateUserId;

    @JsonFormat(pattern = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")
    private LocalDateTime updateTime;

}

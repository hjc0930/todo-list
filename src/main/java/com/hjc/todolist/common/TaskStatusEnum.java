package com.hjc.todolist.common;

public enum TaskStatusEnum {
    PENDING("草稿", 0),
    COMPLETED("已完成", 1);

    private String name;
    private int code;

    TaskStatusEnum(String name, int code) {
        this.name = name;
        this.code = code;
    }

    public String getName() {
        return name;
    }

    public int getCode() {
        return code;
    }
}

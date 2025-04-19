package com.hjc.todolist.util;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

import java.io.Serializable;
import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
@ToString
public class PageResult<T> implements Serializable {
    private int page;

    private int pageSize;

    private int totalCount;

    private List<T> list;

}

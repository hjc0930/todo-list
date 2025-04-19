package com.hjc.todolist.util;

import java.util.LinkedHashMap;
import java.util.Map;

public class PageQueryUtil extends LinkedHashMap<String,Object> {
    private int page;
    private int pageSize;

    public PageQueryUtil(Map<String, Object> params) {
        this.putAll(params);

        // Pagination parameters
        this.page = Integer.parseInt(params.get("page").toString());
        this.pageSize = Integer.parseInt(params.get("pageSize").toString());
        this.put("start", (page - 1) * pageSize);
        this.put("page", page);
        this.put("pageSize", pageSize);
    }

    public int getPage() {
        return page;
    }
    public void setPage(int page) {
        this.page = page;
    }
    public int getPageSize() {
        return pageSize;
    }
    public void setPageSize(int pageSize) {
        this.pageSize = pageSize;
    }

    @Override
    public String toString() {
        return "PageQueryUtil{" +
                "page=" + page +
                ", pageSize=" + pageSize +
                '}';
    }
}

package com.hjc.todolist.controller;

import com.hjc.todolist.dao.TaskMapper;
import com.hjc.todolist.service.TaskService;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;
import com.hjc.todolist.util.Result;
import com.hjc.todolist.util.ResultGenerator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;

@RestController
@RequestMapping("/task")
public class TaskController {
    @Autowired
    private TaskService taskService;

    @RequestMapping(value = "/list", method = RequestMethod.GET)
    public Result getTaskList(
            @RequestParam(required = false, defaultValue = "1") Integer page,
            @RequestParam(required = false, defaultValue = "10") Integer pageSize
    ) {
        HashMap<String, Object> map = new HashMap<>();
        map.put("page", page);
        map.put("pageSize", pageSize);
        PageQueryUtil pageQueryUtil = new PageQueryUtil(map);

        PageResult taskList = taskService.getTaskList(pageQueryUtil);
        return ResultGenerator.getSuccessResult(taskList);
    }
}

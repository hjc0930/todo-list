package com.hjc.todolist.controller;

import com.hjc.todolist.dto.CreateTaskDto;
import com.hjc.todolist.entity.Task;
import com.hjc.todolist.service.TaskService;
import com.hjc.todolist.util.PageQueryUtil;
import com.hjc.todolist.util.PageResult;
import com.hjc.todolist.util.Result;
import com.hjc.todolist.util.ResultGenerator;
import com.hjc.todolist.vo.TaskListVo;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;

@RestController
@RequestMapping("/task")
public class TaskController {
    @Autowired
    private TaskService taskService;

    @RequestMapping(value = "/list", method = RequestMethod.GET)
    public Result<PageResult<TaskListVo>> getTaskList(
            @RequestParam(required = false, defaultValue = "1") Integer page,
            @RequestParam(required = false, defaultValue = "10") Integer pageSize
    ) {
        HashMap<String, Object> map = new HashMap<>();
        map.put("page", page);
        map.put("pageSize", pageSize);
        PageQueryUtil pageQueryUtil = new PageQueryUtil(map);

        PageResult<TaskListVo> taskList = taskService.getTaskList(pageQueryUtil);
        return ResultGenerator.getSuccessResult(taskList);
    }

    @RequestMapping(value = "/create", method = RequestMethod.POST)
    public Result<String> addTask(@RequestBody @Valid CreateTaskDto task) {
        int taskId = taskService.addTask(task);
        if (taskId > 0) {
            return ResultGenerator.getSuccessResult("Task added successfully");
        } else {
            return ResultGenerator.genFailResult("Failed to add task");
        }
    }
}

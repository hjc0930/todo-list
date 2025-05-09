package com.hjc.todolist.config;

import com.hjc.todolist.common.NewBeeMallException;
import com.hjc.todolist.common.ServiceResultEnum;
import com.hjc.todolist.util.Result;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.validation.BindException;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.Objects;

@RestControllerAdvice
public class TodoListGlobalExceptionHandler {
    @ExceptionHandler(BindException.class)
    public Object bindException(BindException e) {
        Result<Object> result = new Result<>();
        result.setResultCode(510);
        BindingResult bindingResult = e.getBindingResult();

        result.setMessage(Objects.requireNonNull(bindingResult.getFieldError()).getDefaultMessage());

        return result;
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public Object bindException(MethodArgumentNotValidException e) {
        Result<Object> result = new Result<>();
        result.setResultCode(510);
        BindingResult bindingResult = e.getBindingResult();
        result.setMessage(Objects.requireNonNull(bindingResult.getFieldError()).getDefaultMessage());
        return result;
    }

    @ExceptionHandler(Exception.class)
    public Object handleException(Exception e, HttpServletRequest req) {
        Result result = new Result();
        result.setResultCode(500);
        //区分是否为自定义异常
        if (e instanceof NewBeeMallException) {
            result.setMessage(e.getMessage());
            if (e.getMessage().equals(ServiceResultEnum.NOT_LOGIN_ERROR.getMessage()) || e.getMessage().equals(ServiceResultEnum.TOKEN_EXPIRE_ERROR.getMessage())) {
                result.setResultCode(416);
            } else if (e.getMessage().equals(ServiceResultEnum.ADMIN_NOT_LOGIN_ERROR.getMessage()) || e.getMessage().equals(ServiceResultEnum.ADMIN_TOKEN_EXPIRE_ERROR.getMessage())) {
                result.setResultCode(419);
            }
        } else {
            e.printStackTrace();
            result.setMessage("未知异常，请查看控制台日志并检查配置文件。");
        }
        return result;

    }
}

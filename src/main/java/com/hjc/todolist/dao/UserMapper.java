package com.hjc.todolist.dao;

import com.hjc.todolist.entity.User;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface UserMapper {
    int insert(User user);

    int insertSelective(User user);

    /**
     * Login methods
     * @param userName
     * @param password
     * @return
     */
    User login(@Param("userName") String userName, @Param("userName") String password);
    User findUserById(@Param("userId") Long userId);
    List<User> findAllById(List<Long> userIds);
    void updateByPrimaryKey(User user);
    void  updateByPrimaryKeySelective(User user);
}

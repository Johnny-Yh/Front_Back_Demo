<template>
    <div class="about">
      <h1>关于员工管理系统</h1>
      <p>
        这是一个用于管理员工信息的前后端分离系统。系统提供了以下功能：
      </p>
      <ul>
        <li>添加新员工</li>
        <li>编辑员工信息</li>
        <li>删除员工</li>
        <li>查看所有员工列表</li>
      </ul>
      <p>
        系统前端使用 Vue 3 和 Element Plus 构建，后端使用 Go 和 Gin 框架，数据库使用SQLite3。
      </p>
    </div>
  </template>
  
  <script>
  export default {
    name: 'AboutView',
  };
  </script>
  
  <style scoped>
  .about {
    padding: 20px;
  }
  h1 {
    margin-bottom: 20px;
  }
  ul {
    margin: 20px 0;
    padding-left: 20px;
  }
  li {
    margin-bottom: 10px;
  }
  </style>
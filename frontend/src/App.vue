<template>
  <el-container>
    <el-header>
      <!-- Logo 和标题 -->
      <div class="header-content">
        <img src="@/assets/logo.png" alt="Logo" class="logo" />
        <h1>员工管理系统</h1>
      </div>

      <!-- 导航菜单 -->
      <el-menu
        mode="horizontal"
        background-color="#409EFF"
        text-color="#fff"
        active-text-color="#ffd04b"
        :router="true"
      >
        <el-menu-item index="/">首页</el-menu-item>
        <el-menu-item index="/employee-list">员工列表</el-menu-item> <!-- 新增员工列表菜单项 -->
        <el-menu-item index="/about">关于</el-menu-item>
      </el-menu>
    </el-header>

    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
body {
  margin: 0;
  font-family: Arial, sans-serif;
}

.el-header {
  background-color: #409EFF; /* 背景颜色 */
  color: white; /* 文字颜色 */
  display: flex; /* 使用 Flexbox 布局 */
  align-items: center; /* 垂直居中 */
  justify-content: space-between; /* 标题和菜单分布在两端 */
  height: 60px; /* 设置高度 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  padding: 0 20px; /* 添加左右内边距 */
}

.header-content {
  display: flex;
  align-items: center; /* Logo 和标题垂直居中 */
}

.logo {
  height: 40px; /* Logo 高度 */
  margin-right: 10px; /* Logo 和标题之间的间距 */
}

h1 {
  font-size: 24px; /* 标题字体大小 */
  margin: 0; /* 移除默认的外边距 */
}

.el-main {
  padding: 20px;
}
</style>
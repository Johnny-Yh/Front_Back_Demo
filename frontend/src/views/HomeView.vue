<template>
  <div class="home">
    <h1>
      <el-icon><User /></el-icon> 员工管理系统
    </h1>
    <p class="welcome-text">欢迎使用员工管理系统，您可以在这里管理员工信息。</p>
    <EmployeeList />
  </div>
</template>

<script>
import EmployeeList from '../components/EmployeeList.vue';
import { User } from '@element-plus/icons-vue'; // 引入图标

export default {
  components: { EmployeeList, User },
};
</script>

<style scoped>
.home {
  padding: 20px;
}
h1 {
  text-align: center;
  margin-bottom: 20px;
  font-size: 24px;
  color: #303133;
  display: flex;
  align-items: center;
  justify-content: center;
}
.welcome-text {
  text-align: center;
  color: #606266;
  margin-bottom: 10px; /* 减小底部间距 */
  line-height: 1.2; /* 调整行高 */
  padding: 0; /* 移除内边距 */
}
</style>
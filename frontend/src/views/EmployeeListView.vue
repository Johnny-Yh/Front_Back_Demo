<template>
    <div class="employee-list-view">
      <h1>员工列表</h1>
      <EmployeeList />
    </div>
  </template>
  
  <script>
  import EmployeeList from '../components/EmployeeList.vue';
  
  export default {
    components: { EmployeeList },
  };
  </script>
  
  <style scoped>
  .employee-list-view {
    padding: 20px;
  }
  h1 {
    text-align: center;
    margin-bottom: 20px;
    font-size: 24px;
    color: #303133;
  }
  </style>
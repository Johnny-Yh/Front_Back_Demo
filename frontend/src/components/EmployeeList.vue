<template>
    <div class="employee-list">
      <h2>员工列表</h2>
      <el-button type="primary" @click="addEmployee">添加新员工</el-button>
      <el-table :data="employees" style="width: 100%" stripe>
        <el-table-column prop="id" label="ID" width="100" />
        <el-table-column prop="name" label="姓名" />
        <el-table-column prop="age" label="年龄" width="100" />
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" @click="editEmployee(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="deleteEmployee(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
  
      <el-dialog v-model="showForm" :title="selectedEmployee.id ? '编辑员工' : '添加新员工'" width="30%">
        <EmployeeForm
          :employee="selectedEmployee"
          @save="handleSave"
          @close="showForm = false"
        />
      </el-dialog>
    </div>
  </template>
  
  <script>
  import EmployeeForm from './EmployeeForm.vue';
  import employeeService from '../services/employeeService';
  
  export default {
    components: { EmployeeForm },
    data() {
      return {
        employees: [],
        showForm: false,
        selectedEmployee: { id: null, name: '', age: '' },
      };
    },
    async created() {
      await this.fetchEmployees();
    },
    methods: {
      async fetchEmployees() {
        this.employees = await employeeService.findAll();
      },
      addEmployee() {
        this.selectedEmployee = { id: null, name: '', age: '' };
        this.showForm = true;
      },
      editEmployee(employee) {
        this.selectedEmployee = { ...employee };
        this.showForm = true;
      },
      async deleteEmployee(id) {
        await employeeService.remove(id);
        await this.fetchEmployees();
      },
      async handleSave(employee) {
        if (employee.id) {
          await employeeService.update(employee);
        } else {
          await employeeService.save(employee);
        }
        this.showForm = false;
        await this.fetchEmployees();
      },
    },
  };
  </script>
  
  <style scoped>
  .employee-list {
    padding: 20px;
  }
  h2 {
    margin-bottom: 20px;
  }
  </style>
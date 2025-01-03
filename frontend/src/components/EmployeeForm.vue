<template>
    <div>
      <el-form :model="localEmployee" label-width="80px">
        <el-form-item label="姓名" required>
          <el-input v-model="localEmployee.name" />
        </el-form-item>
        <el-form-item label="年龄" required>
          <el-input v-model="localEmployee.age" type="number" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button @click="close">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      employee: {
        type: Object,
        default: () => ({ id: null, name: '', age: '' }),
      },
    },
    emits: ['save', 'close'],
    data() {
      return {
        localEmployee: { ...this.employee },
      };
    },
    watch: {
      employee: {
        handler(newVal) {
          this.localEmployee = { ...newVal };
        },
        immediate: true,
      },
    },
    methods: {
      save() {
        const employee = {
            ...this.localEmployee, age: Number(this.localEmployee.age)
        };
        this.$emit('save', employee);
      },
      close() {
        this.$emit('close');
      },
    },
  };
  </script>
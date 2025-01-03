<template>
    <div>
        <el-row class="title-row">
            <h1>员工管理系统</h1>
        </el-row>
        <el-menu class="el-menu-demo" mode="horizontal" :router="true">
            <el-menu-item index="/selectAll">全部员工</el-menu-item>
            <el-menu-item index="/insert">添加</el-menu-item>
        </el-menu>
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    name: "index"
}
</script>

<style scoped>
.title-row {
    text-align: center;
    margin: 20px 0;
}

h1 {
    color: #303133;
    font-size: 28px;
    margin: 0;
}
</style>

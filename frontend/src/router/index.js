import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import AboutView from '../views/AboutView.vue';
import EmployeeListView from '../views/EmployeeListView.vue'; // 引入 EmployeeListView

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView,
  },
  {
    path: '/employee-list',
    name: 'employee-list',
    component: EmployeeListView, // 添加 EmployeeListView 路由
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
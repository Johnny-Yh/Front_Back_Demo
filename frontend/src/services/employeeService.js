import axios from 'axios';

const API_URL = 'http://127.0.0.1:8081/employee';

export default {
    // 获取所有员工
    async findAll() {
        const response = await axios.get(`${API_URL}/findAll`);
        return response.data;
    },

    // 获取单个员工
    async findById(id) {
        const response = await axios.get(`${API_URL}/findById/${id}`);
        return response.data;
    },

    // 添加员工
    async save(employee) {
        const response = await axios.post(`${API_URL}/save`, employee);
        return response.data;
    },

    // 更新员工
    async update(employee) {
        const response = await axios.post(`${API_URL}/update`, employee);
        return response.data;
    },

    // 删除员工
    async remove(id) {
        const response = await axios.delete(`${API_URL}/remove/${id}`);
        return response.data;
    },
};
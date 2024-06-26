import axios from 'axios';

const http = axios.create({
  timeout: 5000,
});

// 请求拦截器
http.interceptors.request.use(config => {

  let token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = token
  }

  // 在请求发送之前做一些处理
  return config;
}, error => {
  return Promise.reject(error);
});

// 返回拦截器
http.interceptors.response.use(response => {
  // 对响应数据做一些处理
  return response.data;
}, error => {
  return Promise.reject(error);
});

export default http;
<template>
    <div class="the-header">
      <h1>我的网站</h1>
      <!-- 其他导航链接或元素 -->
    </div>
  </template>
  
  <script>
  export default {
    name: 'TheHeader'
  }
  </script>
  
  <style scoped>
  .the-header {
    width: 100%;
    position: fixed;
    top: 0;
    left: 0;
    background-color: #333;
    color: white;
    z-index: 1000; /* 确保头部在页面最上层 */
  }
  </style>
<template>
    <div class="sidebar">
      <nav>
        <ul>
          <li><router-link to="/sqllist">SQL监控</router-link></li>
          <li><router-link to="/scriptlist">脚本监控</router-link></li>
          <li><router-link to="/nodelist">工作节点</router-link></li>
          <li><router-link to="/grouplist">告警通知组</router-link></li>
        </ul>
      </nav>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Sidebar'
  }
  </script>
  
  <style scoped>
  .dashboard-container {
    display: flex;
    flex-direction: column;
    height: 100vh; /* 全视口高度 */
  }
  
  .the-header {
    width: 100%;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000; /* 确保头部在页面最上层 */
    background-color: #ffffff; /* 头部背景颜色 */
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 添加阴影以提升层次感 */
    height: 60px; /* 头部高度，根据需要调整 */
    line-height: 60px; /* 使内容垂直居中 */
    padding: 0 20px; /* 头部内边距 */
    /* 如果需要响应式设计，可以添加媒体查询来调整样式 */
  }
  
  .main-container {
    display: flex;
    flex: 1;
    overflow: hidden; /* 防止内容溢出 */
  }
  
  .sidebar {
    width: 200px;
    background-color: #333;
    color: white;
  }
  
  .content-wrapper {
    flex: 1;
    padding: 20px;
    overflow-y: auto; /* 内容区域垂直滚动 */
  }
  
  /* 其他样式... */
  </style>
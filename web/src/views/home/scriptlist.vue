<template>
    <div class="script-container">
      <h2>脚本监控</h2>
      <ul class="script-list">
        <li v-for="script in scripts" :key="script.id">
          {{ script.name }} - {{ script.lastRun }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Script',
    data() {
      return {
        scripts: [
          { id: 1, name: '备份脚本', lastRun: '2024-08-10 10:00' },
          { id: 2, name: '清理脚本', lastRun: '2024-08-09 15:00' },
          // 更多脚本监控项...
        ]
      };
    }
  };
  </script>
  
  <style scoped>
  .script-container {
    padding: 20px;
  }
  
  .script-list {
    list-style: none;
    padding: 0;
  }
  
  .script-list li {
    padding: 8px;
    border-bottom: 1px solid #eee;
  }
  
  .script-list li:last-child {
    border-bottom: none;
  }
  </style>
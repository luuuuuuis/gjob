<template>
    <div class="alert-container">
      <h2>告警通知组</h2>
      <ul class="alert-list">
        <li v-for="alert in alerts" :key="alert.id">
          {{ alert.name }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Alert',
    data() {
      return {
        alerts: [
          { id: 1, name: '告警通知组1' },
          { id: 2, name: '告警通知组2' },
          // 更多告警组...
        ]
      };
    }
  };
  </script>
  
  <style scoped>
  .alert-container {
    padding: 20px;
  }
  
  .alert-list {
    list-style: none;
    padding: 0;
  }
  
  .alert-list li {
    padding: 8px;
    border-bottom: 1px solid #eee;
  }
  
  .alert-list li:last-child {
    border-bottom: none;
  }
  </style>
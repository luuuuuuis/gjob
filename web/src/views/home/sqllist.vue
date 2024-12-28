<template>
    <div class="sql-container">
      <h2>SQL监控</h2>
      <ul class="sql-list">
        <li v-for="sql in sqls" :key="sql.id">
          {{ sql.query }} - {{ sql.status }}
        </li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SQL',
    data() {
      return {
        sqls: [
          { id: 1, query: 'SELECT * FROM users', status: '正常' },
          { id: 2, query: 'SELECT * FROM orders', status: '警告' },
          // 更多SQL监控项...
        ]
      };
    }
  };
  </script>
  
  <style scoped>
  .sql-container {
    padding: 20px;
  }
  
  .sql-list {
    list-style: none;
    padding: 0;
  }
  
  .sql-list li {
    padding: 8px;
    border-bottom: 1px solid #eee;
  }
  
  .sql-list li:last-child {
    border-bottom: none;
  }
  </style>
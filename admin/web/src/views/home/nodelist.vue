<template>
  <div class="alert-container">
    <lay-table :columns="columns2" :data-source="dataSource2" v-model:expandKeys="expandKeys2">
    </lay-table>
  </div>
</template>

<script>
  export default {
    setup() {

      const columns2 = [
        { title:"编号", width:"80px", key:"id", fixed: "left", sort: "desc"},
        { title:"名称", width:"120px", key:"name" },
        { title:"IP", width:"80px", key:"ip"},
        { title:"状态", width:"80px", key:"status"},
      ]

      const dataSource2 = [
        {id:"1", name:"门户", ip:"192.168.1.3", status: true},
        {id:"2", name:"大数据", ip:"192.168.1.4", status: false},
      ]

      return {
        columns2,
        dataSource2,
      }
    }
  }
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
  .expand-content {
    width: 100%;
    padding: 0px 20px 20px 0px;
    .layui-progress {
      margin-top: 24px;
    }
  }
</style>
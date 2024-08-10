<template>
  <div class="layout-container">
    <div class="header-container">
      <lay-header>
        <div class="title">GJOB</div>
        <div class="layui-input-inline">
          <!-- <input type="text" name="keyword" required lay-verify="required" placeholder="请输入关键词" autocomplete="off" class="layui-input">
          <button id="searchButton" type="submit" class="layui-btn layui-btn-primary" lay-submit lay-filter="searchFilter">搜索</button> -->
          <lay-input size="xs" allow-clear=true prefix-icon="layui-icon-search" placeholder="任务名/用户/系统"></lay-input>
          <lay-button id="searchButton" type="primary"><lay-icon type="layui-icon-search"></lay-icon></lay-button>
        </div>
      </lay-header>
    </div>
    
    <div class="main-container">
      <lay-side>
        <lay-menu :selected-key="selectedKey" @change-selected-Key="changeSelectedKey" @change-open-keys="changeOpenKeys" v-model:openKeys="openKeys2" :tree="true">
          <lay-menu-item id="1"><router-link to="/nodelist">工作节点</router-link></lay-menu-item>
          <lay-menu-item id="2"><router-link to="/sqllist">SQL监控</router-link></lay-menu-item>
          <lay-menu-item id="3"><router-link to="/scriptlist">脚本监控</router-link></lay-menu-item> 
          <lay-menu-item id="3"><router-link to="/grouplist">告警通知组</router-link></lay-menu-item> 
        </lay-menu>
      </lay-side>

      <lay-body>
        content
      </lay-body>
    </div>
  </div>
</template>

<script>
// 组件逻辑
export default {
  // ...
}
</script>

<style>
  .title {
    float: left;
    font-size: 24px;
    margin-left: 20px;
    color: white;
    font-weight: bold;
  }
  .layout-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }

  .layout-container .header-container {
    line-height: 60px;
    text-align: center;
    background: #87ca9a;
    color: white;
  }

  .main-container {
    display: flex;
    flex: 1; /* 占据剩余空间 */
    overflow: auto; /* 如果内容溢出，允许滚动 */
  }

  .main-container .layui-side {
    /* 侧边栏样式 */
    background: #77c38c;
    color: white;
    text-align: center;
  }

  .main-container .layui-body {
    /* body内容区域样式 */
    flex: 1; /* 占据剩余空间 */
    background: #5FB878;
    color: white;
  }

  .layui-input-inline {
    float: right;
    margin-right: 20px;
  }

  /* 清除Layui默认的margin和padding */
  body, .layui-layout, .layui-layout * {
    margin: 0;
    padding: 0;
  }

  #searchButton {
    float: right;
  }
</style>
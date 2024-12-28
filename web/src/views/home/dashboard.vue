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
      <lay-page-header class="goback" content="详情页面" @back="handleBack"></lay-page-header>
    </div>
    
    <div class="main-container">
      <lay-side>
        <lay-menu class="sidemenu" :tree="true">
          <lay-menu-item class="sidemenudetail" id="1"><router-link to="/nodelist">工作节点</router-link></lay-menu-item>
          <lay-menu-item class="sidemenudetail" id="2"><router-link to="/sqllist">SQL监控</router-link></lay-menu-item>
          <lay-menu-item class="sidemenudetail" id="3"><router-link to="/scriptlist">脚本监控</router-link></lay-menu-item> 
          <lay-menu-item class="sidemenudetail" id="4"><router-link to="/grouplist">告警通知组</router-link></lay-menu-item> 
        </lay-menu>
      </lay-side>

      <lay-body>
        content
      </lay-body>
    </div>
  </div>
</template>

<script>
  const handleBack=()=>{
    alert('点击返回回调')
  }
</script>

<style>
  .goback {
    /* float: left; */
    /* 离左边60 */
    margin-left: 133px;
    /* 靠下 */
    margin-top: 20px;
  }
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

  .sidemenudetail {
    border-bottom: 2px solid #02205b57; /* 1px宽度的实线黑色边框 */
  }

  .layout-container .header-container {
    line-height: 60px;
    text-align: center;
    background: #16baaa;
    color: white;
  }

  .main-container {
    display: flex;
    flex: 1; /* 占据剩余空间 */
    overflow: auto; /* 如果内容溢出，允许滚动 */
  }

  .main-container .layui-side {
    /* 侧边栏样式 */
    background: #16b777;
    color: white;
    text-align: center;
  }
  /* 导航栏样式 */
  .sidemenu {
    background: #16b777;
  }

  .main-container .layui-body {
    /* body内容区域样式 */
    flex: 1; /* 占据剩余空间 */
    background: #f8f8f8;
    color: rgb(7, 7, 7);
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
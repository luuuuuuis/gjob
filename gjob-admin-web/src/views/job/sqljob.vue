<template>
  <h1>SQL任务</h1>
  <!-- 展示getAllSqlJobs()的返回值 -->
  <p>{{ _res_ }}</p>
</template>

<script setup lang="ts">
defineOptions({
  // name 作为一种规范最好必须写上并且和路由的name保持一致
  name: "ScriptJob"
});
import { getAllSqlJobs } from "@/api/sqljob";
const { _res_ } = getAllSqlJobs();
</script>

// 最简代码，也就是这些字段必须有
export default {
  path: "/job",
  meta: {
    title: "调度任务"
  },
  children: [
    {
      path: "/job/sql",
      name: "SqlJob",
      component: () => import("@/views/job/sqljob.vue"),
      meta: {
        title: "SQL任务"
      }
    },
    {
      path: "/job/script",
      name: "ScriptJob",
      component: () => import("@/views/job/scriptjob.vue"),
      meta: {
        title: "脚本任务"
      }
    }
  ]
};

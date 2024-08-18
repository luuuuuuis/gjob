// 最简代码，也就是这些字段必须有
export default {
  path: "/node",
  meta: {
    title: "工作节点"
  },
  component: () => import("@/views/node/index.vue")
};

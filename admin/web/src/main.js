
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import '@layui/layui-vue/lib/index.css'
// 导入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 导入图标视图
import * as ELIcons from '@element-plus/icons-vue'

const app = createApp(App)

// 将图标注册为全局组件
for (let iconName in ELIcons) {
	app.component(iconName, ELIcons[iconName])
}
// 引用element plus
app.use(ElementPlus)

app.use(router)

app.mount('#app')

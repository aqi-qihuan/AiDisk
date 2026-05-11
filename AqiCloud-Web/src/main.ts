import "@/access";
import "bytemd/dist/index.css";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import zhCn from "element-plus/es/locale/lang/zh-cn";
import enUS from "element-plus/es/locale/lang/en";
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import i18n from "./locales";
import { getStoredLocale, type Locale } from "./locales";
import { useThemeStore } from "@/store/theme";
import "@/styles/global-reset.css"; // 全局样式重置
import "@/styles/design-system.css"; // 设计系统CSS变量
import "@/styles/mobile.css"; // 移动端样式

// 修复 ResizeObserver loop 警告
const resizeObserverErrorHandler = (e: ErrorEvent) => {
  if (e.message === 'ResizeObserver loop completed with undelivered notifications.') {
    e.stopImmediatePropagation();
    e.preventDefault();
    e.stopPropagation();
  }
};

window.addEventListener('error', resizeObserverErrorHandler);

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);

const app = createApp(App);

app.use(router);

// 获取 Element Plus 语言包
function getElementPlusLocale(locale: Locale) {
  return locale === 'zh-CN' ? zhCn : enUS;
}

// 初始化 Element Plus
app.use(ElementPlus, {
  locale: getElementPlusLocale(getStoredLocale())
});

app.use(pinia);
app.use(i18n);

// 初始化主题（暗色模式为默认）
const themeStore = useThemeStore();
themeStore.initTheme();
themeStore.watchSystemTheme();

app.mount("#app");

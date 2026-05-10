import { createI18n } from 'vue-i18n'
import zhCN from './zh-CN.json'
import enUS from './en-US.json'

// 定义语言类型
export type Locale = 'zh-CN' | 'en-US'

// 获取浏览器语言
function getBrowserLocale(): Locale {
  const navigatorLocale = 
    navigator.languages !== undefined
      ? navigator.languages[0]
      : navigator.language
  
  if (!navigatorLocale) {
    return 'zh-CN'
  }

  // 简化语言代码
  const locale = navigatorLocale.split('-')[0]
  return locale === 'zh' ? 'zh-CN' : 'en-US'
}

// 获取存储的语言或浏览器语言
export function getStoredLocale(): Locale {
  const stored = localStorage.getItem('locale') as Locale | null
  return stored || getBrowserLocale()
}

// 存储语言设置
export function setStoredLocale(locale: Locale): void {
  localStorage.setItem('locale', locale)
}

// 创建 i18n 实例
const i18n = createI18n<[typeof zhCN, typeof enUS], Locale>({
  legacy: false, // 使用 Composition API 模式
  locale: getStoredLocale(),
  fallbackLocale: 'zh-CN', // 回退语言
  messages: {
    'zh-CN': zhCN,
    'en-US': enUS
  },
  globalInjection: true // 全局注入 $t
})

export default i18n

// 切换语言的辅助函数
export function changeLocale(locale: Locale): void {
  // @ts-ignore - TypeScript 类型推断问题，legacy: false 时需要 .value
  i18n.global.locale.value = locale
  setStoredLocale(locale)
  
  // 更新 HTML lang 属性
  document.documentElement.lang = locale
  
  // 更新 Element Plus 语言
  // 这里需要手动触发 Element Plus 的语言切换
  window.dispatchEvent(new CustomEvent('locale-change', { detail: locale }))
}

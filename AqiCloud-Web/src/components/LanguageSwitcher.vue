<template>
  <el-dropdown
    trigger="click"
    @command="handleLanguageChange"
    class="language-switcher"
    popper-class="lang-dropdown-hok"
  >
    <button class="language-btn">
      <svg class="globe-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10"></circle>
        <line x1="2" y1="12" x2="22" y2="12"></line>
        <path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path>
      </svg>
      <span class="lang-code">{{ currentLocaleCode }}</span>
      <svg class="chevron-icon" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
        <polyline points="6 9 12 15 18 9"></polyline>
      </svg>
    </button>

    <template #dropdown>
      <el-dropdown-menu class="lang-menu-hok">
        <el-dropdown-item
          v-for="lang in languages"
          :key="lang.value"
          :command="lang.value"
          :class="{ 'is-active': currentLocale === lang.value }"
          class="lang-item-hok"
        >
          <span class="lang-code-dropdown">{{ lang.code }}</span>
          <span class="lang-label">{{ lang.label }}</span>
          <svg
            v-if="currentLocale === lang.value"
            class="check-icon"
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <polyline points="20 6 9 17 4 12"></polyline>
          </svg>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { changeLocale, type Locale } from "@/locales";
import { ElMessage } from "element-plus";

const { locale, t } = useI18n();

// 支持的语言列表
const languages = [
  {
    value: "zh-CN" as Locale,
    label: "简体中文",
    code: "CN",
    flag: "🇨🇳",
  },
  {
    value: "en-US" as Locale,
    label: "English",
    code: "EN",
    flag: "🇺🇸",
  },
  {
    value: "ja-JP" as Locale,
    label: "日本語",
    code: "JP",
    flag: "🇯🇵",
  },
];

// 当前语言
const currentLocale = computed(() => locale.value as Locale);

// 当前语言代码简称
const currentLocaleCode = computed(() => {
  const lang = languages.find((l) => l.value === locale.value);
  return lang?.code || "CN";
});

// 切换语言
const handleLanguageChange = (langValue: Locale) => {
  if (langValue === locale.value) return;

  changeLocale(langValue);

  ElMessage.success({
    message:
      t("language.switch") +
      " " +
      (langValue === "zh-CN" ? "简体中文" : langValue === "ja-JP" ? "日本語" : "English"),
    duration: 2000,
  });
};
</script>

<style scoped>
/* ===== HOK 暗色主题 — 语言切换 ===== */
.language-switcher {
  display: inline-block;
}

.language-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  background: rgba(245, 158, 11, 0.1);
  border: 1px solid rgba(245, 158, 11, 0.25);
  border-radius: 8px;
  color: #fbbf24;
  cursor: pointer;
  transition: all 0.25s ease;
  font-size: 13px;
  font-weight: 700;
  font-family: "Fira Code", monospace;
  letter-spacing: 0.02em;
}

.language-btn:hover {
  background: rgba(245, 158, 11, 0.2);
  border-color: rgba(245, 158, 11, 0.5);
  transform: translateY(-1px);
  box-shadow: 0 2px 12px rgba(245, 158, 11, 0.2);
}

.language-btn:active {
  transform: translateY(0) scale(0.97);
}

.globe-icon {
  width: 16px;
  height: 16px;
  color: rgba(245, 158, 11, 0.7);
  flex-shrink: 0;
  transition: color 0.25s ease;
}

.language-btn:hover .globe-icon {
  color: #fbbf24;
}

.lang-code {
  font-size: 13px;
  line-height: 1;
  color: #fbbf24;
}

.chevron-icon {
  transition: transform 0.3s ease;
  color: rgba(245, 158, 11, 0.6);
}

.language-btn:hover .chevron-icon {
  transform: translateY(1px);
}

/* 移动端优化 */
@media (max-width: 768px) {
  .language-btn {
    padding: 6px 8px;
    gap: 4px;
  }

  .globe-icon {
    width: 14px;
    height: 14px;
  }

  .lang-code {
    font-size: 12px;
  }
}
</style>

<style>
/* 全局样式（非 scoped，覆盖 Element Plus 下拉菜单） */
.lang-dropdown-hok {
  background: #1a1a24 !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  border-radius: 12px !important;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.6) !important;
  padding: 6px !important;
}

.lang-item-hok {
  display: flex !important;
  align-items: center !important;
  gap: 10px !important;
  padding: 10px 14px !important;
  border-radius: 8px !important;
  color: #94a3b8 !important;
  transition: all 0.2s ease !important;
}

.lang-item-hok:hover {
  background: rgba(245, 158, 11, 0.08) !important;
  color: #fbbf24 !important;
}

.lang-item-hok.is-active {
  background: rgba(245, 158, 11, 0.12) !important;
  color: #fbbf24 !important;
  font-weight: 600 !important;
}

.lang-code-dropdown {
  font-size: 14px;
  font-weight: 700;
  font-family: "Fira Code", monospace;
  color: inherit;
  min-width: 24px;
}

.lang-label {
  flex: 1;
  font-size: 14px;
}

.check-icon {
  color: #fbbf24;
  margin-left: auto;
}
</style>

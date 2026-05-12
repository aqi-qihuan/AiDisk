<template>
  <el-dropdown
    trigger="click"
    @command="handleLanguageChange"
    class="language-switcher"
  >
    <button class="language-btn">
      <svg
        class="language-icon"
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <circle cx="12" cy="12" r="10"></circle>
        <line x1="2" y1="12" x2="22" y2="12"></line>
        <path
          d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"
        ></path>
      </svg>
      <span class="language-text">{{ currentLanguageLabel }}</span>
    </button>

    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="lang in languages"
          :key="lang.value"
          :command="lang.value"
          :class="{ 'is-active': currentLocale === lang.value }"
        >
          <span class="lang-flag">{{ lang.flag }}</span>
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
    flag: "🇨🇳",
  },
  {
    value: "en-US" as Locale,
    label: "English",
    flag: "🇺🇸",
  },
];

// 当前语言
const currentLocale = computed(() => locale.value as Locale);

// 当前语言标签
const currentLanguageLabel = computed(() => {
  const lang = languages.find((l) => l.value === locale.value);
  return lang?.flag || "🌐";
});

// 切换语言
const handleLanguageChange = (langValue: Locale) => {
  if (langValue === locale.value) return;

  changeLocale(langValue);

  ElMessage.success({
    message:
      t("language.switch") +
      " " +
      (langValue === "zh-CN" ? "简体中文" : "English"),
    duration: 2000,
  });
};
</script>

<style scoped>
.language-switcher {
  display: inline-block;
}

.language-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  color: var(--text-primary, #333);
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 14px;
  font-weight: 500;
}

.language-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.3);
  transform: translateY(-1px);
}

.language-icon {
  color: var(--text-primary, #333);
  transition: transform 0.3s ease;
}

.language-btn:hover .language-icon {
  transform: rotate(15deg);
}

.language-text {
  font-size: 18px;
  line-height: 1;
}

/* 下拉菜单项 */
:deep(.el-dropdown-menu__item) {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  transition: all 0.2s ease;
}

:deep(.el-dropdown-menu__item:hover) {
  background: rgba(0, 0, 0, 0.05);
}

:deep(.el-dropdown-menu__item.is-active) {
  background: rgba(64, 158, 255, 0.1);
  color: #409eff;
  font-weight: 600;
}

.lang-flag {
  font-size: 20px;
  line-height: 1;
}

.lang-label {
  flex: 1;
  font-size: 14px;
}

.check-icon {
  color: #409eff;
  margin-left: auto;
}

/* 深色模式 */
@media (prefers-color-scheme: dark) {
  .language-btn {
    background: rgba(255, 255, 255, 0.05);
    border-color: rgba(255, 255, 255, 0.1);
    color: #fff;
  }

  .language-btn:hover {
    background: rgba(255, 255, 255, 0.1);
    border-color: rgba(255, 255, 255, 0.2);
  }

  .language-icon {
    color: #fff;
  }
}

/* 移动端优化 */
@media (max-width: 768px) {
  .language-btn {
    padding: 6px 10px;
  }

  .language-text {
    display: none; /* 移动端只显示图标 */
  }
}
</style>

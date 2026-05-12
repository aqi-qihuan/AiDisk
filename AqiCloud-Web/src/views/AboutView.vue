<template>
  <div class="about-view">
    <!-- ===== HOK Epic Hero ===== -->
    <div class="hero-section">
      <div
        class="hero-glow pink"
        style="width: 600px; height: 600px; top: -200px; right: -100px"
      ></div>
      <div
        class="hero-glow gold"
        style="width: 400px; height: 400px; bottom: -100px; left: -80px"
      ></div>
      <div
        class="hero-glow pink"
        style="width: 300px; height: 300px; top: 30%; left: 50%"
      ></div>
      <div class="hero-grid"></div>

      <div class="hero-content">
        <div class="hero-badge">
          <svg
            width="14"
            height="14"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2.5"
          >
            <path d="M12 2L2 7l10 5 10-5-10-5z" />
            <path d="M2 17l10 5 10-5" />
            <path d="M2 12l10 5 10-5" />
          </svg>
          {{ t("page.about.title") }}
        </div>
        <h1 class="hero-title">
          <span class="gold-gradient">关于小七云盘</span>
        </h1>
        <p class="hero-subtitle">{{ t("page.about.subtitle") }}</p>
      </div>
    </div>

    <!-- ===== 主要内容 ===== -->
    <div class="content-section">
      <div class="content-container">
        <!-- 项目介绍 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box pink">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"
                />
                <polyline points="14 2 14 8 20 8" />
                <line x1="16" y1="13" x2="8" y2="13" />
                <line x1="16" y1="17" x2="8" y2="17" />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.projectTitle") }}
            </h2>
          </div>
          <div class="intro-card">
            <p class="intro-text">{{ t("page.about.projectDesc") }}</p>
          </div>
        </section>

        <!-- 产品亮点 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box gold">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2" />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.highlightsTitle") }}
            </h2>
          </div>
          <div class="highlights-list">
            <div
              class="highlight-card"
              v-for="(item, i) in highlights"
              :key="i"
            >
              <div class="highlight-left">
                <div class="hl-icon-box" :class="item.color">
                  <svg
                    width="26"
                    height="26"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                  >
                    <rect x="2" y="3" width="20" height="14" rx="2" ry="2" />
                    <line x1="8" y1="21" x2="16" y2="21" />
                    <line x1="12" y1="17" x2="12" y2="21" />
                  </svg>
                </div>
              </div>
              <div class="highlight-body">
                <h3>{{ t(item.titleKey) }}</h3>
                <p>{{ t(item.descKey) }}</p>
              </div>
              <div class="highlight-tags">
                <span class="hl-tag" v-for="tag in item.tags" :key="tag">{{
                  tag
                }}</span>
              </div>
            </div>
          </div>
        </section>

        <!-- 核心功能 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box pink">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <polygon
                  points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"
                />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.featuresTitle") }}
            </h2>
          </div>
          <div class="features-grid">
            <div class="feat-card" v-for="(feat, i) in features" :key="i">
              <div class="feat-icon-box" :class="feat.color">
                <svg
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                >
                  <path
                    v-if="feat.icon === 'folder'"
                    d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"
                  />
                  <circle v-if="feat.icon === 'search'" cx="11" cy="11" r="8" />
                  <line
                    v-if="feat.icon === 'search'"
                    x1="21"
                    y1="21"
                    x2="16.65"
                    y2="16.65"
                  />
                  <polyline
                    v-if="feat.icon === 'trash'"
                    points="3 6 5 6 21 6"
                  />
                  <path
                    v-if="feat.icon === 'trash'"
                    d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"
                  />
                  <path
                    v-if="feat.icon === 'link'"
                    d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"
                  />
                  <path
                    v-if="feat.icon === 'link'"
                    d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"
                  />
                  <path
                    v-if="feat.icon === 'users'"
                    d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"
                  />
                  <circle v-if="feat.icon === 'users'" cx="9" cy="7" r="4" />
                  <path
                    v-if="feat.icon === 'users'"
                    d="M23 21v-2a4 4 0 0 0-3-3.87"
                  />
                  <path
                    v-if="feat.icon === 'users'"
                    d="M16 3.13a4 4 0 0 1 0 7.75"
                  />
                </svg>
              </div>
              <h3>{{ t(feat.titleKey) }}</h3>
              <p>{{ t(feat.descKey) }}</p>
            </div>
          </div>
        </section>

        <!-- 快速开始 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box gold">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
                <polyline points="17 8 12 3 7 8" />
                <line x1="12" y1="3" x2="12" y2="15" />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.quickstartTitle") }}
            </h2>
          </div>
          <div class="steps-row">
            <div class="step-card" v-for="(step, i) in steps" :key="i">
              <div class="step-num" :class="step.color">{{ i + 1 }}</div>
              <div class="step-body">
                <h3>{{ t(step.titleKey) }}</h3>
                <p>{{ t(step.descKey) }}</p>
              </div>
              <div class="step-arrow" v-if="i < steps.length - 1">
                <svg
                  width="20"
                  height="20"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2.5"
                >
                  <line x1="5" y1="12" x2="19" y2="12" />
                  <polyline points="12 5 19 12 12 19" />
                </svg>
              </div>
            </div>
          </div>
        </section>

        <!-- 开发进度 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box emerald">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <circle cx="12" cy="12" r="10" />
                <polyline points="12 6 12 12 16 14" />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.progressTitle") }}
            </h2>
          </div>
          <div class="progress-grid">
            <div class="prog-card" v-for="(prog, i) in progresses" :key="i">
              <div class="prog-head">
                <span class="prog-check">
                  <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="3"
                  >
                    <polyline points="20 6 9 17 4 12" />
                  </svg>
                </span>
                <h3>{{ t(prog.titleKey) }}</h3>
                <span class="prog-badge completed">{{
                  t("page.about.completed")
                }}</span>
              </div>
              <div class="prog-tags">
                <span class="prog-tag" v-for="tag in prog.tags" :key="tag">{{
                  tag
                }}</span>
              </div>
            </div>
          </div>
        </section>

        <!-- 技术栈 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box blue">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <polyline points="16 18 22 12 16 6" />
                <polyline points="8 6 2 12 8 18" />
              </svg>
            </div>
            <h2 class="section-title-hok">{{ t("page.about.techTitle") }}</h2>
          </div>
          <div class="tech-grid">
            <div class="tech-card" v-for="(tech, i) in techs" :key="i">
              <div class="tech-icon-box" :class="tech.color">
                <svg
                  width="22"
                  height="22"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                >
                  <circle cx="12" cy="12" r="10" />
                  <circle cx="12" cy="12" r="6" />
                  <circle cx="12" cy="12" r="2" />
                </svg>
              </div>
              <div class="tech-body">
                <h4>{{ t(tech.titleKey) }}</h4>
                <p>{{ t(tech.descKey) }}</p>
              </div>
            </div>
          </div>
        </section>

        <!-- 联系我们 -->
        <section class="hok-section">
          <div class="section-title-row">
            <div class="section-icon-box gold">
              <svg
                width="22"
                height="22"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
                />
              </svg>
            </div>
            <h2 class="section-title-hok">
              {{ t("page.about.contactTitle") }}
            </h2>
          </div>
          <div class="contact-grid">
            <div class="contact-item">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"
                />
                <polyline points="22,6 12,13 2,6" />
              </svg>
              <span>{{ t("page.about.email") }}</span>
            </div>
            <div class="contact-item">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"
                />
                <path
                  d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"
                />
              </svg>
              <span>{{ t("page.about.website") }}</span>
            </div>
            <div class="contact-item">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z" />
                <circle cx="12" cy="10" r="3" />
              </svg>
              <span>{{ t("page.about.location") }}</span>
            </div>
          </div>
        </section>
      </div>
    </div>

    <!-- Footer -->
    <footer class="footer-hok">
      <div class="hok-divider"></div>
      <div class="footer-inner">
        <p class="copyright">
          {{ t("page.about.copyright", { year: new Date().getFullYear() }) }}
        </p>
        <p class="tech-stack">{{ t("page.about.techStack") }}</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const highlights = [
  {
    titleKey: "page.about.highlightCloud",
    descKey: "page.about.highlightCloudDesc",
    color: "emerald",
    tags: ["图片", "视频", "文档", "音乐"],
  },
  {
    titleKey: "page.about.highlightSpeed",
    descKey: "page.about.highlightSpeedDesc",
    color: "pink",
    tags: ["极速上传", "秒传技术", "CDN加速"],
  },
  {
    titleKey: "page.about.highlightSync",
    descKey: "page.about.highlightSyncDesc",
    color: "gold",
    tags: ["多端同步", "实时更新", "离线访问"],
  },
];

const features = [
  {
    titleKey: "page.about.featureFile",
    descKey: "page.about.featureFileDesc",
    icon: "folder",
    color: "pink",
  },
  {
    titleKey: "page.about.featureSearch",
    descKey: "page.about.featureSearchDesc",
    icon: "search",
    color: "blue",
  },
  {
    titleKey: "page.about.featureRecycle",
    descKey: "page.about.featureRecycleDesc",
    icon: "trash",
    color: "emerald",
  },
  {
    titleKey: "page.about.featureShare",
    descKey: "page.about.featureShareDesc",
    icon: "link",
    color: "gold",
  },
  {
    titleKey: "page.about.featureMultiUser",
    descKey: "page.about.featureMultiUserDesc",
    icon: "users",
    color: "pink",
  },
];

const steps = [
  {
    titleKey: "page.about.step1Title",
    descKey: "page.about.step1Desc",
    color: "pink",
  },
  {
    titleKey: "page.about.step2Title",
    descKey: "page.about.step2Desc",
    color: "gold",
  },
  {
    titleKey: "page.about.step3Title",
    descKey: "page.about.step3Desc",
    color: "emerald",
  },
];

const progresses = [
  {
    titleKey: "page.about.progressUser",
    tags: ["用户注册", "用户登录", "用户管理"],
  },
  {
    titleKey: "page.about.progressFile",
    tags: ["文件上传", "文件管理", "文件预览", "文件下载"],
  },
  {
    titleKey: "page.about.progressShare",
    tags: ["创建分享", "分享管理", "文件转存"],
  },
];

const techs = [
  {
    titleKey: "page.about.techFrontend",
    descKey: "page.about.techFrontendDesc",
    color: "blue",
  },
  {
    titleKey: "page.about.techUI",
    descKey: "page.about.techUIDesc",
    color: "pink",
  },
  {
    titleKey: "page.about.techBackend",
    descKey: "page.about.techBackendDesc",
    color: "emerald",
  },
  {
    titleKey: "page.about.techAI",
    descKey: "page.about.techAIDesc",
    color: "gold",
  },
];
</script>

<style>
/* Global CSS variables - must be in unscoped style to apply to :root */
:root {
  --p-primary: #db2777;
  --p-secondary: #f472b6;
  --p-gold: #d97706;
  --p-gold-light: #fbbf24;
  --p-bg-dark: #0b0b10;
  --p-bg-card: #1a1a24;
  --p-bg-surface: #14141c;
  --p-text: #f8fafc;
  --p-text-secondary: #94a3b8;
  --p-text-tertiary: #64748b;
  --p-border: rgba(255, 255, 255, 0.06);
  --p-glass-bg: rgba(255, 255, 255, 0.04);
  --p-glass-hover: rgba(255, 255, 255, 0.08);
  --p-radius: 16px;
  --p-glows: 0 0 24px rgba(219, 39, 119, 0.25);
  --p-glowg: 0 0 24px rgba(217, 119, 6, 0.3);
  --font-display: "Plus Jakarta Sans", "Fira Code", sans-serif;
  --font-body: "Fira Sans", sans-serif;
}
</style>

<style scoped>
/* ===== HOK Design System - About Page ===== */
@import url("https://fonts.googleapis.com/css2?family=Fira+Code:wght@400;500;600;700&family=Fira+Sans:wght@300;400;500;600;700&family=Plus+Jakarta+Sans:wght@400;500;600;700;800&display=swap");

.about-view {
  min-height: 100vh;
  background: var(--p-bg-dark);
  font-family: var(--font-body);
  overflow-x: hidden;
  color: var(--p-text);
}

/* ===== Hero Section ===== */
.hero-section {
  position: relative;
  min-height: 50vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  overflow: hidden;
  padding: 100px 24px;
}

.hero-glow {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  pointer-events: none;
  opacity: 0.4;
}
.hero-glow.pink {
  background: radial-gradient(circle, #db2777, transparent);
}
.hero-glow.gold {
  background: radial-gradient(circle, #d97706, transparent);
}

.hero-grid {
  position: absolute;
  inset: 0;
  background: url("data:image/svg+xml,%3Csvg width='60' height='60' xmlns='http://www.w3.org/2000/svg'%3E%3Cdefs%3E%3Cpattern id='g' width='60' height='60' patternUnits='userSpaceOnUse'%3E%3Cpath d='M 60 0 L 0 0 0 60' fill='none' stroke='rgba(255,255,255,0.02)' stroke-width='0.5'/%3E%3C/pattern%3E%3C/defs%3E%3Crect width='100%25' height='100%25' fill='url(%23g)'/%3E%3C/svg%3E");
  opacity: 0.6;
  z-index: 0;
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 800px;
  animation: fadeUp 0.8s ease-out;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 16px;
  background: rgba(219, 39, 119, 0.1);
  border: 1px solid rgba(219, 39, 119, 0.2);
  border-radius: 999px;
  font-size: 13px;
  color: var(--p-secondary);
  font-weight: 500;
  margin-bottom: 24px;
  backdrop-filter: blur(10px);
}

.hero-title {
  font-family: var(--font-display);
  font-size: 64px;
  font-weight: 800;
  line-height: 1.05;
  margin-bottom: 20px;
  letter-spacing: -0.03em;
}

.gold-gradient {
  background: linear-gradient(135deg, #fbbf24 0%, #fde68a 50%, #f472b6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  filter: drop-shadow(0 0 24px rgba(251, 191, 36, 0.2));
}

.hero-subtitle {
  font-size: 18px;
  color: var(--p-text-secondary);
  line-height: 1.7;
  max-width: 600px;
  margin: 0 auto;
}

/* ===== Content Area ===== */
.content-section {
  padding: 60px 24px 40px;
}

.content-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 40px;
}

/* ===== Section ===== */
.hok-section {
  background: var(--p-bg-card);
  border: 1px solid var(--p-border);
  border-radius: var(--p-radius);
  padding: 32px;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  animation: fadeUp 0.6s ease-out backwards;
}

.hok-section:nth-child(1) {
  animation-delay: 0.05s;
}
.hok-section:nth-child(2) {
  animation-delay: 0.1s;
}
.hok-section:nth-child(3) {
  animation-delay: 0.15s;
}
.hok-section:nth-child(4) {
  animation-delay: 0.2s;
}
.hok-section:nth-child(5) {
  animation-delay: 0.25s;
}
.hok-section:nth-child(6) {
  animation-delay: 0.3s;
}
.hok-section:nth-child(7) {
  animation-delay: 0.35s;
}

.hok-section::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(251, 191, 36, 0.2),
    rgba(219, 39, 119, 0.2),
    transparent
  );
  opacity: 0;
  transition: opacity 0.25s ease;
}

.hok-section:hover {
  border-color: rgba(255, 255, 255, 0.1);
  transform: translateY(-3px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.3);
}

.hok-section:hover::before {
  opacity: 1;
}

/* ===== Section Title ===== */
.section-title-row {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 28px;
}

.section-icon-box {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.section-icon-box.pink {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.15),
    rgba(219, 39, 119, 0.05)
  );
  border: 1px solid rgba(219, 39, 119, 0.2);
  color: var(--p-secondary);
  box-shadow: 0 0 12px rgba(219, 39, 119, 0.08);
}

.section-icon-box.gold {
  background: linear-gradient(
    135deg,
    rgba(217, 119, 6, 0.15),
    rgba(217, 119, 6, 0.05)
  );
  border: 1px solid rgba(217, 119, 6, 0.2);
  color: var(--p-gold-light);
  box-shadow: 0 0 12px rgba(217, 119, 6, 0.08);
}

.section-icon-box.emerald {
  background: linear-gradient(
    135deg,
    rgba(16, 185, 129, 0.15),
    rgba(16, 185, 129, 0.05)
  );
  border: 1px solid rgba(16, 185, 129, 0.2);
  color: #34d399;
  box-shadow: 0 0 12px rgba(16, 185, 129, 0.08);
}

.section-icon-box.blue {
  background: linear-gradient(
    135deg,
    rgba(59, 130, 246, 0.15),
    rgba(59, 130, 246, 0.05)
  );
  border: 1px solid rgba(59, 130, 246, 0.2);
  color: #60a5fa;
  box-shadow: 0 0 12px rgba(59, 130, 246, 0.08);
}

.section-title-hok {
  font-family: var(--font-display);
  font-size: 22px;
  font-weight: 700;
  color: var(--p-text);
  margin: 0;
}

/* ===== Intro ===== */
.intro-card {
  padding: 24px;
  background: rgba(255, 255, 255, 0.02);
  border-radius: 12px;
  border-left: 4px solid #db2777;
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-left-width: 4px;
}

.intro-text {
  font-size: 16px;
  line-height: 1.8;
  color: var(--p-text-secondary);
  margin: 0;
}

/* ===== Highlights ===== */
.highlights-list {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.highlight-card {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 20px 24px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.highlight-card:hover {
  background: rgba(255, 255, 255, 0.04);
  border-color: rgba(255, 255, 255, 0.08);
  transform: translateX(6px);
}

.hl-icon-box {
  width: 52px;
  height: 52px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: transform 0.25s ease;
}

.highlight-card:hover .hl-icon-box {
  transform: scale(1.08) rotate(5deg);
}

.hl-icon-box.emerald {
  background: linear-gradient(
    135deg,
    rgba(16, 185, 129, 0.12),
    rgba(16, 185, 129, 0.04)
  );
  border: 1px solid rgba(16, 185, 129, 0.15);
  color: #34d399;
}

.hl-icon-box.pink {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12),
    rgba(219, 39, 119, 0.04)
  );
  border: 1px solid rgba(219, 39, 119, 0.15);
  color: var(--p-secondary);
}

.hl-icon-box.gold {
  background: linear-gradient(
    135deg,
    rgba(217, 119, 6, 0.12),
    rgba(217, 119, 6, 0.04)
  );
  border: 1px solid rgba(217, 119, 6, 0.15);
  color: var(--p-gold-light);
}

.highlight-body {
  flex: 1;
  min-width: 0;
}

.highlight-body h3 {
  font-family: var(--font-display);
  font-size: 17px;
  font-weight: 700;
  color: var(--p-text);
  margin: 0 0 4px 0;
}

.highlight-body p {
  font-size: 14px;
  color: var(--p-text-secondary);
  margin: 0;
}

.highlight-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
  flex-shrink: 0;
}

.hl-tag {
  display: inline-flex;
  padding: 5px 12px;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 999px;
  font-size: 12px;
  color: var(--p-text-tertiary);
  font-weight: 500;
  transition: all 0.2s ease;
  cursor: pointer;
}

.hl-tag:hover {
  background: rgba(219, 39, 119, 0.1);
  border-color: rgba(219, 39, 119, 0.2);
  color: var(--p-secondary);
}

/* ===== Features ===== */
.features-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 16px;
}

.feat-card {
  text-align: center;
  padding: 24px 14px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.feat-card::after {
  content: "";
  position: absolute;
  bottom: 0;
  left: 15%;
  right: 15%;
  height: 2px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(251, 191, 36, 0.3),
    transparent
  );
  opacity: 0;
  transition: opacity 0.25s ease;
}

.feat-card:hover::after {
  opacity: 1;
}

.feat-card:hover {
  transform: translateY(-6px);
  border-color: rgba(255, 255, 255, 0.1);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.3);
}

.feat-card:hover .feat-icon-box {
  transform: scale(1.1) rotate(8deg);
}

.feat-icon-box {
  width: 56px;
  height: 56px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 14px;
  transition: transform 0.3s ease;
}

.feat-icon-box.pink {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12),
    rgba(219, 39, 119, 0.04)
  );
  border: 1px solid rgba(219, 39, 119, 0.15);
  color: var(--p-secondary);
}

.feat-icon-box.blue {
  background: linear-gradient(
    135deg,
    rgba(59, 130, 246, 0.12),
    rgba(59, 130, 246, 0.04)
  );
  border: 1px solid rgba(59, 130, 246, 0.15);
  color: #60a5fa;
}

.feat-icon-box.emerald {
  background: linear-gradient(
    135deg,
    rgba(16, 185, 129, 0.12),
    rgba(16, 185, 129, 0.04)
  );
  border: 1px solid rgba(16, 185, 129, 0.15);
  color: #34d399;
}

.feat-icon-box.gold {
  background: linear-gradient(
    135deg,
    rgba(217, 119, 6, 0.12),
    rgba(217, 119, 6, 0.04)
  );
  border: 1px solid rgba(217, 119, 6, 0.15);
  color: var(--p-gold-light);
}

.feat-card h3 {
  font-family: var(--font-display);
  font-size: 15px;
  font-weight: 700;
  color: var(--p-text);
  margin-bottom: 6px;
}

.feat-card p {
  font-size: 13px;
  color: var(--p-text-secondary);
  line-height: 1.5;
  margin: 0;
}

/* ===== Steps ===== */
.steps-row {
  display: flex;
  gap: 20px;
  align-items: stretch;
}

.step-card {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 22px 20px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.step-card:hover {
  background: rgba(255, 255, 255, 0.04);
  border-color: rgba(255, 255, 255, 0.1);
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
}

.step-num {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: var(--font-display);
  font-size: 18px;
  font-weight: 800;
  flex-shrink: 0;
}

.step-num.pink {
  background: rgba(219, 39, 119, 0.12);
  border: 1px solid rgba(219, 39, 119, 0.2);
  color: var(--p-secondary);
}

.step-num.gold {
  background: rgba(217, 119, 6, 0.12);
  border: 1px solid rgba(217, 119, 6, 0.2);
  color: var(--p-gold-light);
}

.step-num.emerald {
  background: rgba(16, 185, 129, 0.12);
  border: 1px solid rgba(16, 185, 129, 0.2);
  color: #34d399;
}

.step-body h3 {
  font-family: var(--font-display);
  font-size: 16px;
  font-weight: 700;
  color: var(--p-text);
  margin: 0 0 4px 0;
}

.step-body p {
  font-size: 13px;
  color: var(--p-text-secondary);
  margin: 0;
}

.step-arrow {
  color: var(--p-text-tertiary);
  flex-shrink: 0;
  display: flex;
  align-items: center;
  margin-left: auto;
}

/* ===== Progress ===== */
.progress-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.prog-card {
  padding: 22px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  border-left: 3px solid #10b981;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.prog-card:hover {
  transform: translateY(-4px);
  border-color: rgba(255, 255, 255, 0.08);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
}

.prog-head {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 14px;
}

.prog-check {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(16, 185, 129, 0.15);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #10b981;
  flex-shrink: 0;
}

.prog-head h3 {
  flex: 1;
  font-family: var(--font-display);
  font-size: 16px;
  font-weight: 700;
  color: var(--p-text);
  margin: 0;
}

.prog-badge {
  font-size: 11px;
  padding: 3px 10px;
  border-radius: 999px;
  font-weight: 600;
  letter-spacing: 0.03em;
}

.prog-badge.completed {
  background: rgba(16, 185, 129, 0.12);
  color: #34d399;
  border: 1px solid rgba(16, 185, 129, 0.15);
}

.prog-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.prog-tag {
  font-size: 12px;
  color: var(--p-text-tertiary);
  background: rgba(255, 255, 255, 0.03);
  padding: 5px 12px;
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.04);
  transition: all 0.2s ease;
  cursor: pointer;
}

.prog-tag:hover {
  background: rgba(219, 39, 119, 0.08);
  border-color: rgba(219, 39, 119, 0.15);
  color: var(--p-secondary);
}

/* ===== Tech ===== */
.tech-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.tech-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 18px 20px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.tech-card:hover {
  transform: translateY(-4px);
  border-color: rgba(255, 255, 255, 0.08);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
}

.tech-card:hover .tech-icon-box {
  transform: rotate(360deg) scale(1.1);
}

.tech-icon-box {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.tech-icon-box.blue {
  background: linear-gradient(
    135deg,
    rgba(59, 130, 246, 0.12),
    rgba(59, 130, 246, 0.04)
  );
  border: 1px solid rgba(59, 130, 246, 0.15);
  color: #60a5fa;
}

.tech-icon-box.pink {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12),
    rgba(219, 39, 119, 0.04)
  );
  border: 1px solid rgba(219, 39, 119, 0.15);
  color: var(--p-secondary);
}

.tech-icon-box.emerald {
  background: linear-gradient(
    135deg,
    rgba(16, 185, 129, 0.12),
    rgba(16, 185, 129, 0.04)
  );
  border: 1px solid rgba(16, 185, 129, 0.15);
  color: #34d399;
}

.tech-icon-box.gold {
  background: linear-gradient(
    135deg,
    rgba(217, 119, 6, 0.12),
    rgba(217, 119, 6, 0.04)
  );
  border: 1px solid rgba(217, 119, 6, 0.15);
  color: var(--p-gold-light);
}

.tech-body h4 {
  font-family: var(--font-display);
  font-size: 15px;
  font-weight: 700;
  color: var(--p-text);
  margin: 0 0 3px 0;
}

.tech-body p {
  font-size: 13px;
  color: var(--p-text-secondary);
  margin: 0;
}

/* ===== Contact ===== */
.contact-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 18px 20px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 14px;
  color: var(--p-text-secondary);
  cursor: pointer;
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

.contact-item:hover {
  color: var(--p-text);
  background: rgba(219, 39, 119, 0.04);
  border-color: rgba(219, 39, 119, 0.15);
  transform: translateY(-4px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
}

.contact-item svg {
  flex-shrink: 0;
  color: var(--p-secondary);
  transition: transform 0.25s ease;
}

.contact-item:hover svg {
  transform: scale(1.15) rotate(-5deg);
}

.contact-item span {
  font-size: 14px;
}

/* ===== Footer ===== */
.footer-hok {
  padding: 32px 24px;
}

.hok-divider {
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(251, 191, 36, 0.1),
    transparent
  );
  margin-bottom: 24px;
}

.footer-inner {
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.copyright {
  font-size: 14px;
  color: var(--p-text-tertiary);
  margin-bottom: 6px;
}

.tech-stack {
  font-size: 13px;
  color: var(--p-text-tertiary);
  opacity: 0.7;
}

/* ===== Animations ===== */
@keyframes fadeUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* ===== Responsive ===== */
@media (max-width: 1200px) {
  .features-grid {
    grid-template-columns: repeat(3, 1fr);
  }
  .tech-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 1024px) {
  .progress-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .steps-row {
    flex-direction: column;
  }
  .step-arrow {
    display: none;
  }
}

@media (max-width: 768px) {
  .hero-section {
    min-height: 40vh;
    padding: 60px 16px;
  }
  .hero-title {
    font-size: 40px;
  }
  .hero-subtitle {
    font-size: 16px;
  }
  .content-section {
    padding: 40px 16px;
  }
  .hok-section {
    padding: 24px;
  }
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
  }
  .progress-grid,
  .tech-grid,
  .contact-grid {
    grid-template-columns: 1fr;
  }
  .highlight-card {
    flex-direction: column;
    text-align: center;
  }
  .highlight-tags {
    justify-content: center;
  }
  .section-title-row {
    flex-direction: column;
    text-align: center;
  }
}

@media (max-width: 576px) {
  .hero-title {
    font-size: 32px;
  }
  .features-grid {
    grid-template-columns: 1fr;
  }
  .step-card {
    flex-direction: column;
    text-align: center;
  }
  .contact-grid {
    grid-template-columns: 1fr;
  }
}
</style>

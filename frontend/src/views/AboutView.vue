<template>
  <div class="about-view">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-background">
        <div class="gradient-orb orb-1"></div>
        <div class="gradient-orb orb-2"></div>
        <div class="gradient-orb orb-3"></div>
      </div>
      <div class="header-content">
        <div class="icon-wrapper">
          <el-icon :size="48"><InfoFilled /></el-icon>
        </div>
        <h1 class="page-title">关于我们</h1>
        <p class="page-subtitle">了解小七云盘智能文件管理系统</p>
      </div>
    </div>

    <!-- 主要内容区域 -->
    <div class="content-section">
      <div class="content-container">
        <!-- 项目介绍 -->
        <section class="intro-section glass-section">
          <div class="section-header">
            <div class="section-icon">
              <el-icon><Document /></el-icon>
            </div>
            <h2>项目介绍</h2>
          </div>
          <div class="intro-content">
            <p class="intro-text">
              小七云盘是一款基于 <span class="highlight">MinIO</span> 对象存储构建的现代云存储解决方案，提供安全、高效的文件管理服务。
            </p>
          </div>
        </section>

        <!-- 产品亮点 -->
        <section class="highlights-section glass-section">
          <div class="section-header">
            <div class="section-icon gradient">
              <el-icon><Lightning /></el-icon>
            </div>
            <h2>产品亮点</h2>
          </div>
          <div class="highlights-list">
            <div class="highlight-card glass-card">
              <div class="highlight-icon-wrapper green">
                <el-icon><Box /></el-icon>
              </div>
              <div class="highlight-content">
                <h3>云端存储</h3>
                <p>10GB 免费空间</p>
              </div>
              <div class="highlight-tags">
                <span class="highlight-tag">
                  <el-icon><Picture /></el-icon>
                  图片
                </span>
                <span class="highlight-tag">
                  <el-icon><VideoCamera /></el-icon>
                  视频
                </span>
              </div>
            </div>
            <div class="highlight-card glass-card">
              <div class="highlight-icon-wrapper purple">
                <el-icon><Lightning /></el-icon>
              </div>
              <div class="highlight-content">
                <h3>极速传输</h3>
                <p>秒级上传下载</p>
              </div>
            </div>
            <div class="highlight-card glass-card">
              <div class="highlight-icon-wrapper orange">
                <el-icon><Connection /></el-icon>
              </div>
              <div class="highlight-content">
                <h3>多端同步</h3>
                <p>实时数据同步</p>
              </div>
            </div>
          </div>
        </section>

        <!-- 核心功能 -->
        <section class="features-section glass-section">
          <div class="section-header">
            <div class="section-icon purple">
              <el-icon><Star /></el-icon>
            </div>
            <h2>核心功能</h2>
          </div>
          <div class="features-grid">
            <div class="feature-card glass-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path>
                </svg>
              </div>
              <h3>文件管理</h3>
              <p>上传、下载、预览、分享一站式解决</p>
            </div>
            <div class="feature-card glass-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="11" cy="11" r="8"></circle>
                  <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
                </svg>
              </div>
              <h3>智能搜索</h3>
              <p>快速定位文件，支持全文检索</p>
            </div>
            <div class="feature-card glass-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="3 6 5 6 21 6"></polyline>
                  <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
                  <line x1="10" y1="11" x2="10" y2="17"></line>
                  <line x1="14" y1="11" x2="14" y2="17"></line>
                </svg>
              </div>
              <h3>回收站</h3>
              <p>文件删除与恢复，数据更安全</p>
            </div>
            <div class="feature-card glass-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path>
                  <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path>
                </svg>
              </div>
              <h3>文件分享</h3>
              <p>生成分享链接，权限可控</p>
            </div>
            <div class="feature-card glass-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                  <circle cx="9" cy="7" r="4"></circle>
                  <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                  <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                </svg>
              </div>
              <h3>多用户支持</h3>
              <p>用户注册与管理，权限隔离</p>
            </div>
          </div>
        </section>

        <!-- 快速开始 -->
        <section class="quickstart-section glass-section">
          <div class="section-header">
            <div class="section-icon green">
              <el-icon><Upload /></el-icon>
            </div>
            <h2>快速开始</h2>
          </div>
          <div class="steps-container">
            <div class="step-card glass-card">
              <div class="step-number">1</div>
              <div class="step-content">
                <h3>数据库配置</h3>
                <p>运行 SQL 脚本，创建数据库</p>
              </div>
              <div class="step-arrow">→</div>
            </div>
            <div class="step-card glass-card">
              <div class="step-number">2</div>
              <div class="step-content">
                <h3>服务配置</h3>
                <p>配置数据库和 MinIO 连接信息</p>
              </div>
              <div class="step-arrow">→</div>
            </div>
            <div class="step-card glass-card">
              <div class="step-number">3</div>
              <div class="step-content">
                <h3>启动项目</h3>
                <p>运行项目，开始体验</p>
              </div>
            </div>
          </div>
        </section>

        <!-- 开发进度 -->
        <section class="progress-section glass-section">
          <div class="section-header">
            <div class="section-icon orange">
              <el-icon><Cpu /></el-icon>
            </div>
            <h2>开发进度</h2>
          </div>
          <div class="progress-grid">
            <div class="progress-card glass-card">
              <div class="progress-header">
                <div class="progress-icon completed">
                  <el-icon><Check /></el-icon>
                </div>
                <h3>用户模块</h3>
                <span class="status-badge completed">已完成</span>
              </div>
              <div class="progress-tags">
                <span class="tag">用户注册</span>
                <span class="tag">用户登录</span>
                <span class="tag">用户管理</span>
              </div>
            </div>
            <div class="progress-card glass-card">
              <div class="progress-header">
                <div class="progress-icon completed">
                  <el-icon><Check /></el-icon>
                </div>
                <h3>文件模块</h3>
                <span class="status-badge completed">已完成</span>
              </div>
              <div class="progress-tags">
                <span class="tag">文件上传</span>
                <span class="tag">文件管理</span>
                <span class="tag">文件预览</span>
                <span class="tag">文件下载</span>
              </div>
            </div>
            <div class="progress-card glass-card">
              <div class="progress-header">
                <div class="progress-icon completed">
                  <el-icon><Check /></el-icon>
                </div>
                <h3>分享模块</h3>
                <span class="status-badge completed">已完成</span>
              </div>
              <div class="progress-tags">
                <span class="tag">创建分享</span>
                <span class="tag">分享管理</span>
                <span class="tag">文件转存</span>
              </div>
            </div>
          </div>
        </section>

        <!-- 技术栈 -->
        <section class="tech-section glass-section">
          <div class="section-header">
            <div class="section-icon blue">
              <el-icon><Cpu /></el-icon>
            </div>
            <h2>技术栈</h2>
          </div>
          <div class="tech-grid">
            <div class="tech-card glass-card">
              <div class="tech-icon">
                <svg class="tech-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10"></circle>
                  <circle cx="12" cy="12" r="6"></circle>
                  <circle cx="12" cy="12" r="2"></circle>
                </svg>
              </div>
              <div class="tech-info">
                <h4>前端技术</h4>
                <p>Vue 3 + TypeScript + Vite</p>
              </div>
            </div>
            <div class="tech-card glass-card">
              <div class="tech-icon">
                <svg class="tech-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="3"></circle>
                  <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
                </svg>
              </div>
              <div class="tech-info">
                <h4>UI框架</h4>
                <p>Element Plus + 自定义设计系统</p>
              </div>
            </div>
            <div class="tech-card glass-card">
              <div class="tech-icon">
                <svg class="tech-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09z"></path>
                  <path d="m12 15-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"></path>
                  <path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0"></path>
                  <path d="M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"></path>
                </svg>
              </div>
              <div class="tech-info">
                <h4>后端服务</h4>
                <p>Go (Gin) + MySQL + Redis + GORM</p>
              </div>
            </div>
            <div class="tech-card glass-card">
              <div class="tech-icon">
                <svg class="tech-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M12 8V4H8"></path>
                  <rect width="16" height="12" x="4" y="8" rx="2" ry="2"></rect>
                  <path d="M2 14h2"></path>
                  <path d="M20 14h2"></path>
                  <path d="M15 13v2"></path>
                  <path d="M9 13v2"></path>
                </svg>
              </div>
              <div class="tech-info">
                <h4>AI能力</h4>
                <p>Cloudwego Eino + GPT集成 + 智能文档处理</p>
              </div>
            </div>
          </div>
        </section>

        <!-- 联系我们 -->
        <section class="contact-section glass-section">
          <div class="section-header">
            <div class="section-icon pink">
              <el-icon><Message /></el-icon>
            </div>
            <h2>联系我们</h2>
          </div>
          <div class="contact-grid">
            <div class="contact-item glass-card">
              <el-icon><Message /></el-icon>
              <span>support@aipan.com</span>
            </div>
            <div class="contact-item glass-card">
              <el-icon><Link /></el-icon>
              <span>www.aipan.com</span>
            </div>
            <div class="contact-item glass-card">
              <el-icon><Location /></el-icon>
              <span>中国 · 广州</span>
            </div>
          </div>
        </section>
      </div>
    </div>

    <!-- 页脚版权信息 -->
    <div class="footer-section">
      <div class="footer-content">
        <div class="footer-info">
          <p class="copyright">© {{ new Date().getFullYear() }} 小七云盘. All rights reserved.</p>
          <p class="tech-stack">Made with ❤️ using Vue 3 + TypeScript</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  InfoFilled,
  Document,
  Cpu,
  Star,
  Upload,
  Message,
  Link,
  Location,
  Check,
  Lightning,
  Box,
  Connection,
  Picture,
  VideoCamera
} from '@element-plus/icons-vue';
</script>

<style scoped>
/* 引入 Inter 字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap');

/* CSS 变量定义 */
::root {
  --primary: #DB2777;
  --secondary: #F472B6;
  --cta: #CA8A04;
  --background: #FDF2F8;
  --text-primary: #831843;
  --text-secondary: #475569;
  --font-primary: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  
  /* Glassmorphism 变量 */
  --glass-bg: rgba(255, 255, 255, 0.72);
  --glass-border: rgba(255, 255, 255, 0.5);
  --glass-shadow: 0 8px 32px rgba(0, 0, 0, 0.08);
  --glass-backdrop: blur(20px) saturate(180%);
}

.about-view {
  min-height: 100vh;
  background: linear-gradient(180deg, #FDF2F8 0%, #FCE7F3 50%, #FBCFE8 100%);
  font-family: var(--font-primary);
  overflow-x: hidden;
}

/* ===== 页面头部 ===== */
.page-header {
  position: relative;
  padding: 100px 24px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 50%, #CA8A04 100%);
  overflow: hidden;
}

.header-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.5;
  animation: float 25s infinite ease-in-out;
}

.orb-1 {
  width: 500px;
  height: 500px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.4) 0%, rgba(219, 39, 119, 0.3) 100%);
  top: -200px;
  right: -150px;
  animation-delay: 0s;
}

.orb-2 {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, rgba(244, 114, 182, 0.4) 0%, rgba(202, 138, 4, 0.3) 100%);
  bottom: -150px;
  left: -100px;
  animation-delay: 8s;
}

.orb-3 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, rgba(202, 138, 4, 0.4) 0%, rgba(219, 39, 119, 0.3) 100%);
  top: 40%;
  left: 30%;
  transform: translate(-50%, -50%);
  animation-delay: 16s;
}

.header-content {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
  animation: fadeInUp 0.8s ease-out;
}

.icon-wrapper {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 120px;
  height: 120px;
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 32px;
  margin-bottom: 32px;
  color: white;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  animation: float 6s infinite ease-in-out;
}

.page-title {
  font-size: 3.5rem;
  font-weight: 800;
  font-family: var(--font-primary);
  color: white;
  margin-bottom: 16px;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  letter-spacing: -0.02em;
}

.page-subtitle {
  font-size: 1.25rem;
  font-family: var(--font-primary);
  color: rgba(255, 255, 255, 0.95);
  font-weight: 400;
  letter-spacing: 0.01em;
}

/* ===== 内容区域 ===== */
.content-section {
  padding: 60px 24px;
}

.content-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 48px;
}

/* ===== Glassmorphism Section ===== */
.glass-section {
  background: var(--glass-bg);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid var(--glass-border);
  border-radius: 24px;
  padding: 32px;
  box-shadow: var(--glass-shadow);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: fadeInUp 0.6s ease-out backwards;
}

.glass-section:nth-child(1) { animation-delay: 0.1s; }
.glass-section:nth-child(2) { animation-delay: 0.15s; }
.glass-section:nth-child(3) { animation-delay: 0.2s; }
.glass-section:nth-child(4) { animation-delay: 0.25s; }
.glass-section:nth-child(5) { animation-delay: 0.3s; }
.glass-section:nth-child(6) { animation-delay: 0.35s; }
.glass-section:nth-child(7) { animation-delay: 0.4s; }

.glass-section:hover {
  box-shadow: 0 16px 48px rgba(219, 39, 119, 0.12);
  transform: translateY(-2px);
}

/* ===== Glassmorphism Card ===== */
.glass-card {
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(12px) saturate(150%);
  -webkit-backdrop-filter: blur(12px) saturate(150%);
  border: 1px solid rgba(255, 255, 255, 0.6);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* ===== 区块通用样式 ===== */
.section-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 32px;
}

.section-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}

.section-icon.purple {
  background: linear-gradient(135deg, #8b5cf6 0%, #a855f7 100%);
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
}

.section-icon.green {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.section-icon.orange {
  background: linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.section-icon.blue {
  background: linear-gradient(135deg, #3b82f6 0%, #60a5fa 100%);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.section-icon.pink {
  background: linear-gradient(135deg, #ec4899 0%, #f472b6 100%);
  box-shadow: 0 4px 12px rgba(236, 72, 153, 0.3);
}

.section-icon.gradient {
  background: linear-gradient(135deg, #f59e0b 0%, #ef4444 50%, #8b5cf6 100%);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.section-header h2 {
  font-size: 24px;
  font-weight: 700;
  color: #1e1b4b;
  margin: 0;
}

/* ===== 产品亮点 ===== */
.highlights-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.highlight-card {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 24px;
  border-radius: 16px;
  border: 1px solid rgba(99, 102, 241, 0.1);
  cursor: pointer;
}

.highlight-card:hover {
  transform: translateX(8px) scale(1.01);
  box-shadow: 0 8px 24px rgba(99, 102, 241, 0.15);
  border-color: rgba(99, 102, 241, 0.3);
}

.highlight-icon-wrapper {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  flex-shrink: 0;
  transition: transform 0.3s ease;
}

.highlight-card:hover .highlight-icon-wrapper {
  transform: rotate(8deg) scale(1.1);
}

.highlight-icon-wrapper.green {
  background: linear-gradient(135deg, #10b981 0%, #34d399 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.highlight-icon-wrapper.purple {
  background: linear-gradient(135deg, #8b5cf6 0%, #a855f7 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
}

.highlight-icon-wrapper.orange {
  background: linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.highlight-content {
  flex: 1;
}

.highlight-content h3 {
  font-size: 18px;
  font-weight: 600;
  color: #1e1b4b;
  margin: 0 0 4px 0;
}

.highlight-content p {
  font-size: 14px;
  color: #64748b;
  margin: 0;
}

.highlight-tags {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.highlight-tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 6px 12px;
  background: rgba(255, 255, 255, 0.9);
  border: 1px solid rgba(99, 102, 241, 0.2);
  border-radius: 20px;
  font-size: 13px;
  color: #6366f1;
  font-weight: 500;
  transition: all 0.2s ease;
}

.highlight-tag:hover {
  background: #6366f1;
  color: white;
  transform: scale(1.05);
}

.highlight-tag .el-icon {
  font-size: 14px;
}

/* ===== 项目介绍 ===== */
.intro-content {
  padding: 24px;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 16px;
  border-left: 4px solid #6366f1;
  backdrop-filter: blur(8px);
}

.intro-text {
  font-size: 17px;
  line-height: 1.8;
  color: #475569;
  margin: 0;
}

.intro-text .highlight {
  color: #6366f1;
  font-weight: 600;
  background: linear-gradient(135deg, rgba(99, 102, 241, 0.15) 0%, rgba(139, 92, 246, 0.15) 100%);
  padding: 2px 8px;
  border-radius: 4px;
}

/* ===== 核心功能 ===== */
.features-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 20px;
}

.feature-card {
  text-align: center;
  padding: 24px 16px;
  border-radius: 20px;
  border: 1px solid rgba(99, 102, 241, 0.15);
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(219, 39, 119, 0.05) 0%, rgba(202, 138, 4, 0.05) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card:hover::before {
  opacity: 1;
}

.feature-card:hover {
  transform: translateY(-12px) rotateX(5deg) rotateY(5deg);
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.2);
  border-color: transparent;
}

.feature-card:hover .feature-icon-wrapper {
  transform: scale(1.15) rotate(10deg);
}

.feature-icon-wrapper {
  width: 64px;
  height: 64px;
  margin: 0 auto 16px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 50%, #CA8A04 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  z-index: 1;
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 6px 20px rgba(219, 39, 119, 0.3);
}

/* 渐变边框效果 */
.feature-icon-wrapper::after {
  content: '';
  position: absolute;
  inset: -2px;
  background: linear-gradient(135deg, #DB2777 0%, #CA8A04 100%);
  border-radius: 18px;
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card:hover .feature-icon-wrapper::after {
  opacity: 1;
}

.feature-svg-icon {
  width: 32px;
  height: 32px;
  color: white;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.2));
}

.feature-card h3 {
  font-size: 16px;
  font-weight: 600;
  color: #1e1b4b;
  margin-bottom: 8px;
  position: relative;
  z-index: 1;
}

.feature-card p {
  font-size: 13px;
  color: #64748b;
  line-height: 1.5;
  margin: 0;
  position: relative;
  z-index: 1;
}

/* ===== 快速开始 ===== */
.steps-container {
  display: flex;
  gap: 24px;
  align-items: stretch;
}

.step-card {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 24px;
  border-radius: 20px;
  border: 2px solid rgba(99, 102, 241, 0.2);
  position: relative;
  cursor: pointer;
}

.step-card:hover {
  border-color: #6366f1;
  box-shadow: 0 12px 32px rgba(99, 102, 241, 0.15);
  transform: translateY(-4px);
}

.step-card:hover .step-number {
  animation: rotateBorder 2s linear infinite;
}

.step-number {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 700;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
  position: relative;
  transition: transform 0.3s ease;
}

/* 步骤编号渐变旋转边框 */
.step-number::before {
  content: '';
  position: absolute;
  inset: -3px;
  background: conic-gradient(from 0deg, #6366f1, #a855f7, #ec4899, #6366f1);
  border-radius: 50%;
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.step-card:hover .step-number::before {
  opacity: 1;
}

@keyframes rotateBorder {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.step-content h3 {
  font-size: 16px;
  font-weight: 600;
  color: #1e1b4b;
  margin-bottom: 4px;
}

.step-content p {
  font-size: 14px;
  color: #64748b;
  margin: 0;
}

.step-arrow {
  font-size: 24px;
  color: #cbd5e1;
  margin-left: auto;
}

.step-card:last-child .step-arrow {
  display: none;
}

/* ===== 开发进度 ===== */
.progress-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.progress-card {
  padding: 24px;
  border-radius: 16px;
  border: 1px solid rgba(16, 185, 129, 0.2);
  border-left: 4px solid #10b981;
  cursor: pointer;
}

.progress-card:hover {
  box-shadow: 0 12px 32px rgba(16, 185, 129, 0.12);
  transform: translateY(-4px);
}

.progress-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.progress-icon {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.progress-icon.completed {
  background: #10b981;
  color: white;
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.3);
}

.progress-header h3 {
  flex: 1;
  font-size: 16px;
  font-weight: 600;
  color: #1e1b4b;
  margin: 0;
}

.status-badge {
  font-size: 12px;
  padding: 4px 12px;
  border-radius: 20px;
  font-weight: 500;
}

.status-badge.completed {
  background: rgba(16, 185, 129, 0.15);
  color: #10b981;
}

.progress-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag {
  font-size: 13px;
  color: #64748b;
  background: rgba(255, 255, 255, 0.7);
  padding: 6px 12px;
  border-radius: 8px;
  border: 1px solid rgba(0, 0, 0, 0.05);
  transition: all 0.2s ease;
  cursor: pointer;
}

.tag:hover {
  background: #6366f1;
  color: white;
  transform: scale(1.05);
}

/* ===== 技术栈 ===== */
.tech-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.tech-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px;
  border-radius: 16px;
  border: 1px solid rgba(59, 130, 246, 0.15);
  cursor: pointer;
}

.tech-card:hover {
  transform: translateY(-6px) rotateX(5deg);
  box-shadow: 0 16px 32px rgba(59, 130, 246, 0.15);
  border-color: #3b82f6;
}

.tech-card:hover .tech-icon {
  transform: rotate(360deg) scale(1.2);
}

.tech-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #3b82f6 0%, #60a5fa 100%);
  border-radius: 12px;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.tech-svg-icon {
  width: 24px;
  height: 24px;
  color: white;
}

.tech-info h4 {
  font-size: 15px;
  font-weight: 600;
  color: #1e1b4b;
  margin-bottom: 4px;
}

.tech-info p {
  font-size: 13px;
  color: #64748b;
  margin: 0;
}

/* ===== 联系我们 ===== */
.contact-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.contact-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 20px;
  border-radius: 12px;
  color: #475569;
  border: 1px solid rgba(236, 72, 153, 0.15);
  cursor: pointer;
}

.contact-item:hover {
  color: #ec4899;
  transform: translateY(-4px);
  box-shadow: 0 12px 24px rgba(236, 72, 153, 0.12);
  border-color: #ec4899;
}

.contact-item:hover .el-icon {
  transform: scale(1.2) rotate(-5deg);
}

.contact-item .el-icon {
  font-size: 24px;
  color: #ec4899;
  transition: transform 0.3s ease;
}

/* ===== 页脚 ===== */
.footer-section {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border-top: 1px solid rgba(255, 255, 255, 0.5);
  padding: 32px 24px;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  text-align: center;
}

.copyright {
  font-size: 14px;
  color: #64748b;
  margin-bottom: 8px;
}

.tech-stack {
  font-size: 13px;
  color: #94a3b8;
}

/* ===== 动画 ===== */
@keyframes float {
  0%, 100% {
    transform: translateY(0) scale(1);
  }
  50% {
    transform: translateY(-30px) scale(1.05);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-12px);
  }
}

/* ===== 响应式设计 ===== */
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
  
  .steps-container {
    flex-direction: column;
  }
  
  .step-arrow {
    display: none;
  }
}

@media (max-width: 768px) {
  .page-header {
    padding: 60px 16px;
  }
  
  .content-section {
    padding: 32px 16px;
  }
  
  .page-title {
    font-size: 32px;
  }
  
  .page-subtitle {
    font-size: 16px;
  }
  
  .icon-wrapper {
    width: 80px;
    height: 80px;
  }
  
  .glass-section {
    padding: 24px;
  }
  
  .features-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }
  
  .progress-grid,
  .tech-grid,
  .contact-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 576px) {
  .features-grid {
    grid-template-columns: 1fr;
  }
  
  .step-card {
    flex-direction: column;
    text-align: center;
  }
  
  .section-header {
    flex-direction: column;
    text-align: center;
  }
}
</style>

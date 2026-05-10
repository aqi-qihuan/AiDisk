<template>
  <div id="user-layout">
    <div class="content">
      <!-- 左侧装饰区域 -->
      <div class="left-section">
        <!-- 动态背景装饰 -->
        <div class="hero-bg">
          <div class="gradient-orb orb-1"></div>
          <div class="gradient-orb orb-2"></div>
          <div class="gradient-orb orb-3"></div>
          <div class="grid-pattern"></div>
        </div>

        <div class="hero-content">
          <div class="logo-section">
            <div class="logo-wrapper">
              <img src="../assets/logo.svg" alt="小七云盘" class="logo-icon" />
            </div>
            <span class="logo-text">小七云盘</span>
          </div>

          <div class="hero-main">
            <h1 class="hero-title">
              <span class="title-highlight">安全存储</span>
              <br />随时随地访问
            </h1>
            <p class="hero-desc">新一代智能云存储平台，为您提供安全、便捷、高效的文件管理服务</p>
          </div>

          <div class="hero-features">
            <div class="feature-card">
              <div class="feature-icon-wrapper">
                <svg class="feature-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M18 10h-1.26A8 8 0 1 0 9 20h9a5 5 0 0 0 0-10z"></path>
                </svg>
              </div>
              <div class="feature-info">
                <span class="feature-title">云端存储</span>
                <span class="feature-desc">10GB 免费空间</span>
              </div>
            </div>
            <div class="feature-card">
              <div class="feature-icon-wrapper purple">
                <svg class="feature-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"></polygon>
                </svg>
              </div>
              <div class="feature-info">
                <span class="feature-title">极速传输</span>
                <span class="feature-desc">秒级上传下载</span>
              </div>
            </div>
            <div class="feature-card">
              <div class="feature-icon-wrapper orange">
                <svg class="feature-icon" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path>
                  <rect x="2" y="9" width="4" height="12"></rect>
                  <circle cx="4" cy="4" r="2"></circle>
                </svg>
              </div>
              <div class="feature-info">
                <span class="feature-title">多端同步</span>
                <span class="feature-desc">实时数据同步</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 装饰性插图 -->
        <div class="hero-illustration">
          <div class="floating-card card-1">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="16" y1="13" x2="8" y2="13"></line>
              <line x1="16" y1="17" x2="8" y2="17"></line>
              <polyline points="10 9 9 9 8 9"></polyline>
            </svg>
            <span>文档</span>
          </div>
          <div class="floating-card card-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
              <circle cx="8.5" cy="8.5" r="1.5"></circle>
              <polyline points="21 15 16 10 5 21"></polyline>
            </svg>
            <span>图片</span>
          </div>
          <div class="floating-card card-3">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polygon points="23 7 16 12 23 17 23 7"></polygon>
              <rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect>
            </svg>
            <span>视频</span>
          </div>
          <div class="connection-lines">
            <svg viewBox="0 0 400 300" class="lines-svg">
              <path class="line line-1" d="M50,150 Q200,50 350,100" />
              <path class="line line-2" d="M50,150 Q200,200 350,200" />
              <path class="line line-3" d="M50,150 Q200,250 350,250" />
            </svg>
          </div>
        </div>
      </div>

      <!-- 右侧表单区域 -->
      <div class="right-section">
        <div class="right-content">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * 用户登录/注册页面的布局组件
 * 左侧展示欢迎信息和背景图，右侧显示登录/注册表单
 */
</script>

<style scoped>
#user-layout {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f7fa;
}

.content {
  display: flex;
  width: 100%;
  max-width: 1200px;
  min-height: 700px;
  background-color: #fff;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12);
}

/* 左侧 Hero 区域 */
.left-section {
  flex: 1;
  background: linear-gradient(135deg, #1E1B4B 0%, #312E81 50%, #4338CA 100%);
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 60px;
  position: relative;
  overflow: hidden;
  color: #fff;
}

/* 动态背景 */
.hero-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.5;
  animation: float 20s ease-in-out infinite;
}

.orb-1 {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #6366F1, #8B5CF6);
  top: -100px;
  right: -100px;
  animation-delay: 0s;
}

.orb-2 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, #EC4899, #F472B6);
  bottom: 10%;
  left: -50px;
  animation-delay: -7s;
}

.orb-3 {
  width: 250px;
  height: 250px;
  background: linear-gradient(135deg, #10B981, #34D399);
  bottom: -50px;
  right: 20%;
  animation-delay: -14s;
}

@keyframes float {
  0%, 100% { transform: translate(0, 0) scale(1); }
  25% { transform: translate(30px, -30px) scale(1.05); }
  50% { transform: translate(-20px, 20px) scale(0.95); }
  75% { transform: translate(20px, 30px) scale(1.02); }
}

.grid-pattern {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255,255,255,0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255,255,255,0.03) 1px, transparent 1px);
  background-size: 50px 50px;
}

.hero-content {
  max-width: 420px;
  color: white;
  z-index: 2;
  position: relative;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 48px;
}

.logo-wrapper {
  width: 48px;
  height: 48px;
  background: rgba(255,255,255,0.15);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255,255,255,0.2);
}

.logo-icon {
  width: 28px;
  height: 28px;
  filter: brightness(0) invert(1);
}

.logo-text {
  font-size: 26px;
  font-weight: 700;
  letter-spacing: -0.5px;
}

.hero-main {
  margin-bottom: 40px;
}

.hero-title {
  font-size: 42px;
  font-weight: 800;
  margin: 0 0 16px 0;
  line-height: 1.2;
  letter-spacing: -1px;
}

.title-highlight {
  background: linear-gradient(135deg, #FBBF24, #F59E0B);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-desc {
  font-size: 16px;
  opacity: 0.75;
  margin: 0;
  line-height: 1.6;
  max-width: 360px;
}

.hero-features {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.feature-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 14px 18px;
  background: rgba(255,255,255,0.08);
  border-radius: 12px;
  border: 1px solid rgba(255,255,255,0.1);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.feature-card:hover {
  background: rgba(255,255,255,0.12);
  transform: translateX(8px);
}

.feature-icon-wrapper {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #10B981, #059669);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.feature-icon-wrapper.purple {
  background: linear-gradient(135deg, #8B5CF6, #7C3AED);
}

.feature-icon-wrapper.orange {
  background: linear-gradient(135deg, #F59E0B, #D97706);
}

.feature-icon {
  font-size: 20px;
  color: white;
}

.feature-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.feature-title {
  font-size: 14px;
  font-weight: 600;
}

.feature-desc {
  font-size: 12px;
  opacity: 0.7;
}

/* 装饰插图 */
.hero-illustration {
  position: absolute;
  right: 40px;
  top: 50%;
  transform: translateY(-50%);
  width: 320px;
  height: 350px;
  z-index: 0;
}

.floating-card {
  position: absolute;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  background: rgba(255,255,255,0.12);
  border-radius: 10px;
  border: 1px solid rgba(255,255,255,0.2);
  backdrop-filter: blur(10px);
  color: white;
  font-size: 13px;
  font-weight: 500;
  animation: cardFloat 6s ease-in-out infinite;
}

.floating-card .el-icon {
  font-size: 18px;
}

.card-1 {
  top: 20%;
  left: 0;
  animation-delay: 0s;
}

.card-2 {
  top: 45%;
  right: 0;
  animation-delay: -2s;
}

.card-3 {
  bottom: 20%;
  left: 15%;
  animation-delay: -4s;
}

@keyframes cardFloat {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
}

.connection-lines {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.lines-svg {
  width: 100%;
  height: 100%;
}

.line {
  fill: none;
  stroke: rgba(255,255,255,0.2);
  stroke-width: 2;
  stroke-dasharray: 8, 4;
  animation: dash 20s linear infinite;
}

.line-1 { animation-delay: 0s; }
.line-2 { animation-delay: -7s; }
.line-3 { animation-delay: -14s; }

@keyframes dash {
  to { stroke-dashoffset: -100; }
}

/* 右侧表单区域 */
.right-section {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px;
  background: #F8FAFC;
}

.right-content {
  width: 100%;
  max-width: 400px;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .content {
    max-width: 900px;
  }

  .hero-illustration {
    display: none;
  }
}

@media (max-width: 768px) {
  #user-layout {
    background: #fff;
  }

  .content {
    flex-direction: column;
    box-shadow: none;
    border-radius: 0;
    max-width: 100%;
    min-height: auto;
  }

  .left-section {
    display: none;
  }

  .right-section {
    padding: 40px 24px;
    background: #fff;
  }

  .right-content {
    max-width: 100%;
  }
}
</style>

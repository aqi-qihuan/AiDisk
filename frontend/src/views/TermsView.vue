<template>
  <div class="terms-page">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-background">
        <div class="gradient-orb orb-1"></div>
        <div class="gradient-orb orb-2"></div>
        <div class="gradient-orb orb-3"></div>
      </div>
      <div class="header-content">
        <div class="icon-wrapper">
          <svg class="header-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
            <polyline points="14 2 14 8 20 8"></polyline>
            <line x1="16" y1="13" x2="8" y2="13"></line>
            <line x1="16" y1="17" x2="8" y2="17"></line>
            <polyline points="10 9 9 9 8 9"></polyline>
          </svg>
        </div>
        <h1 class="header-title">使用条款</h1>
        <p class="header-subtitle">Terms of Service</p>
        <div class="header-divider"></div>
      </div>
    </div>

    <!-- 主体内容 -->
    <div class="terms-body">
      <div class="terms-container">

        <!-- 侧边导航 + 内容 -->
        <div class="terms-layout">
          <!-- 目录导航 -->
          <nav class="toc-nav glass-nav">
            <div class="toc-title">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
              </svg>
              目录
            </div>
            <ul class="toc-list">
              <li v-for="(item, index) in tocItems" :key="index" class="toc-item">
                <a :href="'#section-' + (index + 1)" class="toc-link" @click.prevent="scrollToSection(index)">
                  <span class="toc-number">{{ index + 1 }}</span>
                  <span>{{ item }}</span>
                </a>
              </li>
            </ul>
          </nav>

          <!-- 条款内容 -->
          <div class="terms-content">
            <section id="section-1" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">1</div>
                <h2 class="section-title">服务说明</h2>
              </div>
              <p class="section-text">
                AI Pan（以下简称"本平台"）是一款基于人工智能技术的智能云盘服务。通过使用本平台，您可以享受文件存储、智能问答、文档分析、AI 聊天等服务。
              </p>
              <p class="section-text">
                使用本平台即表示您同意遵守本使用条款。如果您不同意这些条款，请勿使用本平台。
              </p>
            </section>

            <section id="section-2" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">2</div>
                <h2 class="section-title">用户账户</h2>
              </div>
              <p class="section-text">
                2.1 您需要注册账户才能使用本平台的全部功能。注册时，您需要提供真实、准确、完整的个人信息。
              </p>
              <p class="section-text">
                2.2 您有责任保护您的账户安全，包括保管好密码。如发现账户被盗用，请立即联系我们。
              </p>
              <p class="section-text">
                2.3 一个账户仅限一人使用，禁止分享账户或密码。
              </p>
            </section>

            <section id="section-3" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">3</div>
                <h2 class="section-title">服务内容</h2>
              </div>
              <p class="section-text">
                3.1 本平台提供的服务包括但不限于：
              </p>
              <ul class="terms-list">
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  文件上传、下载、存储和管理
                </li>
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  AI 智能问答服务
                </li>
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  文档智能分析和摘要生成
                </li>
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  AI 聊天助手
                </li>
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  智能文件搜索
                </li>
                <li>
                  <svg class="list-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  文件分享功能
                </li>
              </ul>
              <p class="section-text">
                3.2 我们保留随时修改、暂停或终止部分或全部服务的权利，恕不另行通知。
              </p>
            </section>

            <section id="section-4" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number warning">4</div>
                <h2 class="section-title">用户行为规范</h2>
              </div>
              <p class="section-text">
                4.1 您同意不会使用本平台进行以下行为：
              </p>
              <ul class="terms-list warning-list">
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  上传、存储或分享违法、侵权、淫秽、暴力等内容
                </li>
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  侵犯他人知识产权、隐私权或其他合法权益
                </li>
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  传播病毒、恶意软件或其他有害代码
                </li>
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  未经授权访问他人账户或系统
                </li>
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  滥用 AI 服务生成违法或有害内容
                </li>
                <li>
                  <svg class="list-icon warning" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="15" y1="9" x2="9" y2="15"></line><line x1="9" y1="9" x2="15" y2="15"></line></svg>
                  干扰或破坏本平台的正常运行
                </li>
              </ul>
            </section>

            <section id="section-5" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">5</div>
                <h2 class="section-title">知识产权</h2>
              </div>
              <p class="section-text">
                5.1 您上传的文件和内容归您所有或您已获得合法授权。您授予本平台有限的权限，以便我们为您提供服务。
              </p>
              <p class="section-text">
                5.2 本平台的所有技术、设计、代码、商标等知识产权归我们所有。
              </p>
              <p class="section-text">
                5.3 AI 生成的内容仅供您个人使用，商业使用需遵守相关法律法规。
              </p>
            </section>

            <section id="section-6" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number info">6</div>
                <h2 class="section-title">服务限制与免责声明</h2>
              </div>
              <div class="disclaimer-box glass-card">
                <svg class="disclaimer-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
                  <line x1="12" y1="9" x2="12" y2="13"></line>
                  <line x1="12" y1="17" x2="12.01" y2="17"></line>
                </svg>
                <p class="section-text" style="margin: 0;">
                  6.1 本平台按"现状"提供服务，不对服务的连续性、安全性、准确性作出任何保证。
                </p>
              </div>
              <p class="section-text">
                6.2 AI 生成的内容仅供参考，不构成专业建议。对于因使用 AI 内容而产生的任何损失，我们不承担责任。
              </p>
              <p class="section-text">
                6.3 因不可抗力、系统维护、网络故障等原因导致的服务中断，我们不承担责任。
              </p>
            </section>

            <section id="section-7" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">7</div>
                <h2 class="section-title">账户终止</h2>
              </div>
              <p class="section-text">
                7.1 如您违反本使用条款，我们有权暂停或终止您的账户。
              </p>
              <p class="section-text">
                7.2 您也可以随时注销账户。账户注销后，您的数据将在合理期限内被删除。
              </p>
            </section>

            <section id="section-8" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">8</div>
                <h2 class="section-title">条款修改</h2>
              </div>
              <p class="section-text">
                我们保留随时修改本使用条款的权利。修改后的条款将在本页面公布，继续使用本平台即表示您接受修改后的条款。
              </p>
            </section>

            <section id="section-9" class="terms-section glass-section">
              <div class="section-header">
                <div class="section-number">9</div>
                <h2 class="section-title">联系我们</h2>
              </div>
              <p class="section-text">
                如您对本使用条款有任何疑问，请通过以下方式联系我们：
              </p>
              <div class="contact-cards">
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
                      <polyline points="22,6 12,13 2,6"></polyline>
                    </svg>
                  </div>
                  <div>
                    <div class="contact-label">邮箱</div>
                    <div class="contact-value">support@aipan.com</div>
                  </div>
                </div>
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path>
                      <circle cx="12" cy="10" r="3"></circle>
                    </svg>
                  </div>
                  <div>
                    <div class="contact-label">地址</div>
                    <div class="contact-value">北京市海淀区xxx路xxx号</div>
                  </div>
                </div>
              </div>
            </section>

            <!-- 底部信息 -->
            <div class="terms-footer">
              <div class="effective-date glass-card">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                <span>本条款最后更新日期：2024年1月1日</span>
              </div>
              <router-link to="/" class="back-link">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <line x1="19" y1="12" x2="5" y2="12"></line>
                  <polyline points="12 19 5 12 12 5"></polyline>
                </svg>
                返回首页
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const tocItems = [
  '服务说明',
  '用户账户',
  '服务内容',
  '用户行为规范',
  '知识产权',
  '服务限制与免责声明',
  '账户终止',
  '条款修改',
  '联系我们'
];

const scrollToSection = (index: number) => {
  const el = document.getElementById(`section-${index + 1}`);
  if (el) {
    el.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }
};
</script>

<style scoped>
/* 引入 Inter 字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap');

/* CSS 变量 */
::root {
  --primary: #DB2777;
  --secondary: #F472B6;
  --cta: #CA8A04;
  --background: #FDF2F8;
  --text-primary: #831843;
  --text-secondary: #475569;
  --font-primary: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;

  --glass-bg: rgba(255, 255, 255, 0.72);
  --glass-border: rgba(255, 255, 255, 0.5);
  --glass-shadow: 0 8px 32px rgba(0, 0, 0, 0.08);
  --glass-backdrop: blur(20px) saturate(180%);
}

.terms-page {
  min-height: 100vh;
  background: linear-gradient(180deg, #FDF2F8 0%, #FCE7F3 50%, #FBCFE8 100%);
  font-family: var(--font-primary);
  overflow-x: hidden;
}

/* ===== 页面头部 ===== */
.page-header {
  position: relative;
  padding: 80px 24px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 50%, #CA8A04 100%);
  overflow: hidden;
}

.header-background {
  position: absolute;
  inset: 0;
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
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.4) 0%, rgba(219, 39, 119, 0.3) 100%);
  top: -150px;
  right: -100px;
}

.orb-2 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, rgba(244, 114, 182, 0.4) 0%, rgba(202, 138, 4, 0.3) 100%);
  bottom: -100px;
  left: -80px;
  animation-delay: 8s;
}

.orb-3 {
  width: 250px;
  height: 250px;
  background: linear-gradient(135deg, rgba(202, 138, 4, 0.4) 0%, rgba(219, 39, 119, 0.3) 100%);
  top: 40%;
  left: 35%;
  transform: translate(-50%, -50%);
  animation-delay: 16s;
}

.header-content {
  position: relative;
  z-index: 1;
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  animation: fadeInUp 0.8s ease-out;
}

.icon-wrapper {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 100px;
  height: 100px;
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 28px;
  margin-bottom: 24px;
  color: white;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  animation: float 6s infinite ease-in-out;
}

.header-svg-icon {
  width: 48px;
  height: 48px;
  filter: drop-shadow(0 2px 8px rgba(0, 0, 0, 0.15));
}

.header-title {
  font-size: 3rem;
  font-weight: 800;
  color: white;
  margin: 0 0 8px;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  letter-spacing: -0.02em;
}

.header-subtitle {
  font-size: 1.125rem;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500;
  letter-spacing: 0.05em;
  text-transform: uppercase;
}

.header-divider {
  width: 60px;
  height: 4px;
  background: rgba(255, 255, 255, 0.4);
  border-radius: 2px;
  margin: 20px auto 0;
}

/* ===== 主体区域 ===== */
.terms-body {
  padding: 48px 24px 60px;
}

.terms-container {
  max-width: 1100px;
  margin: 0 auto;
}

/* ===== 布局：目录 + 内容 ===== */
.terms-layout {
  display: flex;
  gap: 32px;
  align-items: flex-start;
}

/* ===== 目录导航 ===== */
.toc-nav {
  position: sticky;
  top: 88px;
  width: 220px;
  flex-shrink: 0;
  background: var(--glass-bg);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid var(--glass-border);
  border-radius: 20px;
  padding: 24px;
  box-shadow: var(--glass-shadow);
  animation: fadeInLeft 0.6s ease-out;
}

.toc-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 700;
  color: #1e1b4b;
  margin-bottom: 16px;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.toc-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.toc-item {
  margin-bottom: 4px;
}

.toc-link {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 12px;
  border-radius: 10px;
  font-size: 13px;
  font-weight: 500;
  color: #64748b;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.25s ease;
}

.toc-link:hover {
  background: rgba(219, 39, 119, 0.08);
  color: #DB2777;
  transform: translateX(4px);
}

.toc-number {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(219, 39, 119, 0.1);
  color: #DB2777;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 700;
  flex-shrink: 0;
  transition: all 0.25s ease;
}

.toc-link:hover .toc-number {
  background: #DB2777;
  color: white;
  transform: rotate(10deg) scale(1.1);
}

/* ===== 条款内容 ===== */
.terms-content {
  flex: 1;
  min-width: 0;
}

/* ===== Glassmorphism Section ===== */
.glass-section {
  background: var(--glass-bg);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid var(--glass-border);
  border-radius: 20px;
  padding: 28px 32px;
  box-shadow: var(--glass-shadow);
  margin-bottom: 20px;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: fadeInUp 0.5s ease-out backwards;
}

.glass-section:nth-child(1) { animation-delay: 0.05s; }
.glass-section:nth-child(2) { animation-delay: 0.1s; }
.glass-section:nth-child(3) { animation-delay: 0.15s; }
.glass-section:nth-child(4) { animation-delay: 0.2s; }
.glass-section:nth-child(5) { animation-delay: 0.25s; }
.glass-section:nth-child(6) { animation-delay: 0.3s; }
.glass-section:nth-child(7) { animation-delay: 0.35s; }
.glass-section:nth-child(8) { animation-delay: 0.4s; }
.glass-section:nth-child(9) { animation-delay: 0.45s; }

.glass-section:hover {
  box-shadow: 0 12px 40px rgba(219, 39, 119, 0.1);
  transform: translateY(-2px);
}

.glass-card {
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(12px) saturate(150%);
  -webkit-backdrop-filter: blur(12px) saturate(150%);
  border: 1px solid rgba(255, 255, 255, 0.6);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* ===== Section Header ===== */
.section-header {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 20px;
}

.section-number {
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 100%);
  color: white;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 800;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
  transition: transform 0.3s ease;
}

.glass-section:hover .section-number {
  transform: rotate(8deg) scale(1.1);
}

.section-number.warning {
  background: linear-gradient(135deg, #ef4444 0%, #f87171 100%);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.section-number.info {
  background: linear-gradient(135deg, #f59e0b 0%, #fbbf24 100%);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.section-title {
  font-size: 20px;
  font-weight: 700;
  color: #1e1b4b;
  margin: 0;
}

/* ===== 文本内容 ===== */
.section-text {
  font-size: 15px;
  color: #475569;
  line-height: 1.85;
  margin: 0 0 12px;
}

.section-text:last-child {
  margin-bottom: 0;
}

/* ===== 列表 ===== */
.terms-list {
  list-style: none;
  padding: 0;
  margin: 16px 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.terms-list li {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  font-size: 15px;
  color: #475569;
  line-height: 1.7;
  padding: 8px 14px;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.6);
  transition: all 0.2s ease;
}

.terms-list li:hover {
  background: rgba(255, 255, 255, 0.8);
  transform: translateX(4px);
}

.list-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
  margin-top: 2px;
  color: #10b981;
}

.list-icon.warning {
  color: #ef4444;
}

/* ===== 免责声明框 ===== */
.disclaimer-box {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px 20px;
  border-radius: 14px;
  border: 1px solid rgba(245, 158, 11, 0.2);
  background: rgba(245, 158, 11, 0.06);
  margin-bottom: 16px;
  cursor: default;
}

.disclaimer-box:hover {
  border-color: rgba(245, 158, 11, 0.4);
  box-shadow: 0 4px 16px rgba(245, 158, 11, 0.1);
}

.disclaimer-icon {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
  margin-top: 1px;
  color: #f59e0b;
}

/* ===== 联系卡片 ===== */
.contact-cards {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  margin-top: 16px;
}

.contact-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 20px;
  border-radius: 14px;
  border: 1px solid rgba(219, 39, 119, 0.15);
  cursor: pointer;
}

.contact-card:hover {
  border-color: #DB2777;
  box-shadow: 0 8px 24px rgba(219, 39, 119, 0.12);
  transform: translateY(-4px);
}

.contact-card:hover .contact-icon {
  transform: scale(1.15) rotate(-5deg);
}

.contact-icon {
  width: 44px;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 100%);
  border-radius: 12px;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
  transition: transform 0.3s ease;
}

.contact-icon svg {
  width: 22px;
  height: 22px;
  color: white;
}

.contact-label {
  font-size: 12px;
  font-weight: 600;
  color: #94a3b8;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 2px;
}

.contact-value {
  font-size: 14px;
  font-weight: 600;
  color: #1e1b4b;
}

/* ===== 底部 ===== */
.terms-footer {
  margin-top: 40px;
  padding-top: 28px;
  border-top: 1px solid rgba(219, 39, 119, 0.15);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
  animation: fadeInUp 0.6s ease-out 0.5s backwards;
}

.effective-date {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 12px;
  font-size: 13px;
  color: #64748b;
  font-weight: 500;
}

.effective-date svg {
  color: #DB2777;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 100%);
  color: white;
  text-decoration: none;
  border-radius: 14px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  box-shadow: 0 4px 16px rgba(219, 39, 119, 0.3);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.back-link:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 28px rgba(219, 39, 119, 0.4);
}

/* ===== 动画 ===== */
@keyframes float {
  0%, 100% { transform: translateY(0) scale(1); }
  50% { transform: translateY(-20px) scale(1.03); }
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fadeInLeft {
  from { opacity: 0; transform: translateX(-20px); }
  to { opacity: 1; transform: translateX(0); }
}

/* ===== 响应式 ===== */
@media (max-width: 1024px) {
  .toc-nav {
    display: none;
  }
}

@media (max-width: 768px) {
  .page-header {
    padding: 60px 16px;
  }

  .header-title {
    font-size: 2.25rem;
  }

  .icon-wrapper {
    width: 80px;
    height: 80px;
    border-radius: 22px;
  }

  .terms-body {
    padding: 32px 16px 48px;
  }

  .glass-section {
    padding: 20px;
  }

  .contact-cards {
    grid-template-columns: 1fr;
  }

  .terms-footer {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
}

@media (max-width: 576px) {
  .header-title {
    font-size: 1.875rem;
  }

  .section-title {
    font-size: 17px;
  }

  .section-number {
    width: 30px;
    height: 30px;
    font-size: 13px;
    border-radius: 10px;
  }
}
</style>

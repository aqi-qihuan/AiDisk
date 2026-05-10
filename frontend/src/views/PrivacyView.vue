<template>
  <div class="privacy-page">
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
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
        </div>
        <h1 class="header-title">隐私政策</h1>
        <p class="header-subtitle">Privacy Policy</p>
        <div class="header-divider"></div>
      </div>
    </div>

    <!-- 主体内容 -->
    <div class="privacy-body">
      <div class="privacy-container">
        <div class="privacy-layout">
          <!-- 目录导航 -->
          <nav class="toc-nav">
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

          <!-- 隐私内容 -->
          <div class="privacy-content">

            <section id="section-1" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">1</div>
                <h2 class="section-title">引言</h2>
              </div>
              <p class="section-text">
                AI Pan（以下简称"本平台"）非常重视您的隐私保护。本隐私政策旨在向您说明我们如何收集、使用、存储和保护您的个人信息。
              </p>
              <p class="section-text">
                使用本平台即表示您同意本隐私政策中描述的做法。请仔细阅读本政策，了解我们对您个人信息的处理方式。
              </p>
            </section>

            <section id="section-2" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">2</div>
                <h2 class="section-title">信息收集</h2>
              </div>
              <p class="section-text">2.1 我们收集的信息类型：</p>
              <ul class="privacy-list">
                <li>
                  <div class="list-icon-box blue">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                  </div>
                  <div><strong>账户信息</strong>：用户名、邮箱地址、手机号码、密码等</div>
                </li>
                <li>
                  <div class="list-icon-box purple">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path><circle cx="12" cy="7" r="4"></circle></svg>
                  </div>
                  <div><strong>个人资料</strong>：头像、昵称等您主动填写的信息</div>
                </li>
                <li>
                  <div class="list-icon-box green">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path></svg>
                  </div>
                  <div><strong>文件数据</strong>：您上传、存储的文件内容</div>
                </li>
                <li>
                  <div class="list-icon-box orange">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="22 12 18 12 15 21 9 3 6 12 2 12"></polyline></svg>
                  </div>
                  <div><strong>使用数据</strong>：访问日志、操作记录、设备信息等</div>
                </li>
                <li>
                  <div class="list-icon-box pink">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path></svg>
                  </div>
                  <div><strong>AI 交互数据</strong>：您与 AI 助手的对话记录</div>
                </li>
              </ul>
              <p class="section-text">
                2.2 我们使用 Cookie 和类似技术来收集有关您使用本平台的信息，以改善用户体验。
              </p>
            </section>

            <section id="section-3" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">3</div>
                <h2 class="section-title">信息使用</h2>
              </div>
              <p class="section-text">3.1 我们使用您的信息用于以下目的：</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  提供、维护和改进本平台服务
                </li>
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  处理您的文件存储和分享请求
                </li>
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  提供 AI 智能问答、文档分析等服务
                </li>
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  向您发送服务通知和重要更新
                </li>
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  防止欺诈和滥用行为
                </li>
                <li>
                  <svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  遵守法律法规要求
                </li>
              </ul>
              <p class="section-text">
                3.2 AI 服务使用：您的文件内容和对话记录将用于提供 AI 分析服务。我们采用严格的安全措施保护这些数据。
              </p>
            </section>

            <section id="section-4" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number shield">4</div>
                <h2 class="section-title">信息存储与安全</h2>
              </div>
              <div class="highlight-box security">
                <div class="highlight-box-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path></svg>
                </div>
                <p class="section-text" style="margin: 0;">4.1 数据存储：您的数据存储在安全的服务器上，采用加密技术保护。</p>
              </div>
              <p class="section-text">4.2 安全措施：</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg class="check-icon shield-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  使用 SSL/TLS 加密传输数据
                </li>
                <li>
                  <svg class="check-icon shield-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  采用 AES-256 加密存储敏感数据
                </li>
                <li>
                  <svg class="check-icon shield-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  定期进行安全审计和漏洞扫描
                </li>
                <li>
                  <svg class="check-icon shield-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                  严格的访问控制和权限管理
                </li>
              </ul>
              <p class="section-text">
                4.3 尽管我们采取合理措施保护您的信息，但互联网传输无法保证 100% 安全。
              </p>
            </section>

            <section id="section-5" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number info">5</div>
                <h2 class="section-title">信息共享</h2>
              </div>
              <div class="highlight-box success">
                <div class="highlight-box-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path><polyline points="9 12 11 14 15 10"></polyline></svg>
                </div>
                <p class="section-text" style="margin: 0;">5.1 我们不会出售您的个人信息。</p>
              </div>
              <p class="section-text">5.2 在以下情况下，我们可能会共享您的信息：</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg class="check-icon info-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                  <div><strong>服务提供商</strong>：帮助我们运营业务的第三方服务商</div>
                </li>
                <li>
                  <svg class="check-icon info-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                  <div><strong>法律要求</strong>：响应法律程序或政府要求</div>
                </li>
                <li>
                  <svg class="check-icon info-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                  <div><strong>保护权利</strong>：保护本平台、用户或公众的权利和安全</div>
                </li>
                <li>
                  <svg class="check-icon info-check" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="16" x2="12" y2="12"></line><line x1="12" y1="8" x2="12.01" y2="8"></line></svg>
                  <div><strong>业务转让</strong>：合并、收购或资产出售时</div>
                </li>
              </ul>
            </section>

            <section id="section-6" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number rights">6</div>
                <h2 class="section-title">您的权利</h2>
              </div>
              <p class="section-text">根据适用法律，您拥有以下权利：</p>
              <div class="rights-grid">
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">访问权</div>
                    <div class="right-desc">获取我们持有的您的个人信息</div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">更正权</div>
                    <div class="right-desc">更正不准确或不完整的信息</div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">删除权</div>
                    <div class="right-desc">要求删除您的个人信息</div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">限制处理权</div>
                    <div class="right-desc">限制我们对您信息的处理</div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path><polyline points="7 10 12 15 17 10"></polyline><line x1="12" y1="15" x2="12" y2="3"></line></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">数据可携带权</div>
                    <div class="right-desc">获取您的数据副本</div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="4.93" y1="4.93" x2="19.07" y2="19.07"></line></svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">反对权</div>
                    <div class="right-desc">反对某些类型的数据处理</div>
                  </div>
                </div>
              </div>
              <p class="section-text">如需行使这些权利，请通过下方联系方式与我们联系。</p>
            </section>

            <section id="section-7" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">7</div>
                <h2 class="section-title">数据保留</h2>
              </div>
              <p class="section-text">
                7.1 我们将在实现本隐私政策所述目的所需的时间内保留您的个人信息。
              </p>
              <p class="section-text">
                7.2 账户注销后，我们将在合理期限内删除您的个人数据，但法律法规要求保留的除外。
              </p>
            </section>

            <section id="section-8" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number warning">8</div>
                <h2 class="section-title">儿童隐私</h2>
              </div>
              <div class="highlight-box warning">
                <div class="highlight-box-icon">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
                </div>
                <p class="section-text" style="margin: 0;">本平台不面向 14 岁以下儿童。如果我们发现收集了 14 岁以下儿童的个人信息，将立即删除。</p>
              </div>
            </section>

            <section id="section-9" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">9</div>
                <h2 class="section-title">政策更新</h2>
              </div>
              <p class="section-text">
                我们可能会不时更新本隐私政策。更新后的政策将在本页面发布，重大变更我们会通过邮件或平台通知您。
              </p>
            </section>

            <section id="section-10" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">10</div>
                <h2 class="section-title">联系我们</h2>
              </div>
              <p class="section-text">
                如您对本隐私政策有任何疑问或 concerns，请通过以下方式联系我们：
              </p>
              <div class="contact-cards">
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>
                  </div>
                  <div>
                    <div class="contact-label">邮箱</div>
                    <div class="contact-value">privacy@aipan.com</div>
                  </div>
                </div>
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path><circle cx="12" cy="10" r="3"></circle></svg>
                  </div>
                  <div>
                    <div class="contact-label">地址</div>
                    <div class="contact-value">北京市海淀区xxx路xxx号</div>
                  </div>
                </div>
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path></svg>
                  </div>
                  <div>
                    <div class="contact-label">电话</div>
                    <div class="contact-value">400-xxx-xxxx</div>
                  </div>
                </div>
              </div>
            </section>

            <!-- 底部 -->
            <div class="privacy-footer">
              <div class="effective-date glass-card">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
                <span>本政策最后更新日期：2024年1月1日</span>
              </div>
              <router-link to="/" class="back-link">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="19" y1="12" x2="5" y2="12"></line><polyline points="12 19 5 12 12 5"></polyline></svg>
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
  '引言',
  '信息收集',
  '信息使用',
  '信息存储与安全',
  '信息共享',
  '您的权利',
  '数据保留',
  '儿童隐私',
  '政策更新',
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
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap');

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

.privacy-page {
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

.header-background { position: absolute; inset: 0; overflow: hidden; }

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.5;
  animation: float 25s infinite ease-in-out;
}

.orb-1 {
  width: 400px; height: 400px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.4), rgba(219, 39, 119, 0.3));
  top: -150px; right: -100px;
}

.orb-2 {
  width: 300px; height: 300px;
  background: linear-gradient(135deg, rgba(244, 114, 182, 0.4), rgba(202, 138, 4, 0.3));
  bottom: -100px; left: -80px;
  animation-delay: 8s;
}

.orb-3 {
  width: 250px; height: 250px;
  background: linear-gradient(135deg, rgba(202, 138, 4, 0.4), rgba(219, 39, 119, 0.3));
  top: 40%; left: 35%; transform: translate(-50%, -50%);
  animation-delay: 16s;
}

.header-content {
  position: relative; z-index: 1;
  max-width: 800px; margin: 0 auto; text-align: center;
  animation: fadeInUp 0.8s ease-out;
}

.icon-wrapper {
  display: inline-flex; align-items: center; justify-content: center;
  width: 100px; height: 100px;
  background: rgba(255, 255, 255, 0.25);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.4);
  border-radius: 28px; margin-bottom: 24px;
  color: white;
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  animation: float 6s infinite ease-in-out;
}

.header-svg-icon {
  width: 48px; height: 48px;
  filter: drop-shadow(0 2px 8px rgba(0, 0, 0, 0.15));
}

.header-title {
  font-size: 3rem; font-weight: 800; color: white;
  margin: 0 0 8px; text-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
  letter-spacing: -0.02em;
}

.header-subtitle {
  font-size: 1.125rem; color: rgba(255, 255, 255, 0.9);
  font-weight: 500; letter-spacing: 0.05em; text-transform: uppercase;
}

.header-divider {
  width: 60px; height: 4px;
  background: rgba(255, 255, 255, 0.4);
  border-radius: 2px; margin: 20px auto 0;
}

/* ===== 主体 ===== */
.privacy-body { padding: 48px 24px 60px; }
.privacy-container { max-width: 1100px; margin: 0 auto; }

.privacy-layout {
  display: flex; gap: 32px; align-items: flex-start;
}

/* ===== 目录导航 ===== */
.toc-nav {
  position: sticky; top: 88px; width: 220px; flex-shrink: 0;
  background: var(--glass-bg);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid var(--glass-border);
  border-radius: 20px; padding: 24px;
  box-shadow: var(--glass-shadow);
  animation: fadeInLeft 0.6s ease-out;
}

.toc-title {
  display: flex; align-items: center; gap: 8px;
  font-size: 14px; font-weight: 700; color: #1e1b4b;
  margin-bottom: 16px; text-transform: uppercase; letter-spacing: 0.05em;
}

.toc-list { list-style: none; padding: 0; margin: 0; }
.toc-item { margin-bottom: 4px; }

.toc-link {
  display: flex; align-items: center; gap: 10px;
  padding: 8px 12px; border-radius: 10px;
  font-size: 13px; font-weight: 500; color: #64748b;
  text-decoration: none; cursor: pointer;
  transition: all 0.25s ease;
}

.toc-link:hover {
  background: rgba(219, 39, 119, 0.08);
  color: #DB2777; transform: translateX(4px);
}

.toc-number {
  width: 24px; height: 24px;
  display: flex; align-items: center; justify-content: center;
  background: rgba(219, 39, 119, 0.1); color: #DB2777;
  border-radius: 8px; font-size: 12px; font-weight: 700;
  flex-shrink: 0; transition: all 0.25s ease;
}

.toc-link:hover .toc-number {
  background: #DB2777; color: white;
  transform: rotate(10deg) scale(1.1);
}

/* ===== Glassmorphism ===== */
.glass-section {
  background: var(--glass-bg);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid var(--glass-border);
  border-radius: 20px; padding: 28px 32px;
  box-shadow: var(--glass-shadow); margin-bottom: 20px;
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
.glass-section:nth-child(10) { animation-delay: 0.5s; }

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
  display: flex; align-items: center;
  gap: 14px; margin-bottom: 20px;
}

.section-number {
  width: 36px; height: 36px;
  display: flex; align-items: center; justify-content: center;
  background: linear-gradient(135deg, #DB2777, #F472B6);
  color: white; border-radius: 12px;
  font-size: 15px; font-weight: 800; flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
  transition: transform 0.3s ease;
}

.glass-section:hover .section-number {
  transform: rotate(8deg) scale(1.1);
}

.section-number.shield {
  background: linear-gradient(135deg, #3b82f6, #60a5fa);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.section-number.info {
  background: linear-gradient(135deg, #f59e0b, #fbbf24);
  box-shadow: 0 4px 12px rgba(245, 158, 11, 0.3);
}

.section-number.warning {
  background: linear-gradient(135deg, #ef4444, #f87171);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.section-number.rights {
  background: linear-gradient(135deg, #8b5cf6, #a855f7);
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
}

.section-title {
  font-size: 20px; font-weight: 700; color: #1e1b4b; margin: 0;
}

/* ===== 文本 ===== */
.section-text {
  font-size: 15px; color: #475569; line-height: 1.85;
  margin: 0 0 12px;
}

.section-text:last-child { margin-bottom: 0; }

/* ===== 列表 ===== */
.privacy-list {
  list-style: none; padding: 0; margin: 16px 0;
  display: flex; flex-direction: column; gap: 10px;
}

.privacy-list li {
  display: flex; align-items: flex-start; gap: 12px;
  font-size: 15px; color: #475569; line-height: 1.7;
  padding: 10px 14px;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.6);
  transition: all 0.2s ease; cursor: default;
}

.privacy-list li:hover {
  background: rgba(255, 255, 255, 0.85);
  transform: translateX(4px);
}

.privacy-list li strong { color: #1e1b4b; font-weight: 600; }

.privacy-list.simple-list li { gap: 10px; padding: 8px 14px; border-radius: 10px; }

.list-icon-box {
  width: 36px; height: 36px; border-radius: 10px;
  display: flex; align-items: center; justify-content: center;
  flex-shrink: 0; transition: transform 0.3s ease;
}

.privacy-list li:hover .list-icon-box {
  transform: rotate(8deg) scale(1.1);
}

.list-icon-box svg { width: 20px; height: 20px; color: white; }

.list-icon-box.blue { background: linear-gradient(135deg, #3b82f6, #60a5fa); box-shadow: 0 4px 8px rgba(59, 130, 246, 0.25); }
.list-icon-box.purple { background: linear-gradient(135deg, #8b5cf6, #a855f7); box-shadow: 0 4px 8px rgba(139, 92, 246, 0.25); }
.list-icon-box.green { background: linear-gradient(135deg, #10b981, #34d399); box-shadow: 0 4px 8px rgba(16, 185, 129, 0.25); }
.list-icon-box.orange { background: linear-gradient(135deg, #f59e0b, #fbbf24); box-shadow: 0 4px 8px rgba(245, 158, 11, 0.25); }
.list-icon-box.pink { background: linear-gradient(135deg, #ec4899, #f472b6); box-shadow: 0 4px 8px rgba(236, 72, 153, 0.25); }

.check-icon { width: 20px; height: 20px; flex-shrink: 0; margin-top: 2px; color: #10b981; }
.check-icon.shield-check { color: #3b82f6; }
.check-icon.info-check { color: #f59e0b; }

/* ===== 高亮提示框 ===== */
.highlight-box {
  display: flex; align-items: flex-start; gap: 12px;
  padding: 16px 20px; border-radius: 14px;
  margin-bottom: 16px; cursor: default; transition: all 0.3s ease;
}

.highlight-box.security {
  border: 1px solid rgba(59, 130, 246, 0.2);
  background: rgba(59, 130, 246, 0.06);
}
.highlight-box.security .highlight-box-icon { color: #3b82f6; }
.highlight-box.security:hover {
  border-color: rgba(59, 130, 246, 0.4);
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.1);
}

.highlight-box.success {
  border: 1px solid rgba(16, 185, 129, 0.2);
  background: rgba(16, 185, 129, 0.06);
}
.highlight-box.success .highlight-box-icon { color: #10b981; }
.highlight-box.success:hover {
  border-color: rgba(16, 185, 129, 0.4);
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.1);
}

.highlight-box.warning {
  border: 1px solid rgba(245, 158, 11, 0.2);
  background: rgba(245, 158, 11, 0.06);
}
.highlight-box.warning .highlight-box-icon { color: #f59e0b; }
.highlight-box.warning:hover {
  border-color: rgba(245, 158, 11, 0.4);
  box-shadow: 0 4px 16px rgba(245, 158, 11, 0.1);
}

.highlight-box-icon { width: 24px; height: 24px; flex-shrink: 0; }

/* ===== 权利卡片 ===== */
.rights-grid {
  display: grid; grid-template-columns: repeat(3, 1fr);
  gap: 14px; margin: 16px 0 20px;
}

.right-card {
  display: flex; align-items: center; gap: 14px;
  padding: 16px; border-radius: 14px;
  border: 1px solid rgba(139, 92, 246, 0.15);
  cursor: default; transition: all 0.3s ease;
}

.right-card:hover {
  border-color: rgba(139, 92, 246, 0.4);
  box-shadow: 0 8px 20px rgba(139, 92, 246, 0.12);
  transform: translateY(-4px);
}

.right-card:hover .right-icon { transform: rotate(10deg) scale(1.15); }

.right-icon {
  width: 40px; height: 40px;
  display: flex; align-items: center; justify-content: center;
  background: linear-gradient(135deg, #8b5cf6, #a855f7);
  border-radius: 12px; flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
  transition: transform 0.3s ease;
}

.right-icon svg { width: 20px; height: 20px; color: white; }

.right-name { font-size: 14px; font-weight: 700; color: #1e1b4b; margin-bottom: 2px; }
.right-desc { font-size: 12px; color: #64748b; line-height: 1.4; }

/* ===== 联系卡片 ===== */
.contact-cards {
  display: grid; grid-template-columns: repeat(3, 1fr);
  gap: 16px; margin-top: 16px;
}

.contact-card {
  display: flex; align-items: center; gap: 14px;
  padding: 20px; border-radius: 14px;
  border: 1px solid rgba(219, 39, 119, 0.15);
  cursor: pointer; transition: all 0.3s ease;
}

.contact-card:hover {
  border-color: #DB2777;
  box-shadow: 0 8px 24px rgba(219, 39, 119, 0.12);
  transform: translateY(-4px);
}

.contact-card:hover .contact-icon { transform: scale(1.15) rotate(-5deg); }

.contact-icon {
  width: 44px; height: 44px;
  display: flex; align-items: center; justify-content: center;
  background: linear-gradient(135deg, #DB2777, #F472B6);
  border-radius: 12px; flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
  transition: transform 0.3s ease;
}

.contact-icon svg { width: 22px; height: 22px; color: white; }

.contact-label {
  font-size: 12px; font-weight: 600; color: #94a3b8;
  text-transform: uppercase; letter-spacing: 0.05em; margin-bottom: 2px;
}

.contact-value { font-size: 14px; font-weight: 600; color: #1e1b4b; }

/* ===== 底部 ===== */
.privacy-footer {
  margin-top: 40px; padding-top: 28px;
  border-top: 1px solid rgba(219, 39, 119, 0.15);
  display: flex; align-items: center; justify-content: space-between;
  gap: 24px;
  animation: fadeInUp 0.6s ease-out 0.55s backwards;
}

.effective-date {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 10px 20px; border-radius: 12px;
  font-size: 13px; color: #64748b; font-weight: 500;
}

.effective-date svg { color: #DB2777; }

.back-link {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #DB2777, #F472B6);
  color: white; text-decoration: none;
  border-radius: 14px; font-weight: 600; font-size: 14px;
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
  .toc-nav { display: none; }
  .rights-grid { grid-template-columns: repeat(2, 1fr); }
}

@media (max-width: 768px) {
  .page-header { padding: 60px 16px; }
  .header-title { font-size: 2.25rem; }
  .icon-wrapper { width: 80px; height: 80px; border-radius: 22px; }
  .privacy-body { padding: 32px 16px 48px; }
  .glass-section { padding: 20px; }
  .contact-cards { grid-template-columns: 1fr; }
  .privacy-footer { flex-direction: column; align-items: center; text-align: center; }
}

@media (max-width: 576px) {
  .header-title { font-size: 1.875rem; }
  .section-title { font-size: 17px; }
  .section-number { width: 30px; height: 30px; font-size: 13px; border-radius: 10px; }
  .rights-grid { grid-template-columns: 1fr; }
}
</style>

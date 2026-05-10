<template>
  <div class="design-system-demo">
    <div class="design-system-container">
      <h1 class="design-system-title">AqiCloud-AiPan 设计系统</h1>
      <p class="design-system-subtitle">设计系统组件示例和实现指南</p>

      <!-- 色彩系统 -->
      <section class="design-section">
        <h2 class="section-title">色彩系统</h2>
        
        <h3 class="subsection-title">主色调</h3>
        <div class="color-grid">
          <div class="color-swatch" v-for="color in primaryColors" :key="color.name">
            <div class="color-preview" :style="{ backgroundColor: color.hex }"></div>
            <div class="color-info">
              <div class="color-name">{{ color.name }}</div>
              <div class="color-hex">{{ color.hex }}</div>
            </div>
          </div>
        </div>

        <h3 class="subsection-title">功能色</h3>
        <div class="color-grid">
          <div class="color-swatch" v-for="color in functionalColors" :key="color.name">
            <div class="color-preview" :style="{ backgroundColor: color.hex }"></div>
            <div class="color-info">
              <div class="color-name">{{ color.name }}</div>
              <div class="color-hex">{{ color.hex }}</div>
            </div>
          </div>
        </div>
      </section>

      <!-- 排版系统 -->
      <section class="design-section">
        <h2 class="section-title">排版系统</h2>
        
        <div class="typography-demo">
          <h1 class="typo-h1">H1 - 主标题 (48px)</h1>
          <h2 class="typo-h2">H2 - 章节标题 (36px)</h2>
          <h3 class="typo-h3">H3 - 小节标题 (24px)</h3>
          <h4 class="typo-h4">H4 - 组件标题 (20px)</h4>
          <p class="typo-lg">Body Large - 重要正文 (18px)</p>
          <p class="typo-base">Body - 默认正文 (16px)</p>
          <p class="typo-sm">Body Small - 次要正文 (14px)</p>
          <p class="typo-xs">Caption - 辅助文本 (12px)</p>
        </div>
      </section>

      <!-- 按钮组件 -->
      <section class="design-section">
        <h2 class="section-title">按钮组件</h2>
        
        <div class="button-group">
          <div class="button-row">
            <button class="ds-btn ds-btn-primary">主要按钮</button>
            <button class="ds-btn ds-btn-primary" disabled>禁用状态</button>
          </div>
          
          <div class="button-row">
            <button class="ds-btn ds-btn-secondary">次要按钮</button>
            <button class="ds-btn ds-btn-secondary" disabled>禁用状态</button>
          </div>
          
          <div class="button-row">
            <button class="ds-btn ds-btn-cta">CTA按钮</button>
            <button class="ds-btn ds-btn-cta" disabled>禁用状态</button>
          </div>
        </div>
      </section>

      <!-- 输入框组件 -->
      <section class="design-section">
        <h2 class="section-title">输入框组件</h2>
        
        <div class="input-demo">
          <div class="ds-input-group">
            <label class="ds-input-label">用户名</label>
            <input type="text" class="ds-input" placeholder="请输入用户名" />
          </div>
          
          <div class="ds-input-group">
            <label class="ds-input-label">密码</label>
            <input type="password" class="ds-input" placeholder="请输入密码" />
          </div>
          
          <div class="ds-input-group">
            <label class="ds-input-label">禁用状态</label>
            <input type="text" class="ds-input" disabled placeholder="禁用的输入框" />
          </div>
          
          <div class="ds-input-group">
            <label class="ds-input-label">错误状态</label>
            <input type="text" class="ds-input ds-input-error" placeholder="有错误的输入框" />
            <span class="ds-input-error-text">请输入有效的内容</span>
          </div>
        </div>
      </section>

      <!-- 卡片组件 -->
      <section class="design-section">
        <h2 class="section-title">卡片组件</h2>
        
        <div class="card-grid">
          <div class="ds-card">
            <div class="ds-card-header">
              <h3 class="ds-card-title">卡片标题</h3>
            </div>
            <div class="ds-card-body">
              <p class="ds-card-text">这是卡片的内容区域。可以放置文本、图片或其他组件。</p>
            </div>
            <div class="ds-card-footer">
              <button class="ds-btn ds-btn-secondary">取消</button>
              <button class="ds-btn ds-btn-primary">确认</button>
            </div>
          </div>
          
          <div class="ds-card">
            <div class="ds-card-body">
              <h3 class="ds-card-title">简化卡片</h3>
              <p class="ds-card-text">没有头部和底部的简单卡片。</p>
              <button class="ds-btn ds-btn-cta">了解更多</button>
            </div>
          </div>
        </div>
      </section>

      <!-- 标签 -->
      <section class="design-section">
        <h2 class="section-title">标签</h2>
        
        <div class="tag-group">
          <span class="ds-tag ds-tag-primary">主要标签</span>
          <span class="ds-tag ds-tag-success">成功</span>
          <span class="ds-tag ds-tag-warning">警告</span>
          <span class="ds-tag ds-tag-error">错误</span>
          <span class="ds-tag ds-tag-info">信息</span>
        </div>
      </section>

      <!-- 动画效果 -->
      <section class="design-section">
        <h2 class="section-title">动画效果</h2>
        
        <div class="animation-group">
          <button class="ds-btn ds-btn-primary hover-lift">悬停提升</button>
          <button class="ds-btn ds-btn-secondary hover-shadow">悬停阴影</button>
          <button class="ds-btn ds-btn-cta transition-fast">快速过渡</button>
          <button class="ds-btn ds-btn-primary transition-slow">慢速过渡</button>
        </div>
      </section>

      <!-- 响应式示例 -->
      <section class="design-section">
        <h2 class="section-title">响应式网格</h2>
        
        <div class="responsive-grid">
          <div class="grid-card">
            <div class="grid-number">1</div>
            <div class="grid-desc">移动端: 1列</div>
          </div>
          <div class="grid-card">
            <div class="grid-number">2</div>
            <div class="grid-desc">平板: 2列</div>
          </div>
          <div class="grid-card">
            <div class="grid-number">3</div>
            <div class="grid-desc">桌面: 4列</div>
          </div>
          <div class="grid-card">
            <div class="grid-number">4</div>
            <div class="grid-desc">自适应</div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
const primaryColors = [
  { name: 'Primary', hex: '#6366F1' },
  { name: 'Secondary', hex: '#818CF8' },
  { name: 'CTA', hex: '#10B981' },
  { name: 'Background', hex: '#F5F3FF' },
  { name: 'Text Primary', hex: '#1E1B4B' },
];

const functionalColors = [
  { name: 'Success', hex: '#10B981' },
  { name: 'Warning', hex: '#F59E0B' },
  { name: 'Error', hex: '#EF4444' },
  { name: 'Info', hex: '#3B82F6' },
];
</script>

<style scoped>
/* Container */
.design-system-demo {
  min-height: 100vh;
  background-color: #F5F3FF;
  padding: 24px;
  box-sizing: border-box;
  overflow-y: auto;
}

.design-system-container {
  max-width: 1200px;
  margin: 0 auto;
  padding-bottom: 48px;
}

.design-system-title {
  font-size: 48px;
  font-weight: 700;
  color: #1E1B4B;
  margin: 0 0 16px 0;
}

.design-system-subtitle {
  font-size: 18px;
  color: #64748B;
  margin: 0 0 48px 0;
}

/* Sections */
.design-section {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 32px;
  margin-bottom: 32px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.section-title {
  font-size: 36px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0 0 24px 0;
}

.subsection-title {
  font-size: 20px;
  font-weight: 500;
  color: #1E1B4B;
  margin: 24px 0 16px 0;
}

/* Colors */
.color-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.color-swatch {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.color-preview {
  width: 100px;
  height: 100px;
  border-radius: 8px;
  border: 1px solid #E2E8F0;
}

.color-info {
  font-size: 12px;
}

.color-name {
  font-weight: 500;
  color: #1E1B4B;
}

.color-hex {
  color: #64748B;
}

/* Typography */
.typography-demo {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.typo-h1 {
  font-size: 48px;
  font-weight: 700;
  color: #1E1B4B;
  margin: 0;
}

.typo-h2 {
  font-size: 36px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0;
}

.typo-h3 {
  font-size: 24px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0;
}

.typo-h4 {
  font-size: 20px;
  font-weight: 500;
  color: #1E1B4B;
  margin: 0;
}

.typo-lg {
  font-size: 18px;
  color: #1E1B4B;
  margin: 0;
}

.typo-base {
  font-size: 16px;
  color: #1E1B4B;
  margin: 0;
}

.typo-sm {
  font-size: 14px;
  color: #1E1B4B;
  margin: 0;
}

.typo-xs {
  font-size: 12px;
  color: #1E1B4B;
  margin: 0;
}

/* Buttons */
.button-group {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.button-row {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.ds-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.ds-btn-primary {
  background-color: #6366F1;
  color: white;
}

.ds-btn-primary:hover:not(:disabled) {
  background-color: #4F46E5;
  transform: translateY(-1px);
}

.ds-btn-primary:disabled {
  background-color: #E2E8F0;
  cursor: not-allowed;
}

.ds-btn-secondary {
  background-color: #FFFFFF;
  color: #6366F1;
  border: 1px solid #E2E8F0;
}

.ds-btn-secondary:hover:not(:disabled) {
  border-color: #6366F1;
  background-color: #F5F3FF;
}

.ds-btn-secondary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.ds-btn-cta {
  background-color: #10B981;
  color: white;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.ds-btn-cta:hover:not(:disabled) {
  background-color: #059669;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
  transform: translateY(-2px);
}

.ds-btn-cta:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Input */
.input-demo {
  display: flex;
  flex-direction: column;
  gap: 16px;
  max-width: 400px;
}

.ds-input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.ds-input-label {
  font-size: 14px;
  font-weight: 500;
  color: #1E1B4B;
}

.ds-input {
  padding: 12px 16px;
  border: 1px solid #E2E8F0;
  border-radius: 8px;
  font-size: 16px;
  color: #1E1B4B;
  transition: all 0.2s ease;
  box-sizing: border-box;
}

.ds-input:focus {
  outline: none;
  border-color: #6366F1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

.ds-input:disabled {
  background-color: #F5F3FF;
  cursor: not-allowed;
}

.ds-input-error {
  border-color: #EF4444;
}

.ds-input-error-text {
  font-size: 12px;
  color: #EF4444;
}

/* Card */
.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 24px;
}

.ds-card {
  background-color: #FFFFFF;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  overflow: hidden;
  transition: box-shadow 0.2s ease;
  border: 1px solid #E2E8F0;
}

.ds-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.ds-card-header {
  padding: 20px 24px;
  border-bottom: 1px solid #E2E8F0;
}

.ds-card-title {
  font-size: 20px;
  font-weight: 500;
  color: #1E1B4B;
  margin: 0;
}

.ds-card-body {
  padding: 24px;
}

.ds-card-text {
  font-size: 16px;
  color: #64748B;
  margin: 0;
}

.ds-card-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid #E2E8F0;
  background-color: #F5F3FF;
}

/* Tags */
.tag-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.ds-tag {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 9999px;
  font-size: 12px;
  font-weight: 500;
}

.ds-tag-primary {
  background-color: rgba(99, 102, 241, 0.1);
  color: #6366F1;
}

.ds-tag-success {
  background-color: rgba(16, 185, 129, 0.1);
  color: #10B981;
}

.ds-tag-warning {
  background-color: rgba(245, 158, 11, 0.1);
  color: #F59E0B;
}

.ds-tag-error {
  background-color: rgba(239, 68, 68, 0.1);
  color: #EF4444;
}

.ds-tag-info {
  background-color: rgba(59, 130, 246, 0.1);
  color: #3B82F6;
}

/* Animation */
.animation-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.hover-lift:hover {
  transform: translateY(-2px);
}

.hover-shadow:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.transition-fast {
  transition-duration: 0.15s;
}

.transition-slow {
  transition-duration: 0.3s;
}

/* Responsive Grid */
.responsive-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
}

@media (min-width: 768px) {
  .responsive-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .responsive-grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

.grid-card {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 24px;
  text-align: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  border: 1px solid #E2E8F0;
}

.grid-number {
  font-size: 48px;
  font-weight: 700;
  color: #6366F1;
  margin-bottom: 8px;
}

.grid-desc {
  font-size: 14px;
  color: #64748B;
}
</style>

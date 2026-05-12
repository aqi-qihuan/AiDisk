<template>
  <div class="component-usage-demo">
    <div class="container">
      <h1 class="page-title">设计系统组件使用示例</h1>
      <p class="page-subtitle">在实际项目中如何使用设计系统组件</p>

      <!-- 按钮组件示例 -->
      <section class="demo-section">
        <h2 class="section-title">按钮组件 (DSButton)</h2>

        <div class="demo-block">
          <h3>基础用法</h3>
          <div class="demo-row">
            <DSButton variant="primary">主要按钮</DSButton>
            <DSButton variant="secondary">次要按钮</DSButton>
            <DSButton variant="cta">CTA按钮</DSButton>
            <DSButton variant="danger">危险按钮</DSButton>
          </div>
        </div>

        <div class="demo-block">
          <h3>不同尺寸</h3>
          <div class="demo-row">
            <DSButton variant="primary" size="small">小按钮</DSButton>
            <DSButton variant="primary" size="medium">中按钮</DSButton>
            <DSButton variant="primary" size="large">大按钮</DSButton>
          </div>
        </div>

        <div class="demo-block">
          <h3>不同状态</h3>
          <div class="demo-row">
            <DSButton variant="primary">正常状态</DSButton>
            <DSButton variant="primary" disabled>禁用状态</DSButton>
            <DSButton variant="primary" :loading="true">加载中</DSButton>
          </div>
        </div>

        <div class="demo-block">
          <h3>带图标</h3>
          <div class="demo-row">
            <DSButton variant="primary">
              <template #icon>
                <span>📁</span>
              </template>
              上传文件
            </DSButton>
            <DSButton variant="secondary">
              <template #icon>
                <span>🗑️</span>
              </template>
              删除
            </DSButton>
          </div>
        </div>

        <div class="code-block">
          <pre><code>&lt;DSButton variant="primary"&gt;主要按钮&lt;/DSButton&gt;
&lt;DSButton variant="primary" size="large" :loading="true"&gt;加载中&lt;/DSButton&gt;
&lt;DSButton variant="primary" @click="handleClick"&gt;点击事件&lt;/DSButton&gt;</code></pre>
        </div>
      </section>

      <!-- 输入框组件示例 -->
      <section class="demo-section">
        <h2 class="section-title">输入框组件 (DSInput)</h2>

        <div class="demo-block">
          <h3>基础用法</h3>
          <div class="demo-row">
            <DSInput
              v-model="inputValue"
              placeholder="请输入内容"
              label="用户名"
            />
          </div>
        </div>

        <div class="demo-block">
          <h3>不同尺寸</h3>
          <div class="demo-row">
            <DSInput v-model="inputValue" size="small" placeholder="小输入框" />
            <DSInput
              v-model="inputValue"
              size="medium"
              placeholder="中输入框"
            />
            <DSInput v-model="inputValue" size="large" placeholder="大输入框" />
          </div>
        </div>

        <div class="demo-block">
          <h3>不同状态</h3>
          <div class="demo-row">
            <DSInput v-model="inputValue" placeholder="正常状态" label="正常" />
            <DSInput
              v-model="inputValue"
              placeholder="禁用状态"
              label="禁用"
              disabled
            />
            <DSInput
              v-model="inputValue"
              placeholder="错误状态"
              label="错误"
              error="请输入有效的内容"
            />
          </div>
        </div>

        <div class="demo-block">
          <h3>带清除功能</h3>
          <div class="demo-row">
            <DSInput
              v-model="inputValue"
              placeholder="输入内容后显示清除按钮"
              label="可清除"
              showClear
              @clear="handleClear"
            />
          </div>
        </div>

        <div class="code-block">
          <pre><code>&lt;DSInput
  v-model="inputValue"
  label="用户名"
  placeholder="请输入用户名"
  required
  :error="errorMessage"
  @change="handleChange"
/&gt;</code></pre>
        </div>
      </section>

      <!-- 卡片组件示例 -->
      <section class="demo-section">
        <h2 class="section-title">卡片组件 (DSCard)</h2>

        <div class="demo-block">
          <h3>基础卡片</h3>
          <div class="demo-row">
            <DSCard title="卡片标题" style="max-width: 400px">
              <p>这是卡片的内容区域，可以放置文本、图片或其他组件。</p>
              <p>卡片支持悬停效果和阴影效果。</p>
            </DSCard>
          </div>
        </div>

        <div class="demo-block">
          <h3>带头部和底部</h3>
          <div class="demo-row">
            <DSCard title="文件信息" hoverable style="max-width: 400px">
              <template #extra>
                <DSTag variant="success">活跃</DSTag>
              </template>
              <p><strong>文件名:</strong> document.pdf</p>
              <p><strong>大小:</strong> 2.5 MB</p>
              <p><strong>上传时间:</strong> 2026-03-11</p>
              <template #footer>
                <DSButton variant="secondary" size="small">取消</DSButton>
                <DSButton variant="primary" size="small">下载</DSButton>
              </template>
            </DSCard>
          </div>
        </div>

        <div class="code-block">
          <pre><code>&lt;DSCard title="卡片标题" hoverable bordered&gt;
  &lt;template #extra&gt;
    &lt;DSTag variant="success"&gt;标签&lt;/DSTag&gt;
  &lt;/template&gt;
  卡片内容
  &lt;template #footer&gt;
    &lt;DSButton&gt;按钮&lt;/DSButton&gt;
  &lt;/template&gt;
&lt;/DSCard&gt;</code></pre>
        </div>
      </section>

      <!-- 标签组件示例 -->
      <section class="demo-section">
        <h2 class="section-title">标签组件 (DSTag)</h2>

        <div class="demo-block">
          <h3>不同颜色</h3>
          <div class="demo-row">
            <DSTag variant="primary">主要</DSTag>
            <DSTag variant="secondary">次要</DSTag>
            <DSTag variant="success">成功</DSTag>
            <DSTag variant="warning">警告</DSTag>
            <DSTag variant="danger">危险</DSTag>
            <DSTag variant="info">信息</DSTag>
          </div>
        </div>

        <div class="demo-block">
          <h3>不同尺寸</h3>
          <div class="demo-row">
            <DSTag variant="primary" size="small">小标签</DSTag>
            <DSTag variant="primary" size="medium">中标签</DSTag>
            <DSTag variant="primary" size="large">大标签</DSTag>
          </div>
        </div>

        <div class="demo-block">
          <h3>可关闭标签</h3>
          <div class="demo-row">
            <DSTag variant="primary" closeable @close="handleClose"
              >可关闭</DSTag
            >
            <DSTag variant="success" closeable @close="handleClose">成功</DSTag>
            <DSTag variant="danger" closeable @close="handleClose">删除</DSTag>
          </div>
        </div>

        <div class="code-block">
          <pre><code>&lt;DSTag variant="success"&gt;成功&lt;/DSTag&gt;
&lt;DSTag variant="danger" closeable @close="handleClose"&gt;可删除&lt;/DSTag&gt;</code></pre>
        </div>
      </section>

      <!-- 综合示例 -->
      <section class="demo-section">
        <h2 class="section-title">综合示例：文件卡片</h2>

        <div class="demo-block">
          <div class="demo-row">
            <DSCard
              title="项目文档.pdf"
              hoverable
              shadow
              style="max-width: 400px"
            >
              <template #extra>
                <DSTag variant="success">已同步</DSTag>
              </template>
              <div style="margin-bottom: 16px">
                <p><strong>文件大小:</strong> 3.2 MB</p>
                <p><strong>修改时间:</strong> 2026-03-11 14:30</p>
                <p><strong>文件类型:</strong> PDF文档</p>
              </div>
              <div style="margin-bottom: 16px">
                <DSInput
                  v-model="fileName"
                  label="重命名"
                  placeholder="输入新文件名"
                  size="small"
                />
              </div>
              <template #footer>
                <DSButton variant="secondary" size="small">分享</DSButton>
                <DSButton variant="primary" size="small">下载</DSButton>
              </template>
            </DSCard>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { DSButton, DSInput, DSCard, DSTag } from "./index";

const inputValue = ref("");
const fileName = ref("项目文档.pdf");

const handleClick = () => {
  console.log("按钮被点击");
};

const handleClear = () => {
  console.log("输入框已清除");
  inputValue.value = "";
};

const handleChange = (value: string) => {
  console.log("输入值变化:", value);
};

const handleClose = () => {
  console.log("标签已关闭");
};
</script>

<style scoped>
.component-usage-demo {
  min-height: 100vh;
  background-color: #f5f3ff;
  padding: 24px;
  box-sizing: border-box;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

.page-title {
  font-size: 48px;
  font-weight: 700;
  color: #1e1b4b;
  margin: 0 0 16px 0;
}

.page-subtitle {
  font-size: 18px;
  color: #64748b;
  margin: 0 0 48px 0;
}

.demo-section {
  background: #ffffff;
  border-radius: 12px;
  padding: 32px;
  margin-bottom: 32px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

.section-title {
  font-size: 36px;
  font-weight: 600;
  color: #1e1b4b;
  margin: 0 0 24px 0;
}

.demo-block {
  margin-bottom: 32px;
}

.demo-block h3 {
  font-size: 18px;
  font-weight: 500;
  color: #64748b;
  margin: 0 0 16px 0;
}

.demo-row {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  align-items: center;
}

.code-block {
  background-color: #1e1b4b;
  border-radius: 8px;
  padding: 16px;
  overflow-x: auto;
}

.code-block pre {
  margin: 0;
}

.code-block code {
  font-family: "Fira Code", "Courier New", monospace;
  font-size: 14px;
  color: #f5f3ff;
  line-height: 1.6;
}
</style>

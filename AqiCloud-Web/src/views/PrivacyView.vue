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
          <svg
            class="header-svg-icon"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
          </svg>
        </div>
        <h1 class="header-title">{{ t("page.privacy.title") }}</h1>
        <p class="header-subtitle">{{ t("page.privacy.subtitle") }}</p>
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
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
              </svg>
              {{ t("page.privacy.tocTitle") }}
            </div>
            <ul class="toc-list">
              <li
                v-for="(item, index) in tocItems"
                :key="index"
                class="toc-item"
              >
                <a
                  :href="'#section-' + (index + 1)"
                  class="toc-link"
                  @click.prevent="scrollToSection(index)"
                >
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
                <h2 class="section-title">
                  {{ t("page.privacy.section1Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section1P1") }}</p>
              <p class="section-text">{{ t("page.privacy.section1P2") }}</p>
            </section>

            <section id="section-2" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">2</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section2Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section2Intro") }}</p>
              <ul class="privacy-list">
                <li>
                  <div class="list-icon-box blue">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"
                      ></path>
                      <circle cx="12" cy="7" r="4"></circle>
                    </svg>
                  </div>
                  <div>
                    <strong>{{ t("page.privacy.section2List1Title") }}</strong
                    >{{ t("page.privacy.section2List1Desc") }}
                  </div>
                </li>
                <li>
                  <div class="list-icon-box purple">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"
                      ></path>
                      <circle cx="12" cy="7" r="4"></circle>
                    </svg>
                  </div>
                  <div>
                    <strong>{{ t("page.privacy.section2List2Title") }}</strong
                    >{{ t("page.privacy.section2List2Desc") }}
                  </div>
                </li>
                <li>
                  <div class="list-icon-box green">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"
                      ></path>
                    </svg>
                  </div>
                  <div>
                    <strong>{{ t("page.privacy.section2List3Title") }}</strong
                    >{{ t("page.privacy.section2List3Desc") }}
                  </div>
                </li>
                <li>
                  <div class="list-icon-box orange">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <polyline
                        points="22 12 18 12 15 21 9 3 6 12 2 12"
                      ></polyline>
                    </svg>
                  </div>
                  <div>
                    <strong>{{ t("page.privacy.section2List4Title") }}</strong
                    >{{ t("page.privacy.section2List4Desc") }}
                  </div>
                </li>
                <li>
                  <div class="list-icon-box pink">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
                      ></path>
                    </svg>
                  </div>
                  <div>
                    <strong>{{ t("page.privacy.section2List5Title") }}</strong
                    >{{ t("page.privacy.section2List5Desc") }}
                  </div>
                </li>
              </ul>
              <p class="section-text">{{ t("page.privacy.section2P1") }}</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section2List5") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section2List6") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section2List7") }}
                </li>
              </ul>
            </section>

            <section id="section-3" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">3</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section3Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section3Intro") }}</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section3List1") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section3List2") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section3List3") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section3List4") }}
                </li>
                <li>
                  <svg
                    class="check-icon"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section3List5") }}
                </li>
              </ul>
              <p class="section-text">{{ t("page.privacy.section3P1") }}</p>
            </section>

            <section id="section-4" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number shield">4</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section4Title") }}
                </h2>
              </div>
              <div class="highlight-box security">
                <div class="highlight-box-icon">
                  <svg
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <path
                      d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
                    ></path>
                  </svg>
                </div>
                <p class="section-text" style="margin: 0">
                  {{ t("page.privacy.section4Intro") }}
                </p>
              </div>
              <p class="section-text">{{ t("page.privacy.section4P1") }}</p>
            </section>

            <section id="section-5" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number info">5</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section5Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section5Intro") }}</p>
              <ul class="privacy-list simple-list">
                <li>
                  <svg
                    class="check-icon shield-check"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section5List1") }}
                </li>
                <li>
                  <svg
                    class="check-icon shield-check"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section5List2") }}
                </li>
                <li>
                  <svg
                    class="check-icon shield-check"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section5List3") }}
                </li>
                <li>
                  <svg
                    class="check-icon shield-check"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12"></polyline></svg
                  >{{ t("page.privacy.section5List4") }}
                </li>
              </ul>
              <p class="section-text">{{ t("page.privacy.section5P1") }}</p>
            </section>

            <section id="section-6" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">6</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section6Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section6P1") }}</p>
              <p class="section-text">{{ t("page.privacy.section6P2") }}</p>
              <p class="section-text">{{ t("page.privacy.section6P3") }}</p>
            </section>

            <section id="section-7" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number rights">7</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section7Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section7Intro") }}</p>
              <div class="rights-grid">
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"
                      ></path>
                      <circle cx="12" cy="12" r="3"></circle>
                    </svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">
                      {{ t("page.privacy.section7List1") }}
                    </div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"
                      ></path>
                      <path
                        d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"
                      ></path>
                    </svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">
                      {{ t("page.privacy.section7List2") }}
                    </div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <polyline points="3 6 5 6 21 6"></polyline>
                      <path
                        d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"
                      ></path>
                    </svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">
                      {{ t("page.privacy.section7List3") }}
                    </div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <rect
                        x="3"
                        y="11"
                        width="18"
                        height="11"
                        rx="2"
                        ry="2"
                      ></rect>
                      <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                    </svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">
                      {{ t("page.privacy.section7List4") }}
                    </div>
                  </div>
                </div>
                <div class="right-card glass-card">
                  <div class="right-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"
                      ></path>
                      <polyline points="7 10 12 15 17 10"></polyline>
                      <line x1="12" y1="15" x2="12" y2="3"></line>
                    </svg>
                  </div>
                  <div class="right-info">
                    <div class="right-name">
                      {{ t("page.privacy.section7List5") }}
                    </div>
                  </div>
                </div>
              </div>
              <p class="section-text">{{ t("page.privacy.section7P1") }}</p>
            </section>

            <section id="section-8" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number warning">8</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section8Title") }}
                </h2>
              </div>
              <div class="highlight-box warning">
                <div class="highlight-box-icon">
                  <svg
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <path
                      d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"
                    ></path>
                    <line x1="12" y1="9" x2="12" y2="13"></line>
                    <line x1="12" y1="17" x2="12.01" y2="17"></line>
                  </svg>
                </div>
                <p class="section-text" style="margin: 0">
                  {{ t("page.privacy.section8P1") }}
                </p>
              </div>
              <p class="section-text">{{ t("page.privacy.section8P2") }}</p>
            </section>

            <section id="section-9" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">9</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section9Title") }}
                </h2>
              </div>
              <p class="section-text">{{ t("page.privacy.section9P1") }}</p>
              <p class="section-text">{{ t("page.privacy.section9P2") }}</p>
            </section>

            <section id="section-10" class="privacy-section glass-section">
              <div class="section-header">
                <div class="section-number">10</div>
                <h2 class="section-title">
                  {{ t("page.privacy.section10Title") }}
                </h2>
              </div>
              <p class="section-text">
                {{ t("page.privacy.section10IntroP1") }}
              </p>
              <div class="contact-cards">
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"
                      ></path>
                      <polyline points="22,6 12,13 2,6"></polyline>
                    </svg>
                  </div>
                  <div>
                    <div class="contact-label">
                      {{ t("page.privacy.contactEmailLabel") }}
                    </div>
                    <div class="contact-value">
                      {{ t("page.privacy.privacyEmail") }}
                    </div>
                  </div>
                </div>
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"
                      ></path>
                      <circle cx="12" cy="10" r="3"></circle>
                    </svg>
                  </div>
                  <div>
                    <div class="contact-label">
                      {{ t("page.terms.contactAddress") }}
                    </div>
                    <div class="contact-value">
                      {{ t("page.terms.addressValue") }}
                    </div>
                  </div>
                </div>
                <div class="contact-card glass-card">
                  <div class="contact-icon">
                    <svg
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"
                      ></path>
                    </svg>
                  </div>
                  <div>
                    <div class="contact-label">
                      {{ t("page.privacy.contactPhoneLabel") }}
                    </div>
                    <div class="contact-value">
                      {{ t("page.privacy.phoneValue") }}
                    </div>
                  </div>
                </div>
              </div>
            </section>

            <!-- 底部 -->
            <div class="privacy-footer">
              <div class="effective-date glass-card">
                <svg
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <circle cx="12" cy="12" r="10"></circle>
                  <polyline points="12 6 12 12 16 14"></polyline>
                </svg>
                <span>{{ t("page.privacy.lastUpdated") }}</span>
              </div>
              <router-link to="/" class="back-link">
                <svg
                  width="18"
                  height="18"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <line x1="19" y1="12" x2="5" y2="12"></line>
                  <polyline points="12 19 5 12 12 5"></polyline>
                </svg>
                {{ t("page.privacy.backHome") }}
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const tocItems = computed(
  () => t("page.privacy.tocItems", { returnObjects: true }) as string[],
);

const scrollToSection = (index: number) => {
  const el = document.getElementById(`section-${index + 1}`);
  if (el) {
    el.scrollIntoView({ behavior: "smooth", block: "start" });
  }
};
</script>

<style>
/* HOK 暗色设计系统 — Privacy */
:root {
  --primary: #db2777;
  --secondary: #f472b6;
  --cta: #d97706;
  --gold: #fbbf24;
  --bg-dark: #0b0b10;
  --bg-surface: #14141c;
  --bg-card: #1a1a24;
  --text-primary: #f8fafc;
  --text-secondary: #94a3b8;
  --text-tertiary: #64748b;
  --font-primary: "Fira Sans", "Plus Jakarta Sans", -apple-system, sans-serif;
  --glass-bg: rgba(255, 255, 255, 0.04);
  --glass-border: 1px solid rgba(255, 255, 255, 0.06);
  --glass-shadow: 0 8px 32px rgba(0, 0, 0, 0.4), 0 2px 8px rgba(0, 0, 0, 0.2);
  --glass-backdrop: blur(24px) saturate(180%);
}
</style>
<style scoped>
/* ===== HOK 暗色主题 - Privacy ===== */
.privacy-page {
  min-height: 100vh;
  background: #14141c;
  font-family: var(--font-primary);
  overflow-x: hidden;
  position: relative;
}

/* ===== HOK Glow Orbs ===== */
.privacy-page::before,
.privacy-page::after {
  content: "";
  position: fixed;
  border-radius: 50%;
  filter: blur(120px);
  pointer-events: none;
  z-index: 0;
}
.privacy-page::before {
  width: 500px;
  height: 500px;
  background: radial-gradient(circle, #db2777, transparent);
  top: -200px;
  right: -100px;
  opacity: 0.25;
}
.privacy-page::after {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, #d97706, transparent);
  bottom: -150px;
  left: -100px;
  opacity: 0.2;
}

/* ===== 页面头部 ===== */
.page-header {
  position: relative;
  padding: 80px 24px;
  background: linear-gradient(135deg, #14141c 0%, #1a1a24 50%, #14141c 100%);
  overflow: hidden;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
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
  opacity: 0.3;
  animation: float 25s infinite ease-in-out;
}
.orb-1 {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, rgba(219, 39, 119, 0.3), transparent);
  top: -150px;
  right: -100px;
}
.orb-2 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, rgba(217, 119, 6, 0.25), transparent);
  bottom: -100px;
  left: -80px;
  animation-delay: 8s;
}
.orb-3 {
  width: 250px;
  height: 250px;
  background: radial-gradient(circle, rgba(219, 39, 119, 0.15), transparent);
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
  background: rgba(255, 255, 255, 0.04);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 28px;
  margin-bottom: 24px;
  color: #fbbf24;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  animation: float 6s infinite ease-in-out;
}
.header-svg-icon {
  width: 48px;
  height: 48px;
  filter: drop-shadow(0 2px 8px rgba(0, 0, 0, 0.3));
}
.header-title {
  font-size: 3rem;
  font-weight: 800;
  background: linear-gradient(135deg, #fbbf24 0%, #fde68a 40%, #f472b6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0 0 8px;
  letter-spacing: -0.02em;
}
.header-subtitle {
  font-size: 1.125rem;
  color: #94a3b8;
  font-weight: 500;
  letter-spacing: 0.05em;
  text-transform: uppercase;
}
.header-divider {
  width: 60px;
  height: 4px;
  background: linear-gradient(90deg, #db2777, #fbbf24);
  border-radius: 2px;
  margin: 20px auto 0;
}
.privacy-body {
  padding: 48px 24px 60px;
  position: relative;
  z-index: 1;
}
.privacy-container {
  max-width: 1100px;
  margin: 0 auto;
}
.privacy-layout {
  display: flex;
  gap: 32px;
  align-items: flex-start;
}

/* TOC (HOK 暗色) */
.toc-nav {
  position: sticky;
  top: 88px;
  width: 220px;
  flex-shrink: 0;
  background: rgba(255, 255, 255, 0.04);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 20px;
  padding: 24px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  animation: fadeInLeft 0.6s ease-out;
}
.toc-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 700;
  color: #f8fafc;
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
  color: #94a3b8;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.25s ease;
}
.toc-link:hover {
  background: rgba(219, 39, 119, 0.08);
  color: #db2777;
  transform: translateX(4px);
}
.toc-number {
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(219, 39, 119, 0.15);
  color: #db2777;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 700;
  flex-shrink: 0;
  transition: all 0.25s ease;
}
.toc-link:hover .toc-number {
  background: #db2777;
  color: white;
  transform: rotate(10deg) scale(1.1);
}

/* Glass Sections */
.glass-section {
  background: rgba(26, 26, 36, 0.85);
  backdrop-filter: var(--glass-backdrop);
  -webkit-backdrop-filter: var(--glass-backdrop);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 20px;
  padding: 28px 32px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  margin-bottom: 20px;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  animation: fadeInUp 0.5s ease-out backwards;
}
.glass-section:nth-child(1) {
  animation-delay: 0.05s;
}
.glass-section:nth-child(2) {
  animation-delay: 0.1s;
}
.glass-section:nth-child(3) {
  animation-delay: 0.15s;
}
.glass-section:nth-child(4) {
  animation-delay: 0.2s;
}
.glass-section:nth-child(5) {
  animation-delay: 0.25s;
}
.glass-section:nth-child(6) {
  animation-delay: 0.3s;
}
.glass-section:nth-child(7) {
  animation-delay: 0.35s;
}
.glass-section:nth-child(8) {
  animation-delay: 0.4s;
}
.glass-section:nth-child(9) {
  animation-delay: 0.45s;
}
.glass-section:nth-child(10) {
  animation-delay: 0.5s;
}
.glass-section:hover {
  box-shadow:
    0 12px 40px rgba(0, 0, 0, 0.5),
    0 0 40px rgba(219, 39, 119, 0.05);
  transform: translateY(-2px);
}
.glass-card {
  background: rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(12px) saturate(150%);
  -webkit-backdrop-filter: blur(12px) saturate(150%);
  border: 1px solid rgba(255, 255, 255, 0.06);
  transition: all 0.3s ease;
}

/* Section Header */
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
  background: linear-gradient(135deg, #db2777, #f472b6);
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
  font-size: 20px;
  font-weight: 700;
  color: #f8fafc;
  margin: 0;
}
.section-text {
  font-size: 15px;
  color: #94a3b8;
  line-height: 1.85;
  margin: 0 0 12px;
}
.section-text:last-child {
  margin-bottom: 0;
}

/* Lists */
.privacy-list {
  list-style: none;
  padding: 0;
  margin: 16px 0;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.privacy-list li {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  font-size: 15px;
  color: #94a3b8;
  line-height: 1.7;
  padding: 10px 14px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.06);
  transition: all 0.2s ease;
  cursor: default;
}
.privacy-list li:hover {
  background: rgba(255, 255, 255, 0.06);
  transform: translateX(4px);
}
.privacy-list li strong {
  color: #f8fafc;
  font-weight: 600;
}
.privacy-list.simple-list li {
  gap: 10px;
  padding: 8px 14px;
  border-radius: 10px;
}
.list-icon-box {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: transform 0.3s ease;
}
.privacy-list li:hover .list-icon-box {
  transform: rotate(8deg) scale(1.1);
}
.list-icon-box svg {
  width: 20px;
  height: 20px;
  color: white;
}
.list-icon-box.blue {
  background: linear-gradient(135deg, #3b82f6, #60a5fa);
  box-shadow: 0 4px 8px rgba(59, 130, 246, 0.25);
}
.list-icon-box.purple {
  background: linear-gradient(135deg, #8b5cf6, #a855f7);
  box-shadow: 0 4px 8px rgba(139, 92, 246, 0.25);
}
.list-icon-box.green {
  background: linear-gradient(135deg, #10b981, #34d399);
  box-shadow: 0 4px 8px rgba(16, 185, 129, 0.25);
}
.list-icon-box.orange {
  background: linear-gradient(135deg, #f59e0b, #fbbf24);
  box-shadow: 0 4px 8px rgba(245, 158, 11, 0.25);
}
.list-icon-box.pink {
  background: linear-gradient(135deg, #ec4899, #f472b6);
  box-shadow: 0 4px 8px rgba(236, 72, 153, 0.25);
}
.check-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
  margin-top: 2px;
  color: #10b981;
}
.check-icon.shield-check {
  color: #3b82f6;
}
.check-icon.info-check {
  color: #f59e0b;
}

/* Highlight Boxes */
.highlight-box {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px 20px;
  border-radius: 14px;
  margin-bottom: 16px;
  cursor: default;
  transition: all 0.3s ease;
}
.highlight-box.security {
  border: 1px solid rgba(59, 130, 246, 0.2);
  background: rgba(59, 130, 246, 0.08);
}
.highlight-box.security .highlight-box-icon {
  color: #3b82f6;
}
.highlight-box.security:hover {
  border-color: rgba(59, 130, 246, 0.4);
  box-shadow: 0 4px 16px rgba(59, 130, 246, 0.1);
}
.highlight-box.success {
  border: 1px solid rgba(16, 185, 129, 0.2);
  background: rgba(16, 185, 129, 0.08);
}
.highlight-box.success .highlight-box-icon {
  color: #10b981;
}
.highlight-box.success:hover {
  border-color: rgba(16, 185, 129, 0.4);
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.1);
}
.highlight-box.warning {
  border: 1px solid rgba(245, 158, 11, 0.2);
  background: rgba(245, 158, 11, 0.08);
}
.highlight-box.warning .highlight-box-icon {
  color: #f59e0b;
}
.highlight-box.warning:hover {
  border-color: rgba(245, 158, 11, 0.4);
  box-shadow: 0 4px 16px rgba(245, 158, 11, 0.1);
}
.highlight-box-icon {
  width: 24px;
  height: 24px;
  flex-shrink: 0;
}

/* Rights Grid */
.rights-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 14px;
  margin: 16px 0 20px;
}
.right-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 16px;
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(139, 92, 246, 0.15);
  cursor: default;
  transition: all 0.3s ease;
}
.right-card:hover {
  border-color: rgba(139, 92, 246, 0.4);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.3);
  transform: translateY(-4px);
  background: rgba(255, 255, 255, 0.05);
}
.right-card:hover .right-icon {
  transform: rotate(10deg) scale(1.15);
}
.right-icon {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #8b5cf6, #a855f7);
  border-radius: 12px;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);
  transition: transform 0.3s ease;
}
.right-icon svg {
  width: 20px;
  height: 20px;
  color: white;
}
.right-name {
  font-size: 14px;
  font-weight: 700;
  color: #f8fafc;
  margin-bottom: 2px;
}
.right-desc {
  font-size: 12px;
  color: #64748b;
  line-height: 1.4;
}

/* Contact Cards (HOK 暗色) */
.contact-cards {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-top: 16px;
}
.contact-card {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 20px;
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
  cursor: pointer;
  transition: all 0.3s ease;
}
.contact-card:hover {
  border-color: rgba(219, 39, 119, 0.3);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
  transform: translateY(-4px);
  background: rgba(255, 255, 255, 0.05);
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
  background: linear-gradient(135deg, #db2777, #f472b6);
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
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 2px;
}
.contact-value {
  font-size: 14px;
  font-weight: 600;
  color: #f8fafc;
}

/* Footer */
.privacy-footer {
  margin-top: 40px;
  padding-top: 28px;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
  animation: fadeInUp 0.6s ease-out 0.55s backwards;
}
.effective-date {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
  font-size: 13px;
  color: #64748b;
  font-weight: 500;
}
.effective-date svg {
  color: #db2777;
}
.back-link {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background: linear-gradient(135deg, #d97706 0%, #f59e0b 50%, #fbbf24 100%);
  color: #0b0b10;
  text-decoration: none;
  border-radius: 14px;
  font-weight: 700;
  font-size: 14px;
  cursor: pointer;
  box-shadow:
    0 4px 20px rgba(217, 119, 6, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.back-link:hover {
  transform: translateY(-3px);
  box-shadow:
    0 8px 32px rgba(217, 119, 6, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
}

/* Animations */
@keyframes float {
  0%,
  100% {
    transform: translateY(0) scale(1);
  }
  50% {
    transform: translateY(-20px) scale(1.03);
  }
}
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(24px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
@keyframes fadeInLeft {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* Responsive */
@media (max-width: 1024px) {
  .toc-nav {
    display: none;
  }
  .rights-grid {
    grid-template-columns: repeat(2, 1fr);
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
  .privacy-body {
    padding: 32px 16px 48px;
  }
  .glass-section {
    padding: 20px;
  }
  .contact-cards {
    grid-template-columns: 1fr;
  }
  .privacy-footer {
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
  .rights-grid {
    grid-template-columns: 1fr;
  }
}
</style>

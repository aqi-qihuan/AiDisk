import { routes } from "@/router/routes";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior() {
    // 页面切换时滚动到顶部
    return { top: 0, behavior: "smooth" };
  },
});

export default router;

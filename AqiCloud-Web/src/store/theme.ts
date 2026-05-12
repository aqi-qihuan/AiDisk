/**
 * Theme Store - Dark/Light Mode Management
 *
 * Features:
 * - Dark mode as default theme
 * - Light mode opt-in toggle
 * - localStorage persistence
 * - System prefers-color-scheme detection
 * - data-theme attribute on <html> element
 */
import { defineStore } from "pinia";
import { ref, watch } from "vue";

export const useThemeStore = defineStore("theme", () => {
  // State
  const isDark = ref(true); // Default to dark mode (HOK theme)

  /**
   * Apply current theme to DOM
   */
  function applyTheme() {
    if (isDark.value) {
      document.documentElement.setAttribute("data-theme", "dark");
      document.documentElement.removeAttribute("data-theme", "light");
    } else {
      document.documentElement.setAttribute("data-theme", "light");
      document.documentElement.removeAttribute("data-theme", "dark");
    }
  }

  /**
   * Toggle between dark and light mode
   */
  function toggleTheme() {
    isDark.value = !isDark.value;
    applyTheme();
    localStorage.setItem("theme", isDark.value ? "dark" : "light");
  }

  /**
   * Set specific theme mode
   */
  function setTheme(dark: boolean) {
    isDark.value = dark;
    applyTheme();
    localStorage.setItem("theme", dark ? "dark" : "light");
  }

  /**
   * Initialize theme from localStorage or system preference
   */
  function initTheme() {
    const saved = localStorage.getItem("theme");

    if (saved === "light") {
      isDark.value = false;
    } else if (saved === "dark") {
      isDark.value = true;
    } else {
      // No saved preference - check system preference
      const prefersDark = window.matchMedia(
        "(prefers-color-scheme: dark)",
      ).matches;
      isDark.value = prefersDark;
    }

    applyTheme();
  }

  /**
   * Watch for system theme changes
   */
  function watchSystemTheme() {
    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    const handler = (e: MediaQueryListEvent) => {
      // Only auto-switch if user hasn't set a preference
      if (!localStorage.getItem("theme")) {
        isDark.value = e.matches;
        applyTheme();
      }
    };
    mediaQuery.addEventListener("change", handler);
    return () => mediaQuery.removeEventListener("change", handler);
  }

  return {
    isDark,
    applyTheme,
    toggleTheme,
    setTheme,
    initTheme,
    watchSystemTheme,
  };
});

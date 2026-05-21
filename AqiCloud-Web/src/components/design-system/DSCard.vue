<template>
  <div :class="['ds-card', `ds-card--${variant}`]">
    <slot />
  </div>
</template>

<script setup lang="ts">
interface Props {
  variant?: "default" | "hoverable" | "glass";
}

withDefaults(defineProps<Props>(), {
  variant: "default",
});
</script>

<style scoped>
.ds-card {
  background: var(--color-bg-card, #1a1a24);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: var(--radius-lg);
  padding: 24px;
  transition: all var(--transition-base);
  position: relative;
  overflow: hidden;
}

.ds-card--hoverable {
  cursor: pointer;
}

.ds-card--hoverable:hover {
  border-color: rgba(212, 168, 83, 0.3);
  box-shadow: var(--shadow-md), var(--glow-pink);
  transform: translateY(-4px);
}

.ds-card--hoverable:active {
  transform: translateY(-2px);
}

/* Glassmorphism Card - Dark Variant */
.ds-card--glass {
  background: var(--glass-bg);
  backdrop-filter: var(--glass-blur);
  -webkit-backdrop-filter: var(--glass-blur);
  border: var(--glass-border);
  box-shadow: var(--glass-shadow);
}

.ds-card--glass::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(
    90deg,
    transparent,
    var(--color-primary),
    transparent
  );
  opacity: 0;
  transition: opacity var(--transition-base);
}

.ds-card--glass:hover {
  background: var(--glass-bg-hover);
  border: var(--glass-border-hover);
  box-shadow: var(--shadow-lg), var(--glow-pink);
  transform: translateY(-4px);
}

.ds-card--glass:hover::before {
  opacity: 1;
}

@media (prefers-reduced-motion: reduce) {
  .ds-card--hoverable:hover,
  .ds-card--glass:hover {
    transform: none;
  }
}
</style>

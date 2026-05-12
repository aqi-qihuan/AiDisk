<template>
  <div
    :class="[
      'ds-avatar',
      `ds-avatar-${size}`,
      { 'ds-avatar-circle': shape === 'circle' },
    ]"
    :style="avatarStyle"
  >
    <img
      v-if="src && !imageError"
      :src="src"
      :alt="alt"
      class="ds-avatar-image"
      @error="handleImageError"
    />
    <span v-else-if="text" class="ds-avatar-text">
      {{ displayText }}
    </span>
    <span v-else class="ds-avatar-icon">
      <slot>
        <svg viewBox="0 0 24 24" width="60%" height="60%">
          <path
            fill="currentColor"
            d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"
          />
        </svg>
      </slot>
    </span>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";

interface Props {
  src?: string;
  alt?: string;
  size?: "small" | "medium" | "large" | number;
  shape?: "circle" | "square";
  text?: string;
  bgColor?: string;
  textColor?: string;
}

const props = withDefaults(defineProps<Props>(), {
  src: "",
  alt: "",
  size: "medium",
  shape: "circle",
  text: "",
  bgColor: "",
  textColor: "",
});

const imageError = ref(false);

const displayText = computed(() => {
  if (props.text) {
    return props.text.slice(0, 2).toUpperCase();
  }
  return "";
});

const avatarStyle = computed(() => {
  const style: Record<string, string> = {};

  if (typeof props.size === "number") {
    style.width = `${props.size}px`;
    style.height = `${props.size}px`;
  }

  if (props.bgColor) {
    style.backgroundColor = props.bgColor;
  }

  if (props.textColor) {
    style.color = props.textColor;
  }

  return style;
});

const handleImageError = () => {
  imageError.value = true;
};
</script>

<style scoped>
.ds-avatar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background-color: #e0e7ff;
  color: #6366f1;
  overflow: hidden;
  flex-shrink: 0;
  box-sizing: border-box;
}

/* Shapes */
.ds-avatar-circle {
  border-radius: 50%;
}

.ds-avatar[shape="square"] {
  border-radius: 8px;
}

/* Sizes */
.ds-avatar-small {
  width: 32px;
  height: 32px;
  font-size: 12px;
}

.ds-avatar-medium {
  width: 40px;
  height: 40px;
  font-size: 14px;
}

.ds-avatar-large {
  width: 56px;
  height: 56px;
  font-size: 18px;
}

.ds-avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.ds-avatar-text {
  font-weight: 600;
  line-height: 1;
  text-transform: uppercase;
}

.ds-avatar-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

/* Hover effect for clickable avatars */
button .ds-avatar:hover,
a .ds-avatar:hover {
  opacity: 0.85;
}
</style>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';

// Theme state
const theme = ref<'light' | 'dark'>('light');

// Toggle theme function
const toggleTheme = () => {
    theme.value = theme.value === 'light' ? 'dark' : 'light';
};

// Load theme from localStorage on mount
onMounted(() => {
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | null;
    if (savedTheme) {
        theme.value = savedTheme;
    }
    // Apply initial theme
    document.documentElement.classList.toggle('dark', theme.value === 'dark');
});

// Save theme to localStorage and apply to document when it changes
watch(theme, (newTheme) => {
    localStorage.setItem('theme', newTheme);
    document.documentElement.classList.toggle('dark', newTheme === 'dark');
});
</script>

<template>
    <button @click="toggleTheme" 
        class="px-4 py-2 bg-secondary  text-content  rounded-md hover:bg-muted  transition duration-200 flex items-center gap-2 border border-r-2 border-muted">
        {{ theme === 'light' ? '☀️ Light' : '🌙 Dark' }}
    </button>
</template>
<script setup lang="ts">
import { Option } from '@/constants/Options';
defineProps<{
    chip: Option
}>()
</script>

<template>
    <div  :key="chip.id || chip.name" :style="{ background: chip.color }"
        class="px-2 py-1 hover:bg-secondary  cursor-pointer rounded-lg m-2 text-content w-min h-min whitespace-nowrap">
        {{ chip.name }}
    </div>
</template>
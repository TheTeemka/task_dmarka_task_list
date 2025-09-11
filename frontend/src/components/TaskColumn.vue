<script setup lang="ts">
import { inject, readonly } from 'vue';

const props = defineProps<{
    field: string;
    header: string;
    sortable?: boolean;
}>();

const registerColumn = inject<(col: any) => void>('registerColumn');
    if (registerColumn) {
        registerColumn(readonly(props));
    }else{
        console.error("registerColumn not found");
    }

</script>

<template></template>
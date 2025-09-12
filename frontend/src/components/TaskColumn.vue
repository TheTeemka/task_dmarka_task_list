<script setup lang="ts">
import { inject } from 'vue';

export interface Column {
    field: string;
    header: string;
    sortable?: boolean;
    chippable?: boolean;
    editable?: boolean;
    width?: string;  // Add width prop
}

const props = defineProps<Column>();

const registerColumn = inject<(col: Column) => void>('registerColumn');

if (registerColumn) {
    registerColumn(props);
}
</script>

<template>
    <!-- No template needed, as this is just for registration -->
</template>
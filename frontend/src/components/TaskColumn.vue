<script setup lang="ts">
import { inject, onMounted, readonly } from 'vue';

export interface Column {
    field: string;
    header: string;
    sortable?: boolean;
    chippable?: boolean;
}

const props = defineProps<Column>();

const registerColumn = inject<(col: any) => void>('registerColumn');
onMounted(()=>{
    if (registerColumn) {
        console.log(props)
        registerColumn(readonly(props));
    } else {
        console.error("registerColumn not found");
    }
})

</script>

<template></template>
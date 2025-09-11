<script setup lang="ts">
import { inject, onMounted, readonly } from 'vue';

const props = defineProps<{
    field: string;
    header: string;
    chippable?: boolean
}>();

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
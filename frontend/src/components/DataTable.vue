<script setup lang="ts">
import { ref, provide } from 'vue';

interface Column {
    field: string;
    header: string;
    sortable?: boolean;
}

defineProps<{
    tasks: any[];
    loading?: boolean;
}>();

const columns = ref<Column[]>([]);

provide('registerColumn', (col: Column) => {
    console.log(col);
    columns.value.push(col);
});

</script>

<template>
    <slot/>
    <div v-if="loading" class="text-gray-500">Loading...</div>
    <table v-else  class="min-w-full table-auto border-collapse ">
        <thead>
            <tr class="bg-gray-100">
                <th v-for="col in columns" :key="col.field"  class="border border-gray-200 px-4 py-2 cursor-pointer">
                    {{ col.header }}
                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in tasks" :key="item.ID || item.id" class="hover:bg-gray-50">
                <td v-for="col in columns" :key="col.field" class="border border-gray-200 px-4 py-2">
                    {{ item[col.field] }}   
                </td>
            </tr>
            <tr v-if="tasks.length === 0" class="text-gray-500">
                <td :colspan="columns.length" class="border border-gray-300 px-4 py-2 text-center">No data available.</td>
            </tr>
        </tbody>
    </table>
</template>
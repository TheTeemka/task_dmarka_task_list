<script setup lang="ts">
import { ref, provide } from 'vue';
import ChipOption from '@/components/ChipOption.vue';
import { getStatusOptionByID, getPriorityOptionByID } from '@/constants/Options';
interface Column {
    field: string;
    header: string;
    sortable?: boolean;
    chippable?: boolean;
}

defineProps<{
    tasks: any[];
    loading?: boolean;
}>();

const columns = ref<Column[]>([]);

provide('registerColumn', (col: Column) => {
    columns.value.push(col);
});

</script>

<template>
    <slot />
    <div v-if="loading" class="text-gray-500">Loading...</div>
    <table v-else class="min-w-full table-auto border-collapse ">
        <thead>
            <tr class="bg-gray-100">
                <th v-for="col in columns" :key="col.field" class="border border-gray-200 px-4 py-2 cursor-pointer">
                    {{ col.header }}
                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in tasks" :key="item.ID || item.id" class="hover:bg-gray-50">
                <td v-for="col in columns" :key="col.field" class="border border-gray-200 px-4 ">
                    <div v-if="col.chippable">
                        <ChipOption v-if="col.field === 'Status'" :chip="getStatusOptionByID(item[col.field])" />
                        <ChipOption v-else-if="col.field === 'Priority'"
                            :chip="getPriorityOptionByID(item[col.field])" />
                    </div>
                    <div v-else>
                        {{ item[col.field] }}
                    </div>
                </td>
            </tr>
            <tr v-if="tasks.length === 0" class="text-gray-500">
                <td :colspan="columns.length" class="border border-gray-300 px-4 py-2 text-center">No data available.
                </td>
            </tr>
        </tbody>
    </table>
</template>
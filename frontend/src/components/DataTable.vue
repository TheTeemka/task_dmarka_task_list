<script setup lang="ts">
import { ref, provide } from 'vue';
import ChipOption from '@/components/ChipOption.vue';
import { getStatusOptionByID, getPriorityOptionByID } from '@/constants/Options';
import { models } from '@go/models';
import { Column } from '@/components/TaskColumn.vue'

const props = defineProps<{
    filter: models.TaskFilter
    tasks: any[];
    loading?: boolean;
}>();

const columns = ref<Column[]>([]);

provide('registerColumn', (col: Column) => {
    columns.value.push(col);
});

function sort(key: string) {
    const stdKey = key.toLowerCase()
    if (props.filter.SortBy === stdKey) {
        props.filter.SortOrder = props.filter.SortOrder === "asc" ? "desc" : "asc";
    } else {
        props.filter.SortBy = stdKey
        props.filter.SortOrder = "asc";
    }
    console.log(JSON.stringify(props.filter))
};

</script>

<template>
    <slot />
    <div v-if="loading" class="text-gray-500">Loading...</div>
    <table v-else class="min-w-full table-auto border-collapse ">
        <thead>
            <tr class="bg-gray-100">
                <th v-for="col in columns" :key="col.field"
                    :class="['border border-gray-200 px-4 py-2', col.sortable ? 'cursor-pointer hover:bg-gray-200' : '']"
                    @click="col.sortable ? sort(col.field) : null">
                    {{ col.header }}
                    <span
                        :style="{ visibility: (col.sortable && filter.SortBy === col.field.toLowerCase()) ? 'visible' : 'hidden' }"
                        class="ml-2">
                        {{ filter.SortOrder === "asc" ? '↓' : '↑' }}
                    </span>
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
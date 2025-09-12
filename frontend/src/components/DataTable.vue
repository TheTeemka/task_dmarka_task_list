<script setup lang="ts">
import { ref, provide } from 'vue';
import ChipOption from '@/components/ChipOption.vue';
import { getStatusOptionByID, getPriorityOptionByID, statusOptionsWithColor, priorityOptionsWithColor } from '@/constants/Options';
import { models } from '@go/models';
import { Column } from '@/components/TaskColumn.vue'
import SelectOptions from './SelectOptions.vue';

const props = defineProps<{
    filter: models.TaskFilter
    tasks: any[];
    loading?: boolean;
    updateTask: (id: number, task: models.TaskDTO) => Promise<void>;  // New prop

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


const editing = ref<string | null>(null);

const startEdit = (taskId: string, field: string) => {
    editing.value = `${taskId}-${field}`;
};

const stopEdit = async (task: any, field: string, value: any) => {
    task[field] = value;
    editing.value = null;  
    await props.updateTask(task.ID, task);
};

const onEnter = (event: KeyboardEvent, task: any, field: string) => {
    if (event.key === 'Enter') {
        stopEdit(task, field, (event.target as HTMLInputElement).value);
    }
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
                    @click="col.sortable ? sort(col.field) : null" :style="{ width: col.width }">
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
                <td v-for="col in columns" :key="col.field" class="border border-gray-200 px-4 py-2"
                    :style="{ width: col.width }">
                    <!-- Editable Text Fields -->
                    <input v-if="col.editable && !col.chippable && editing === `${item.ID}-${col.field}`"
                        :value="item[col.field]"
                        @blur="stopEdit(item, col.field, ($event.target as HTMLInputElement)?.value || '')"
                        @keydown="onEnter($event, item, col.field)"
                        class="w-full border border-gray-300 rounded px-2 py-1" />
                    <!-- Editable Select for Status -->
                    <SelectOptions
                        v-else-if="col.editable && col.field === 'Status' && editing === `${item.ID}-${col.field}`"
                        :modelValue="getStatusOptionByID(item[col.field])" :options="statusOptionsWithColor"
                        :isOpen="true" 
                        @update:modelValue="stopEdit(item, col.field, $event?.id || null)" />

                    <!-- Editable Select for Priority -->
                    <SelectOptions
                        v-else-if="col.editable && col.field === 'Priority' && editing === `${item.ID}-${col.field}`"
                        :modelValue="getPriorityOptionByID(item[col.field])" :options="priorityOptionsWithColor"
                        :isOpen="true" 
                        @update:modelValue="stopEdit(item, col.field, $event?.id || null)" />

                    <div v-else @click="col.editable ? startEdit(item.ID, col.field) : null"
                        :class="col.editable ? 'cursor-pointer' : ''">
                        <ChipOption v-if="col.chippable"
                            :chip="col.field === 'Status' ? getStatusOptionByID(item[col.field]) : getPriorityOptionByID(item[col.field])" />
                        <span v-else>{{ item[col.field] }}</span>
                    </div>
                </td>
            </tr>
        </tbody>
    </table>
</template>
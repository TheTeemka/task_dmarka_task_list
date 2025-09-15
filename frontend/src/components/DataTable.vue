<script setup lang="ts">
import { ref, provide } from 'vue';
import ChipOption from '@/components/ChipOption.vue';
import { getStatusOptionByID, getPriorityOptionByID, statusOptionsWithColor, priorityOptionsWithColor } from '@/types/Options';
import { models } from '@go/models';
import { Column } from '@/components/TaskColumn.vue'
import SelectOptions from '@/components/SelectOptions.vue';
import  {disablePastDates, formatDate} from '@/utils/dateFormatter'
import Datepicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';


const props = defineProps<{
    filter: models.TaskFilterDTO
    tasks: models.TaskDTO[];
    loading?: boolean;
    updateTask: (id: number, task: models.TaskDTO) => Promise<void>;  // New prop
    deleteTask: (id: number) => Promise<void>;  // Add delete prop
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

const startEdit = (taskId: number, field: string) => {
    editing.value = `${taskId}-${field}`;
};

const stopEdit = async (task: any, field: string, value: any) => {
    task[field] = value;
    editing.value = null;
    await props.updateTask(task.id, task);
};

const onEnter = (event: KeyboardEvent, task: any, field: string) => {
    if (event.key === 'Enter') {
        stopEdit(task, field, (event.target as HTMLInputElement).value);
    }
};

</script>

<template>
    <slot />
    <div v-if="loading" class="text-muted">Loading...</div>
    <table v-else class="min-w-full table-auto border-collapse ">
        <thead>
            <tr class="bg-secondary ">
                <th v-for="col in columns" :key="col.field"
                    class="border border-muted px-2 py-2"
                    :class="[ col.sortable ? 'cursor-pointer hover:bg-muted ' : '']"
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
            <tr v-for="item in tasks" :key="item.id"
                :class="item.status === 3 ?' ' : ''">
                <td v-for="col in columns" :key="col.field" class="border border-muted px-4 py-2"
                    :style="{ width: col.width }" @click="col.editable ? startEdit(item.id, col.field) : null">
                    <!-- Actions Column -->
                    <template v-if="col.field === 'actions'">
                        <button @click="props.deleteTask(item.id)"
                            class="px-2 py-1 bg-error text-content rounded hover:bg-accent">Delete</button>
                    </template>
                    <!-- Editable Text Fields -->
                    <template v-else>
                        <!-- Editable Select for Status -->
                        <SelectOptions
                            v-if="col.editable && col.chippable && col.field === 'status' && editing === `${item.id}-${col.field}`"
                            :modelValue="getStatusOptionByID(item[col.field])" :options="statusOptionsWithColor"
                            :isOpen="true" @update:modelValue="stopEdit(item, col.field, $event?.id || null)" />

                        <!-- Editable Select for Priority -->
                        <SelectOptions
                            v-else-if="col.editable && col.chippable && col.field === 'priority' && editing === `${item.id}-${col.field}`"
                            :modelValue="getPriorityOptionByID(item[col.field])" :options="priorityOptionsWithColor"
                            :isOpen="true" @update:modelValue="stopEdit(item, col.field, $event?.id || null)" />

                        <!-- Editable Date Input for DueDate -->
                        <Datepicker v-else-if="col.editable && col.isDate && editing === `${item.id}-${col.field}`"
                            :model-value="item[col.field as keyof models.TaskDTO]"
                            :enable-time-picker="false" :format="'yyyy-MM-dd'" :text-input="true" class="w-full"
                            :clearable="false"
                            :hide-input-icon="true"
                            :auto-apply="true" @update:model-value="stopEdit(item, col.field, $event)"
                            :disabled-dates="disablePastDates" />

                        <textarea v-else-if="col.editable && editing === `${item.id}-${col.field}`"
                            :value="item[col.field as keyof models.TaskDTO]"     @blur="stopEdit(item, col.field, ($event.target as HTMLTextAreaElement)?.value || '')"

                            @keydown="onEnter($event, item, col.field)" :min-height="40" :max-height="120"
                            class="w-full border border-muted rounded px-2 py-1" />

                        <div v-else :class="col.editable ? 'cursor-pointer' : ''" class="w-full h-full">
                            <ChipOption v-if="col.chippable"
                                :chip="col.field === 'status' ? getStatusOptionByID(item[col.field as keyof models.TaskDTO] as number) : getPriorityOptionByID(item[col.field as keyof models.TaskDTO] as number)" />

                            <span v-else-if="col.isDate">{{ formatDate(item[col.field as keyof models.TaskDTO] as
                                string) }}</span>
                            <span v-else>{{ item[col.field as keyof models.TaskDTO] }}</span>
                        </div>
                    </template>
                </td>
            </tr>
        </tbody>
    </table>
</template>


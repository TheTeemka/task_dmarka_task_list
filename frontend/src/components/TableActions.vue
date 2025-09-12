<script setup lang="ts">
import { ref, watch } from 'vue';
import { models } from '@go/models';
import AddTaskModal from './AddTaskModal.vue';
import SelectOptions from './SelectOptions.vue';
import { Option, priorityOptionsWithColor, statusOptionsWithColor } from '@/constants/Options';
const props = defineProps<{
    filter: models.TaskFilter;
    addTask: (task: models.TaskDTO) => void;
}>();

const filterStatus = ref<Option|null>(null);
watch(filterStatus, (newStatus) => {
    if (newStatus?.id !== null) {
        props.filter.Status = [newStatus?.id as number]
    }else{
        props.filter.Status = []
    }
})

const filterPriority = ref<Option|null>(null);
watch(filterPriority, (newPriority) => {
    if (newPriority?.id !== null) {
        props.filter.Priority = [newPriority?.id as number]
    }else{
        props.filter.Priority = []
    }
})

const search = ref('');
watch(search, (newSearch) => {
    props.filter.Search = newSearch;
})

const showAdvanced = ref(false);

const addPageVisible = ref(false);

</script>

<template>
    <div class="table-actions flex flex-wrap items-center justify-between p-4  border-b border-muted">
        <!-- Add New Task -->
        <div class="flex items-center gap-2">
            <button @click="addPageVisible = !addPageVisible"
                class="px-4 py-2 bg-primary text-content rounded-md hover:bg-primary focus:outline-none focus:ring-2 focus:ring-primary">
                Add Task
            </button>
            <AddTaskModal :visible="addPageVisible" @add="addTask($event)" @close="addPageVisible = false" />
        </div>

        <div class="flex items-center gap-2">
            <input v-model="search" type="text" placeholder="Search tasks..." class="px-3 py-2 border border-muted rounded-md focus:outline-none focus:ring-2 focus:ring-primary" />
            <span class="text-lg font-bold mr-10">Filters</span>
            <SelectOptions v-model="filterStatus" :options="statusOptionsWithColor" />
            <SelectOptions v-model="filterPriority" :options="priorityOptionsWithColor" />
        </div>
    </div>
</template>
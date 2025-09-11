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
    props.filter.Status = newStatus?.id
})

const filterPriority = ref<Option|null>(null);
watch(filterPriority, (newPriority) => {
    props.filter.Priority = newPriority?.id
})

const addPageVisible = ref(false);

</script>

<template>
    <div class="table-actions flex flex-wrap items-center gap-4 p-4 bg-gray-50 border-b border-gray-200">
        <!-- Add New Task -->
        <div class="flex items-center gap-2">
            <button @click="addPageVisible = !addPageVisible"
                class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
                Add Task
            </button>
            <AddTaskModal :visible="addPageVisible" @add="addTask($event)" @close="addPageVisible = false" />
        </div>

        <div class="flex items-center gap-2">
            <SelectOptions v-model="filterStatus"  :options="statusOptionsWithColor"/>
            <SelectOptions v-model="filterPriority" :options="priorityOptionsWithColor" />
        </div>

    </div>
</template>
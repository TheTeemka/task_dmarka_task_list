<script setup lang="ts">
import { ref } from 'vue';
import { models } from '@go/models';

const props = defineProps<{
    onAdd: (task: models.TaskDTO) => void;
    onFilter: (filters: models.TaskFilter) => void;
    onSort: (sortBy: string, sortOrder: string) => void;
    onSearch: (query: string) => void;
}>();

const newTaskTitle = ref('');
const searchQuery = ref('');
const filterPriority = ref('');
const sortBy = ref('Title');
const sortOrder = ref('asc');

const addTask = () => {
    if (!newTaskTitle.value.trim()) return;
    const task = new models.TaskDTO({
        Title: newTaskTitle.value.trim(),
        Description: '',
        Status: 'pending',
        Priority: 'low',
    });
    props.onAdd(task);
    newTaskTitle.value = '';
};

const applyFilter = () => {
    const filters = new models.TaskFilter({
        PriorityID: filterPriority.value ? parseInt(filterPriority.value) : undefined,
    });
    props.onFilter(filters);
};

const applySort = () => {
    props.onSort(sortBy.value, sortOrder.value);
};

const performSearch = () => {
    props.onSearch(searchQuery.value);
};
</script>

<template>
    <div class="table-actions flex flex-wrap items-center gap-4 p-4 bg-gray-50 border-b border-gray-200">
        <!-- Add New Task -->
        <div class="flex items-center gap-2">
            <input
                v-model="newTaskTitle"
                type="text"
                placeholder="Add new task..."
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                @keyup.enter="addTask"
            />
            <button
                @click="addTask"
                class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
                Add
            </button>
        </div>

        <!-- Search -->
        <div class="flex items-center gap-2">
            <input
                v-model="searchQuery"
                type="text"
                placeholder="Search tasks..."
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                @input="performSearch"
            />
        </div>

        <!-- Filter -->
        <div class="flex items-center gap-2">
            <select
                v-model="filterPriority"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                @change="applyFilter"
            >
                <option value="">All Priorities</option>
                <option value="0">Low</option>
                <option value="1">Medium</option>
                <option value="2">High</option>
                <option value="3">Urgent</option>
            </select>
        </div>

        <!-- Sort -->
        <div class="flex items-center gap-2">
            <select
                v-model="sortBy"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                @change="applySort"
            >
                <option value="Title">Title</option>
                <option value="Status">Status</option>
                <option value="Priority">Priority</option>
                <option value="CreatedDate">Created Date</option>
            </select>
            <button
                @click="sortOrder = sortOrder === 'asc' ? 'desc' : 'asc'; applySort()"
                class="px-3 py-2 border border-gray-300 rounded-md hover:bg-gray-100 focus:outline-none"
            >
                {{ sortOrder === 'asc' ? '↑' : '↓' }}
            </button>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { CreateNewTask, GetListByFilters, UpdateTask } from '@go/service/TaskService';
import { models } from '@go/models';
import DataTable from '@/components/DataTable.vue';
import Column from '@/components/TaskColumn.vue';
import TableActions from '@/components/TableActions.vue';
import ThemeToggle from '@/components/ThemeToggle.vue';  // New import

const tasks = ref<models.TaskDTO[]>([]);
const loading = ref<boolean>(false);
const error = ref<string | null>(null);
const filter = ref<models.TaskFilter>(new models.TaskFilter({}));

const fetchTasks = async () => {
    loading.value = true;
    error.value = null;
    try {
        const result = await GetListByFilters(filter.value);
        if (result === null) {
            tasks.value = []
        } else {
            tasks.value = result;
        }
    } catch (e) {
        alert(e)
        error.value = 'Failed to fetch tasks';
    } finally {
        console.log(tasks.value)
        loading.value = false;
    }
};

watch(filter, () => {
    fetchTasks();
}, { deep: true });


const addTask = async (task: models.TaskDTO) => {
    await CreateNewTask(task);
    await fetchTasks();
}

const updateTask = async (id: number, updatedTask: models.TaskDTO) => {
    try {
        await UpdateTask(id, updatedTask);
        await fetchTasks();
    } catch (e) {
        alert('Failed to update task: ' + e);
    }
};

onMounted(fetchTasks);
</script>

<template>
    <div class="task-manager ">
        <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-bold text-content">Task Manager</h2>
            <ThemeToggle /> <!-- Add the theme toggle here -->
        </div>
        <TableActions :addTask="addTask" :filter="filter" />
        <DataTable :tasks="tasks" :loading="loading" :filter="filter" :updateTask="updateTask">
            <Column field="ID" header="ID" sortable width="4px" />
            <Column field="Title" header="Title" sortable editable width="20%" />
            <Column field="Status" header="Status" chippable sortable editable width="10%" />
            <Column field="Priority" header="Priority" chippable sortable editable width="10%" />
            <Column field="Description" header="Description" editable />
        </DataTable>
    </div>
</template>
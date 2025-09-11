<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { CreateNewTask, GetListByFilters } from '@go/service/TaskService';
import { models } from '@go/models';
import DataTable from '@/components/DataTable.vue';
import Column from '@/components/TaskColumn.vue';
import TableActions from '@/components/TableActions.vue';

const tasks = ref<models.TaskDTO[]>([]);
const loading = ref<boolean>(false);
const error = ref<string | null>(null);

const fetchTasks = async () => {
    loading.value = true;
    error.value = null;
    try {
        const filter = new models.TaskFilter({});
        const result = await GetListByFilters(filter);
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


const fetchTasksWithFilter = async () => {
    await fetchTasks();
}

const sortTasks = async () => {
    await fetchTasks();
}

const searchTasks = async () => {
    await fetchTasks();
}

const addTask = async (task: models.TaskDTO) => {
   await CreateNewTask(task);
   await fetchTasks();
}

onMounted(fetchTasks);
</script>

<template>
    <div class="task-manager">
        <h2 class="text-xl font-bold mb-4">Task Manager</h2>
        <TableActions :onAdd="addTask" :onFilter="fetchTasksWithFilter" :onSort="sortTasks" :onSearch="searchTasks" />
        <DataTable :tasks="tasks" :loading="loading">
            <Column field="ID" header="ID" />
            <Column field="Title" header="Title" />
            <Column field="Status" header="Status" chippable/>
            <Column field="Priority" header="Priority" chippable/>
        </DataTable>
    </div>
</template>
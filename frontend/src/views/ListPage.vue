<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { CreateNewTask, GetListByFilters } from '@go/service/TaskService';
import { models } from '@go/models';
import DataTable from '@/components/DataTable.vue';
import Column from '@/components/TaskColumn.vue';

const newTask = ref<string>('');
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
        }else{
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

const addTask = async () => {
    if (!newTask.value.trim()) return;
    try {
        const dto = new models.TaskDTO({
            Title: newTask.value.trim(),
            Description: '',
            Status: 'pending',
            Priority: 0,
        });
        await CreateNewTask(dto);
        newTask.value = '';
        await fetchTasks();
    } catch (e) {
        alert(e);
        error.value = 'Failed to add task';
    }
};

onMounted(fetchTasks);
</script>

<template>
    <div class="task-manager">
        <h2 class="text-xl font-bold mb-4">Task Manager</h2>
        <form @submit.prevent="addTask" class="mb-4 flex gap-2">
            <input v-model="newTask" type="text" placeholder="Enter new task" class="border rounded px-2 py-1 flex-1"
                required />
            <button type="submit" class="bg-blue-500 text-white px-4 py-1 rounded">Add</button>
        </form>
        <div v-if="error" class="text-red-500 mb-4">{{ error }}</div>
        <DataTable :tasks="tasks" :loading="loading" >
            <Column field="ID" header="ID" />
            <Column field="Title" header="Title"  />
            <Column field="Status" header="Status"  />
            <Column field="Priority" header="Priority" />
        </DataTable>
        <div v-if="tasks.length === 0 && !loading" class="text-gray-500 mt-4">No tasks yet.</div>
    </div>
</template>
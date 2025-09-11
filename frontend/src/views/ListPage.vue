<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { CreateNewTask, GetListByFilters } from '@go/service/TaskService';
import { models } from '@go/models';
import DataTable from '@/components/DataTable.vue';
import Column from '@/components/TaskColumn.vue';
import TableActions from '@/components/TableActions.vue';

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

onMounted(fetchTasks);
</script>

<template>
    <div class="task-manager">
        <h2 class="text-xl font-bold mb-4">Task Manager</h2>
        <TableActions :addTask="addTask" :filter="filter" />
        <DataTable :tasks="tasks" :loading="loading" :filter="filter">
            <Column field="ID" header="ID" sortable/>
            <Column field="Title" header="Title" sortable/>
            <Column field="Status" header="Status" chippable sortable/>
            <Column field="Priority" header="Priority" chippable sortable/>
        </DataTable>
    </div>
</template>
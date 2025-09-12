<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { CreateNewTask, GetListByFilters, UpdateTask, DeleteTask } from '@go/service/TaskService';
import { models } from '@go/models';
import DataTable from '@/components/DataTable.vue';
import Column from '@/components/TaskColumn.vue';
import TableActions from '@/components/TableActions.vue';
import ThemeToggle from '@/components/ThemeToggle.vue';  // New import
import ConfirmModal from '@/components/ConfirmModal.vue';

const tasks = ref<models.TaskDTO[]>([]);
const loading = ref<boolean>(false);
const error = ref<string | null>(null);
const filter = ref<models.TaskFilter>(new models.TaskFilter({}));

// Confirm modal state
const confirmModal = ref({
    visible: false,
    title: '',
    message: '',
    onConfirm: () => { },
});

const fetchTasks = async () => {
    loading.value = true;
    error.value = null;
    try {
        const result = await GetListByFilters(filter.value);
        console.log(result);
        if (result === null) {
            tasks.value = []
        } else {
            tasks.value = result;
        }
    } catch (e) {
        alert(e)
        error.value = 'Failed to fetch tasks';
    } finally {
        loading.value = false;
    }
};

watch(filter, () => {
    fetchTasks();
}, { deep: true });


const addTask = async (task: models.TaskDTO) => {
    console.log(task)
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

const deleteTask = async (id: number) => {
    confirmModal.value = {
        visible: true,
        title: 'Delete Task',
        message: 'Are you sure you want to delete this task? This action cannot be undone.',
        onConfirm: async () => {
            try {
                await DeleteTask(id);
                await fetchTasks();
            } catch (e) {
                alert('Failed to delete task: ' + e);
            }
        },
    };
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
        <DataTable :tasks="tasks" :loading="loading" :filter="filter" :updateTask="updateTask" :deleteTask="deleteTask">
            <Column field="id" header="ID" sortable width="4px" />
            <Column field="title" header="Title" sortable editable width="20%" />
            <Column field="status" header="Status" chippable sortable editable width="10%" />
            <Column field="priority" header="Priority" chippable sortable editable width="10%" />
            <Column field="created_at" header="Created Date" sortable  isDate width="12%" />
            <Column field="due_date" header="Due Date" sortable editable isDate width="12%" />
            <Column field="description" header="Description" editable />
            <Column field="actions" header="Actions" width="10%" />
        </DataTable>

        <ConfirmModal :visible="confirmModal.visible" :title="confirmModal.title" :message="confirmModal.message"
            @confirm="confirmModal.onConfirm()" @close="confirmModal.visible = false" />
    </div>
</template>
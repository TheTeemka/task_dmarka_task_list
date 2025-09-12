<script setup lang="ts">
import { ref, watch } from 'vue';
import { models } from '@go/models';
import { getStatusOptionByID, getPriorityOptionByID, statusOptionsWithColor, priorityOptionsWithColor, Option } from '@/constants/Options';
import SelectOptions from './SelectOptions.vue';

defineProps<{
    visible: boolean;
}>();

const emit = defineEmits<{
    add: [task: models.TaskDTO];
    close: [];
}>();



const title = ref<string>("");
const description = ref<string>('');
const status = ref<Option>(getStatusOptionByID(0));
const priority = ref<Option>(getPriorityOptionByID(0));

const addTask = () => {
    if (!title.value.trim()) return;

    const taskData = {
        Title: title.value.trim(),
        Description: description.value.trim(),
        Status: status.value?.id as number,
        Priority: priority.value?.id as number,
    };

    // Validate before creating (optional, but improves safety)
    if (!taskData.Status) {
        alert('Please select a status');
        return;
    }

    const task = new models.TaskDTO(taskData);
    emit('add', task);
    resetForm();
    emit('close');
};

const resetForm = () => {
    title.value = '';
    description.value = '';
    status.value = getStatusOptionByID(0);
    priority.value = getPriorityOptionByID(0);
};

const closeModal = () => {
    resetForm();
    emit('close');
};

watch(status, () => {
    console.log(status)
})

</script>

<template>
    <div v-if="visible" class="fixed inset-0  bg-black/20  flex items-center justify-center z-50">
        <div class="bg-white rounded-lg shadow-lg p-6 w-full max-w-xl">
            <h2 class="text-xl font-bold mb-4">Add New Task</h2>
            <form @submit.prevent="addTask">
                <div class="mb-4">
                    <label class="block text-sm font-medium mb-1">Title</label>
                    <input v-model="title" type="text" required
                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500" />
                </div>
                <div class="flex space-x-3">
                    <div class="mb-4">
                        <label class="block text-sm font-medium mb-1">Status</label>
                        <SelectOptions v-model="status" :options="statusOptionsWithColor" />
                    </div>
                    <div class="mb-4">
                        <label class="block text-sm font-medium mb-1">Priority</label>
                        <SelectOptions v-model="priority" :options="priorityOptionsWithColor" />
                    </div>
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium mb-1">Description</label>
                    <textarea v-model="description"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        rows="5"></textarea>
                </div>
                <div class="flex justify-end gap-2">
                    <button type="button" @click="closeModal"
                        class="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-100 focus:outline-none">
                        Cancel
                    </button>
                    <button type="submit"
                        class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
                        Add Task
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>
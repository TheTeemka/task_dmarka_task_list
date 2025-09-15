<script setup lang="ts">
import { ref } from 'vue';
import { models } from '@go/models';
import { getStatusOptionByID, getPriorityOptionByID, statusOptionsWithColor, priorityOptionsWithColor, Option } from '@/types/Options';
import SelectOptions from './SelectOptions.vue';
import Datepicker from '@vuepic/vue-datepicker';
import { formatDate, disablePastDates } from '@/utils/dateFormatter';
import '@vuepic/vue-datepicker/dist/main.css';

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
const dueDate = ref<Date | null>(null);

const addTask = () => {
    if (!title.value.trim()) return;

    const task: models.TaskDTO = {
        id: 0, // Will be set by backend
        title: title.value.trim(),
        description: description.value.trim(),
        status: status.value?.id ?? 0,
        priority: priority.value?.id ?? 0,
        due_date: dueDate.value ? dueDate.value.toISOString() : new Date().toISOString(),
        completed_date: '',
        created_at: '',
    };


    console.log(task, task)
    emit('add', task);
    resetForm();
    emit('close');
};

const resetForm = () => {
    title.value = '';
    description.value = '';
    status.value = getStatusOptionByID(0);
    priority.value = getPriorityOptionByID(0);
    dueDate.value = null;
};

const closeModal = () => {
    resetForm();
    emit('close');
};


// Custom format function for Datepicker
const formatDateForPicker = (date: Date): string => {
    return formatDate(date);
};


</script>

<template>
    <div v-if="visible" class="fixed inset-0  bg-black/20  flex items-center justify-center z-50">
        <div class="bg-surface  rounded-lg shadow-lg p-6 w-full max-w-xl">
            <h2 class="text-xl font-bold mb-4">Add New Task</h2>
            <form @submit.prevent="addTask">
                <div class="mb-4">
                    <label class="block text-sm font-medium mb-1">Title</label>
                    <input v-model="title" type="text" required
                        class="w-full px-3 py-2 border border-muted rounded-md focus:outline-none focus:ring-2 focus:ring-primary" />
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

                    <div class="mb-4 h-full">
                        <label class="block text-sm font-medium mb-1">Due Date</label>
                        <Datepicker v-model="dueDate" :enable-time-picker="false" :format="formatDateForPicker"
                            placeholder="Select due date" :text-input="true" :clearable="true" :auto-apply="true" :disabled-dates="disablePastDates"
                            required />
                    </div>
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium mb-1">Description</label>
                    <textarea v-model="description"
                        class="w-full px-3 py-2 border border-muted rounded-md focus:outline-none focus:ring-2 focus:ring-primary"
                        rows="5"></textarea>
                </div>

                <div class="flex justify-end gap-2">
                    <button type="button" @click="closeModal"
                        class="px-4 py-2 border border-muted rounded-md hover:bg-secondary  focus:outline-none">
                        Cancel
                    </button>
                    <button type="submit"
                        class="px-4 py-2 bg-primary text-content rounded-md hover:bg-primary focus:outline-none focus:ring-2 focus:ring-primary">
                        Add Task
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>
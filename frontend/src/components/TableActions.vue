<script setup lang="ts">
import { ref, watch } from 'vue';
import { models } from '@go/models';
import AddTaskModal from './AddTaskPage.vue';
import SelectOptions from './SelectOptions.vue';
import { getPriorityOptionByID, Option, priorityOptionsWithColor, statusOptionsWithColor } from '@/types/Options';
import Datepicker from '@vuepic/vue-datepicker';
import { formatDate } from '@/utils/dateFormatter';
import '@vuepic/vue-datepicker/dist/main.css';
const props = defineProps<{
    filter: models.TaskFilterDTO;
    addTask: (task: models.TaskDTO) => void;
}>();

const filterStatus = ref<Option | null>(getPriorityOptionByID(0));
watch(filterStatus, (newStatus) => {
    if (newStatus?.id !== null) {
        props.filter.Status = [newStatus?.id as number]
    } else {
        props.filter.Status = []
    }
})

const filterPriority = ref<Option>(getPriorityOptionByID(0));
watch(filterPriority, (newPriority) => {
    console.log(newPriority)
    if (newPriority?.id !== null) {
        props.filter.Priority = [newPriority?.id as number]
    } else {
        props.filter.Priority = []
    }
})

const filterBefore = ref<Date>();
watch(filterBefore, (newBefore) => {
    console.log(newBefore)
    if (newBefore) {
        props.filter.Before = newBefore.toISOString()
    } else {
        props.filter.Before = undefined
    }
})

const search = ref('');
watch(search, (newSearch) => {
    props.filter.Search = newSearch;
})


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
            <span class="text-lg font-bold mr-10">Filters</span>
            <input v-model="search" type="text" placeholder="Search tasks..."
                class="px-3 py-2 border bg-surface border-muted rounded-md focus:outline-none focus:ring-2 focus:ring-primary" />
            <Datepicker v-model="filterBefore" :placeholder="'DueDate Before'"  :text-input="true"  :clearable="true" :auto-apply="true"   :enable-time-picker="false" :format="formatDate" class="bg-surface "/>
            <SelectOptions v-model="filterStatus" :options="statusOptionsWithColor" />
            <SelectOptions v-model="filterPriority"     :options="priorityOptionsWithColor" />
        </div>
    </div>
</template>
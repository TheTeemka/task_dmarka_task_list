<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { Option } from '@/constants/Options';
import ChipOption from '@/components/ChipOption.vue';

const props = defineProps<{
    modelValue: Option | null;
    options: Option[];
    placeholder?: string;
}>();

const emit = defineEmits<{
    'update:modelValue': [value: any];
}>();

const isOpen = ref(false);
const dropdownRef = ref<HTMLElement | null>(null);  // Ref for the dropdown container

const selectedOption = computed((): Option => {
    const option = props.options.find(opt => opt.name === props.modelValue?.name);
    return option ? option : { name: props.placeholder || 'Select...' };
});

const selectOption = (value: Option) => {
    emit('update:modelValue', value);
    isOpen.value = false;
};

const toggleDropdown = () => {
    isOpen.value = !isOpen.value;
};

const handleClickOutside = (event: Event) => {
    if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
        isOpen.value = false;
    }
};


onMounted(() => {
    document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside);
});
</script>

<template>
    <div class="relative" ref="dropdownRef">
        <div @click="toggleDropdown"
            class="w-full  border border-gray-300 rounded-md bg-white cursor-pointer focus:outline-none focus:ring-2 focus:ring-blue-500 flex justify-between items-center">
            <ChipOption :chip="selectedOption" @click="selectOption(selectedOption)" />

        </div>
        <div v-if="isOpen" class="absolute z-10 mt-1 w-full bg-white border border-gray-300 rounded-md shadow-lg">
            <div v-for="option in options" :key="option.id || option.name">
                <ChipOption :chip="option" @click="selectOption(option)" />
            </div>
        </div>
    </div>
</template>
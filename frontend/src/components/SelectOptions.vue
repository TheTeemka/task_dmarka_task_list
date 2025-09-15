<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { Option } from '@/types/Options';
import ChipOption from '@/components/ChipOption.vue';

const props = defineProps<{
    modelValue: Option | null;
    options: Option[];
    placeholder?: string;
    isOpen?: boolean;
}>();

const emit = defineEmits<{
    'update:modelValue': [value: any];
}>();

const isOpen = ref( true);
const dropdownRef = ref<HTMLElement | null>(null);  // Ref for the dropdown container
const childRef = ref<HTMLElement | null>(null);  // Ref for the child dropdown div

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

// Function to set parent width to match child
const setParentWidth = () => {

    if (dropdownRef.value && childRef.value) {
        const childWidth = childRef.value.offsetWidth;
        dropdownRef.value.style.width = `${childWidth}px`;
    }
};


onMounted(async () => {
    isOpen.value = props.isOpen || false
    setParentWidth();
    setTimeout(() => {
        document.addEventListener('click', handleClickOutside);
    }, 0);
});

onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside);
});
</script>

<template>
    <div class="relative" ref="dropdownRef">  
        <div @click="toggleDropdown"
            class="w-full border border-muted rounded-md bg-surface  cursor-pointer focus:outline-none focus:ring-2 focus:ring-primary flex justify-between items-center">
            <ChipOption :chip="selectedOption" @click="selectOption(selectedOption)" />
        </div>
        <div  ref="childRef" class="absolute z-10 mt-1 bg-surface  border border-muted rounded-md shadow-lg" :style="{visibility: isOpen ? 'visible' : 'hidden'}">
            <div v-for="option in options" :key="option.id || option.name">
                <ChipOption :chip="option" @click="selectOption(option)" />
            </div>
        </div>
    </div>
</template>
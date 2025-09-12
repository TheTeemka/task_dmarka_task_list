<script setup lang="ts">
defineProps<{
  visible: boolean;
  title?: string;
  message?: string;
}>();

const emit = defineEmits<{
  confirm: [];
  cancel: [];
  close: [];
}>();

const handleConfirm = () => {
  emit('confirm');
  emit('close');
};

const handleCancel = () => {
  emit('cancel');
  emit('close');
};

const handleOverlayClick = (event: MouseEvent) => {
  if (event.target === event.currentTarget) {
    emit('close');
  }
};
</script>

<template>
  <Teleport to="body">
    <div
      v-if="visible"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/15"
      @click="handleOverlayClick"
    >
      <div class="bg-surface border border-muted rounded-lg shadow-lg p-6 max-w-md w-full mx-4">
        <div class="mb-4">
          <h3 v-if="title" class="text-lg font-bold text-content mb-2">{{ title }}</h3>
          <p v-if="message" class="text-content">{{ message }}</p>
        </div>

        <div class="flex justify-end gap-3">
          <button
            @click="handleCancel"
            class="px-4 py-2 bg-muted text-content rounded hover:bg-accent transition-colors"
          >
            Cancel
          </button>
          <button
            @click="handleConfirm"
            class="px-4 py-2 bg-primary text-content rounded hover:bg-accent transition-colors"
          >
            Confirm
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>
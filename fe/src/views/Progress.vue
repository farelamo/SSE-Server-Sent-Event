<template>
  <div>
    <h2>Progress: {{ progress }}%</h2>
    <progress :value="progress" max="100"></progress>
  </div>
</template>

<script setup>
import { onMounted, computed, watch } from 'vue';
import { useProgressStore } from '@/stores/progress';

const store = useProgressStore();
const progress = computed(() => store.value);

watch(progress, (newVal) => {
  console.log('Progress updated:', newVal);
});

onMounted(() => {
  const source = new EventSource('http://localhost:8080/progress');

  source.onmessage = (event) => {
    const val = Number(event.data);
    console.log('Progress event received:', val);
    store.setProgress(val);
    if (val >= 100) {
      source.close();
    }else if (val >= 50) {
      source.close();
    }
  };

  source.onerror = (err) => {
    console.error('SSE error:', err);
    source.close();
  };
});
</script>

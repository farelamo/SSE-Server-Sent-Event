// stores/progressStore.js
import { defineStore } from 'pinia';

export const useProgressStore = defineStore('progress', {
  state: () => ({
    value: 0
  }),
  actions: {
    setProgress(val) {
      this.value = val;
    }
  }
});

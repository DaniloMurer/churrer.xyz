import {defineStore} from "pinia";
import {ref} from "vue";

export const useVisibleStore = defineStore('visibleStore', () => {
    const visible = ref(false);
    return { visible }
})
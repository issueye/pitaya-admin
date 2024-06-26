import { ref } from "vue";

import { defineStore } from "pinia";

export const useUserStore = defineStore(
    'userStore',
    () => {
        const info = ref({})
        const token = ref('')


        return {
            info,
            token,
        }
    },
    {
        persist: {
            key: 'user_store',
            storage: localStorage,
        },
    }
)
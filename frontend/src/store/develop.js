import { ref } from "vue";
import { apiGetAllTable, apiGetTableInfo } from "@/apis/develop/index";


import { defineStore } from "pinia";

export const useDevelopStore = defineStore(
    'develop',
    () => {
        // 当前步骤
        const active = ref(0);
        // 所有表
        const tables = ref([]);
        // 表结构
        const columns = ref([]);

        // 选中的表
        const selectTable = ref('');

        // 窗体显示 visible
        const visible = ref(false);

        /**
         * 获取所有表
         */
        const getTables = async () => {
            let res = await apiGetAllTable();
            if (res.code === 200) {
                tables.value = res.data;
            }
        }

        /**
         * 获取表结构
         * @param {*} name 
         */
        const getColumns = async () => {
            let res = await apiGetTableInfo(selectTable.value);
            if (res.code === 200) {
                columns.value = res.data.fields;
            }
        }

        /**
         * 重置数据
         */
        const resetData = () => {
            active.value = 0;
            // tables.value = [];
            columns.value = [];
            selectTable.value = '';
        }

        return {
            visible,
            active,
            tables,
            columns,
            selectTable,

            resetData,
            getTables,
            getColumns,
        }
    },
    {
        persist: {
            key: 'develop_store',
            storage: localStorage,
        },
    }
)
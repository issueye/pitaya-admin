<template>
<div class="bs-header-box">
    <div>
        <span class="bs-header-title">{{ title }}</span>
        <el-divider direction="vertical" />
        <span class="bs-header-description">{{ description }}</span>
    </div>
    <div class="bs-header-actions">
        <slot name="actions"></slot>
    </div>
</div>
</template>

<script setup>
import { toRefs } from 'vue';

const props = defineProps({
    title: {
        type: String,
        default: '',
    },
    description: {
        type: String,
        default: '',
    },
})

</script>

<style lang="scss" scoped>
.bs-header-box {
    border-radius: 2px 2px 0 0;
    width: 100%;
    height: 45px;
    display: inline-flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fff;
    border-bottom: 1px solid #D9D9D9;

    .bs-header-actions{
        margin-right: 10px;
    }

    .bs-header-title {
        margin-left: 20px;
        line-height: 30px;
        font-size: 18px;
        font-weight: bold;
        vertical-align: middle;
        color: #262626;
    }

    .bs-header-description{
        color: #BFBFBF;
        font-size: 11px;
    }
}
</style>

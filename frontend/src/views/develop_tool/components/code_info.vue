<template>
  <div class="h-[95%] border border-solid border-gray-300">
    <Codemirror v-mode:value="code" :options="options" ref="cmRef" />
  </div>
</template>

<script setup>
import "codemirror/mode/javascript/javascript.js";
import Codemirror from "codemirror-editor-vue3";
import { onMounted, ref } from "vue";

onMounted(() => {
  cmRef.value.refresh;
});

const code = ref(`

`);

const cmRef = ref(null);

const options = {
  tabSize: 4,
  mode: "text/javascript",
  theme: "neat",
  lineNumbers: true,
  line: true,
};
</script>
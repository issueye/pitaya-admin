<template>
  <div class="flex justify-center items-center h-screen bg-gray-200">
    <div
      class="flex justify-center items-center w-[600px] h-[309px] bg-white rounded-lg shadow-lg"
    >
      <div>
        <img
          class="rounded-lg shadow-lg"
          src="@/assets/images/flower-login.webp"
          alt=""
          width="300px"
          height="339px"
        />
      </div>
      <div class="flex flex-grow flex-col items-center">
        <div class="text-4xl font-bold text-gray-800 mb-10">pitaya</div>
        <div class="flex flex-col justify-center items-center">
          <el-form
            :rules="rules"
            ref="formRef"
            :model="dataForm"
            label-width="auto"
          >
            <el-form-item label="用户" prop="account">
              <el-input placeholder="请输入账户" v-model="dataForm.account" />
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input
                placeholder="请输入密码"
                type="password"
                v-model="dataForm.password"
                @keyup.enter="loginClick"
              />
            </el-form-item>
          </el-form>
          <el-button
            type="primary"
            class="w-[120px] mt-4 rounded-md hover:bg-blue-600"
            @click="loginClick"
            >登录</el-button
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import { apiLogin } from "../apis/sys/user";
import { useUserStore } from "@/store/user";
import { storeToRefs } from "pinia";

const route = useRouter();
const userStore = useUserStore();
let { info, token } = storeToRefs(userStore);

const dataForm = reactive({
  account: "",
  password: "",
});

onMounted(() => {
  localStorage.clear();
  sessionStorage.clear();
});

const formRef = ref();

const rules = reactive({
  account: [
    {
      required: true,
      message: "请输入账户",
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
      trigger: "blur",
    },
  ],
});

const loginClick = async () => {
  formRef.value.validate(async (valid) => {
    if (valid) {
      let res = await apiLogin(dataForm);
      if (res.code === 200) {
        localStorage.setItem("token", res.data.token);
        info.value = res.data;
        token.value = res.data.token;
        route.push("/");
      }
    }
  });
};
</script>

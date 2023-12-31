<template>
    <div class="login-card">
        <n-icon size="20" class="card-close" v-show="props.close" @click="closeClick">
            <close></close>
        </n-icon>
        <div class="card-left">
            <div class="illustrations">
                <login-illustration></login-illustration>
            </div>
        </div>
        <div class="card-right">
            <transition name="fade" mode="out-in">
                <div v-if="isLogin" key="login">
                    <login-form @change-form="isLogin = false" @success="success"></login-form>
                </div>

                <div v-else key="register">
                    <register-form @change-form="isLogin = true"></register-form>
                </div>
            </transition>
        </div>
        <p class="protocol-container">
            <span>注册或登录即代表同意 </span>
            <span class="protocol">《用户协议》</span>
            <span>和</span>
            <span class="protocol">《隐私政策》</span>
        </p>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { NIcon } from 'naive-ui';
import { Close } from '@leaf/icons';
import LoginIllustration from "./component/LoginIllustration.vue";
import LoginForm from './component/LoginForm.vue'
import RegisterForm from './component/RegisterForm.vue'

const emits = defineEmits(["close","success"]);
const props = withDefaults(defineProps<{
    close?: boolean,
}>(), {
    close: false
});

// 是否为登录
const isLogin = ref(true);

const closeClick = () => {
    emits("close");
}

const success = () => {
    emits("success");
}
</script>

<style lang="less" scoped>

.login-card {
    position: absolute;
    display: grid;
    grid-template-columns: 330px auto;
    grid-template-rows: auto 60px;
    top: 50%;
    left: 50%;
    margin: -220px 0 0 -410px;
    width: 820px;
    height: 400px;
    overflow: hidden;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 16px 16px 50px -12px rgba(0, 0, 0, 0.8);

    .card-close {
        right: 26px;
        top: 16px;
        cursor: pointer;
        position: absolute;

        &:hover {
            color: #808080;
        }
    }

    // 插画部分
    .card-left {
        width: 330px;
        grid-row: 1 / span 2;
        background-color: #63b5fb;
        -webkit-clip-path: polygon(98% 17%, 100% 34%, 98% 51%, 100% 68%, 98% 84%, 100% 100%, 0 100%, 0 0, 100% 0);
        clip-path: polygon(98% 17%, 100% 34%, 98% 51%, 100% 68%, 98% 84%, 100% 100%, 0 100%, 0 0, 100% 0);

        .illustrations {
            height: 100%;
            padding: 100px 30px 0 30px;
        }
    }

    // 表单部分
    .card-right {
        padding: 40px 30px 0 0;
    }

    //用户协议
    .protocol-container {
        grid-column: 2;
        color: #9499a0;
        text-align: center;
        font-size: 12px;

        .protocol {
            cursor: pointer;
            color: #63b5fb;
        }
    }
}


.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.fade-enter-to,
.fade-leave-from {
    opacity: 1;
}

.fade-enter-active,
.fade-leave-active {
    transition: all .5s ease;
}
</style>
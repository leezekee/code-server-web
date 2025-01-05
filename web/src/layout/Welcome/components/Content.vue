<template>
    <div class="flex flex-row w-full">
        <div class="relative flex justify-center items-center flex-1 basis-1/2">
            <img src="@/assets/hero-background.webp" height="90%" width="90%" alt="" class="absolute rounded-xl z-1" />
            <img src="@/assets/example.png" alt="" height="85%" width="85%" class="absolute rounded-xl z-2" />
        </div>
        <div class="flex justify-center items-center flex-1 basis-1/2">
            <div class="flex flex-col">
                <div>
                    <h1 class="text-4xl font-bold">欢迎使用OJPLUS</h1>
                    <p class="text-lg">OJPLUS是一个在线编程平台，提供了丰富的题目和编程环境，帮助你提升编程能力。</p>
                </div>
                <div class="pt-10 text-base text-gray-500">
                    <p>本系统的校内网访问地址：https://buctoj.com/；校外网访问地址：https://buctcoder.com/；</p>
                    <p>在校内时推荐使用校内访问地址；</p>
                    <p>推荐用谷歌浏览器（Chrome）访问本系统，其次为新版Edge或火狐（Firefox）浏览器；</p>
                    <p>本系统内的“复制”、“剪切”和“粘贴”操作是允许的，但是禁止从外部粘贴任何代码，否则有可能能被判定为作弊；</p>
                    <p>如果有卡顿，白屏等问题，点击“退出登录”按钮可以重启云端环境</p>
                    <p>进入系统前需申请使用</p>
                </div>
                <div class="pt-16 flex flex-row space-x-10">
                    <Button label="返回主站" severity="secondary" @click="goback" />
                    <Button label="退出登录" severity="secondary" @click="logout" />
                    <TransitionGroup>
                        <Button label="结束使用" severity="danger" :disabled="!hasApplied" @click="finishUsing" />
                        <Button :label="buttonText" @click="applyForUse" />
                    </TransitionGroup>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';

import Button from 'primevue/button';

const emit = defineEmits(['exit', 'apply', 'enter', 'finish', 'back']);
const props = defineProps({
    hasApplied: Boolean,
});

const hasApplied = computed(() => {
    return props.hasApplied;
});
const buttonText = computed(() => {
    return hasApplied.value ? '进入系统' : '申请使用';
});

const goback = () => {
    emit('back');
};

const logout = () => {
    emit('exit');
};

const finishUsing = () => {
    emit('finish');
};

const applyForUse = () => {
    if (hasApplied.value) {
        emit('enter');
    } else {
        emit('apply');
    }
};


</script>

<style>

</style>
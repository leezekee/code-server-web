<template>
    <div class="pl-4 flex flex-row items-center justify-between">
        <div>
            <i class="pi pi-circle-fill" :style=onlineCountStyle></i>
            <span class="ml-2">当前在线人数：{{ onlineCount }}</span>
        </div>
        <div class="pr-4">
            <span class="pr-4 underline hover:cursor-pointer hover:text-stone-400" @click="checkLog">更新日志</span>
            <span class="underline hover:cursor-pointer hover:text-stone-400" @click="checkFeedback">意见反馈</span>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue';

const props = defineProps({
    onlineCount: Number,
});

const onlineCount = computed(() => {
    return props.onlineCount;
});
const onlineCountStyle = reactive({
    color: 'green',
    fontSize: '.5rem',
});

const emit = defineEmits(['feedback', 'log']);

const checkLog = () => {
    emit('log');
};

const checkFeedback = () => {
    emit('feedback');
};

</script>

<style>

</style>
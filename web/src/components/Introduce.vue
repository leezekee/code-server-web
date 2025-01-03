<template>
    <div class="fixed fs z-10 bg-black/75 top-0 left-0 cp" v-if="visible" @click="handleClick">
        <p class="fixed text-white top-12 left-12">↑点击查看须知</p>
        <p class="fixed text-white top-20 left-12">欢迎使用新在线vscode系统</p>
        <p class="fixed text-white top-24 left-12">首次访问必须查看须知！！！</p>
        <p class="fixed text-white top-28 left-12">如已了解可单击任意处跳过</p>
    </div>
</template>

<script setup>
import useVisitStore from '@/stores/visitStore'
import { storeToRefs } from 'pinia'

const visitStore = useVisitStore();
const { firstVisited: visible } = storeToRefs(visitStore);
console.log('visible', visible);
const handleClick = () => {
    visitStore.setFirstVisited(false);
}
</script>

<style>

</style>
<template>
    <Toolbar class="toolbar">
        <template #start>
            <Button icon="pi pi-info-circle" class="mr-2 hi-f p-0 text-xs" severity="info" label="使用须知" @click="handleInfoClick" text />
            <Button icon="pi pi-refresh" class="hi-f p-0 text-xs" severity="info" label="重置左侧" @click="handleRefreshClick" text />
        </template>

        <template #center>
            <Tag :severity="modeSeverity" class="hi-f w-40 text-xs hover:cursor-pointer" :value="modeText" @click="handleModeClick"></Tag>
        </template>

        <template #end>
            <Button icon="pi pi-pause" class="mr-2 hi-f p-0 text-xs" severity="warn" label="暂时离开" @click="handleLeaveClick" text />
            <Button icon="pi pi-sign-out" class="hi-f text-xs" severity="danger" label="退出使用" @click="handleExitClick" text />
        </template>
    </Toolbar>
</template>

<script setup>
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import useModeStore from '@/stores/modeStore';
import { computed } from 'vue';
import { storeToRefs } from 'pinia';

const modeStore = useModeStore();
const { isFreeMode } = storeToRefs(modeStore);
const emit = defineEmits(['click:info', 'click:refresh', 'click:leave', 'click:exit', 'click:mode']);


const modeText = computed(() => {
    return isFreeMode.value ? '自由模式' : '限制模式';
})
const modeSeverity = computed(() => {
    return isFreeMode.value ? 'success' : 'warn';
});

const handleInfoClick = () => {
    emit('click:info');
}

const handleRefreshClick = () => {
    emit('click:refresh');
}

const handleLeaveClick = () => {
    emit('click:leave');
}

const handleExitClick = () => {
    emit('click:exit');
}

const handleModeClick = () => {
    emit('click:mode');
}
</script>

<style>
.toolbar {
    @apply bg-mygray w-full hh-f p-0 !important;
}

</style>

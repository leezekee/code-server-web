<template>
    <Dialog v-model:visible="visible" header="系统模式" :modal="true" :closable="false">
        <p>暂时离开期间可以切换浏览器或电脑，但是离开超过15分钟将会被视为“退出使用”</p>
        <template #footer>
            <Button label="关闭" @click="close()" :severity="modeSeverity" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';
import useModeStore from '@/stores/modeStore';
import { storeToRefs } from 'pinia';

const modeStore = useModeStore();
const { isFreeMode } = storeToRefs(modeStore);
const modeSeverity = computed(() => isFreeMode.value ? 'success' : 'warn');

const emit = defineEmits(['close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const close = () => {
    emit('close');
};
</script>

<style></style>
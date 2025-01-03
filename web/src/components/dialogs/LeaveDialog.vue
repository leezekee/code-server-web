<template>
    <Dialog v-model:visible="visible" header="暂时离开" :modal="true" :closable="false">
        <p>暂时离开期间可以切换浏览器或电脑，但是离开超过15分钟将会被视为“退出使用”</p>
        <template #footer>
            <Button label="确认" @click="comfirmLeave" severity="warn"/>
            <Button label="取消" @click="close()" severity="warn"/>
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';

const emit = defineEmits(['comfirm:leave', 'close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const comfirmLeave = () => {
    emit('comfirm:leave');
};

const close = () => {
    emit('close');
};
</script>

<style>

</style>
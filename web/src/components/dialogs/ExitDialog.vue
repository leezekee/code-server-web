<template>
    <Dialog v-model:visible="visible" header="退出使用" :modal="true" :closable="false">
        <p>确认退出吗？退出前切记保存数据，该操作无法撤销，除“answer”文件夹外的所有数据不会保存</p>
        <template #footer>
            <Button label="确认" @click="comfirmExit" severity="danger" />
            <Button label="取消" @click="close()" severity="danger" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';

const emit = defineEmits(['comfirm:exit', 'close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const comfirmExit = () => {
    emit('comfirm:exit');
};

const close = () => {
    emit('close');
};
</script>

<style></style>
<template>
    <Dialog v-model:visible="visible" header="重置左侧页面" :modal="true" :closable="false">
        <p>左侧网页可能会因为异常原因，跳转至意料之外的界面。使用该功能可以重置左侧网页，是否要重置呢？</p>
        <template #footer>
            <Button label="确认" @click="comfirmRefresh" severity="info" />
            <Button label="取消" @click="close()" severity="info" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';

const emit = defineEmits(['comfirm:refresh', 'close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const comfirmRefresh = () => {
    emit('comfirm:refresh');
};

const close = () => {
    emit('close');
};
</script>

<style></style>
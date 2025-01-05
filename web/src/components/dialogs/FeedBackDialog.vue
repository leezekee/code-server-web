<template>
    <Dialog v-model:visible="visible" header="反馈" :modal="true" :closable="false">
        <p>反馈方式1：加入QQ群（933410835）</p>
        <p>反馈方式2：<a href="https://wj.qq.com/s2/8625740/684f/" target="_blank">https://wj.qq.com/s2/8625740/684f/</a></p>
        <template #footer>
            <Button label="确认" @click="comfirm" severity="secondary" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';

const emit = defineEmits(['comfirm:close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const comfirm = () => {
    emit('comfirm:close');
};
</script>

<style></style>
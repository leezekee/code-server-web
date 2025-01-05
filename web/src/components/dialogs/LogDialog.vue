<template>
    <Dialog v-model:visible="visible" header="更新日志" :modal="true" :closable="false">
        <!-- TODO 获取更新日志 -->
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
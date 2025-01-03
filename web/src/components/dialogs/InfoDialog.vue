<template>
    <Dialog v-model:visible="visible" header="使用须知" :modal="true" :closable="false">
        <p>左侧网页可能会因为异常原因，跳转至意料之外的界面。使用该功能可以重置左侧网页，是否要重置呢？</p>
        <template #footer>
            <Button label="关闭" @click="close" severity="info" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import { computed } from 'vue';
import useVisitStore from '@/stores/visitStore';
import { storeToRefs } from 'pinia';

const visitStore = useVisitStore();
const { firstVisited } = storeToRefs(visitStore);

const emit = defineEmits(['close'])

const props = defineProps({
    visible: {
        type: Boolean,
        required: true
    }
});

const visible = computed(() => props.visible);

const close = () => {
    if (firstVisited.value) {
        visitStore.setFirstVisited(false);
    }
    emit('close');
};
</script>

<style></style>
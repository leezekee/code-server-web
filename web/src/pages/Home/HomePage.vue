<template>
    <HomeLayout 
        @click:info="handleInfoClick"
        @click:refresh="handleRefreshClick"
        @click:mode="handleModeClick"
        @click:leave="handleLeaveClick"
        @click:exit="handleExitClick"
    >
        <template #left>
            <OJPage ref="OJ"/>
        </template>
        <template #right>
            <VscodePage ref="VSC"/>
        </template>
    </HomeLayout>

    <Teleport to="body">
        <!-- ===================以下均为对话框======================== -->
        <LeaveDialog v-model:visible="leaveDialogVisible" @comfirm:leave="hanleLeaveComfirm" @close="handleDialogClose"/>
        <ExitDialog v-model:visible="exitDialogVisible" @comfirm:exit="hanleExitComfirm" @close="handleDialogClose"/>
        <RefreshDialog v-model:visible="refreshDialogVisible" @comfirm:refresh="hanleRefreshComfirm" @close="handleDialogClose"/>
        <InfoDialog v-model:visible="infoDialogVisible" @close="handleDialogClose"/>
        <ModeDialog v-model:visible="modeDialogVisible" @close="handleDialogClose"/>
        <!-- ======================================================== -->
        
        <!-- =======================引导界面========================== -->
        <Introduce />
        <!-- ======================================================== -->
    </Teleport>
</template>

<script setup>
import HomeLayout from '@/layout/Home/HomeLayout.vue';
import LeaveDialog from '@/components/dialogs/LeaveDialog.vue';
import InfoDialog from '@/components/dialogs/InfoDialog.vue';
import RefreshDialog from '@/components/dialogs/RefreshDialog.vue';
import ModeDialog from '@/components/dialogs/ModeDialog.vue';
import ExitDialog from '@/components/dialogs/ExitDialog.vue';
import OJPage from '@/pages/Home/components/OJPage.vue';
import VscodePage from '@/pages/Home/components/VscodePage.vue';
import Introduce from '@/components/Introduce.vue'
import { ref, onMounted } from 'vue';
import { ListenType, newListener } from '@/utils/pasteboardListener';

const leaveDialogVisible = ref(false);
const infoDialogVisible = ref(false);
const refreshDialogVisible = ref(false);
const modeDialogVisible = ref(false);
const exitDialogVisible = ref(false);

const OJ = ref(null);
const VSC = ref(null);

// ==================监听剪贴板======================
/** example:
const listener = newListener();


const handlePaste = (data) => {
    console.log('Paste:', data);
};

onMounted(() => {
    listener.mount(window, ListenType.PASTE, handlePaste);
    listener.start();
});
*/
// ================================================

// ==================打开对话框======================
const handleInfoClick = () => {
    infoDialogVisible.value = true;
};

const handleRefreshClick = () => {
    refreshDialogVisible.value = true;
};

const handleModeClick = () => {
    modeDialogVisible.value = true;
};

const handleLeaveClick = () => {
    leaveDialogVisible.value = true;
};

const handleExitClick = () => {
    exitDialogVisible.value = true;
};
// ================================================

// ==================对话框确认======================
const hanleLeaveComfirm = () => {
    console.log('Leave comfirmed');
    leaveDialogVisible.value = false;
};

const hanleExitComfirm = () => {
    console.log('Exit comfirmed');
    exitDialogVisible.value = false;
};

const hanleRefreshComfirm = () => {
    console.log('Refresh comfirmed');
    refreshDialogVisible.value = false;
};
// ================================================

// ==================对话框确认======================
// 对话框为模态的，只会有一个对话框打开
// 所以关闭对话框时，只需要将所有对话框的visible设置为false
const handleDialogClose = () => {
    leaveDialogVisible.value = false;
    infoDialogVisible.value = false;
    refreshDialogVisible.value = false;
    modeDialogVisible.value = false;
    exitDialogVisible.value = false;
};


</script>

<style>

</style>
<!-- 欢迎界面 -->
<template>
    <WelcomeLayout>
        <template #header>
            OJPLUS —— BUCTOJ在线编程平台
        </template>

        <template #content>
            <Content 
                @exit="logout" 
                @apply="applyForUse" 
                @enter="enterSystem" 
                @finish="finishUsing" 
                @back="goback" 
                :hasApplied="hasApplied"
            />
        </template>

        <template #footer>
            <Footer @feedback="openFeedback" @log="openLog" :onlineCount="onlineCount"/>
        </template>
    </WelcomeLayout>

    <Teleport to="body">
        <LogDialog :visible="logDialogVisible" @comfirm:close="dialogClose" />
        <FeedBackDialog :visible="feedBackDialogVisible" @comfirm:close="dialogClose" />
    </Teleport>
</template>

<script setup>
import WelcomeLayout from '@/layout/Welcome/WelcomeLayout.vue';
import LogDialog from '@/components/dialogs/LogDialog.vue'
import FeedBackDialog from '@/components/dialogs/FeedBackDialog.vue'
import Content from '@/layout/Welcome/components/Content.vue';
import Footer from '@/layout/Welcome/components/Footer.vue';


import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const hasApplied = ref(false);
const onlineCount = ref(0);
const logDialogVisible = ref(false);
const feedBackDialogVisible = ref(false);

onMounted(() => {
    // 获取申请访问状态
    getAppliedInfo();
    // 获取在线人数
    setInterval(getOnlineCount, 3000);
})

// -------------初始化函数---------------

const getAppliedInfo = () => {
    
}

const getOnlineCount = () => {
    onlineCount.value++
}

// ------------------------------------

const goback = () => {
    router.go(-1);
};

const logout = () => {
    // router.push('/login');
};

const finishUsing = () => {
    hasApplied.value = false;
};

const applyForUse = () => {
    hasApplied.value = true;
};

const enterSystem = () => {
    router.push('/code');
};

// ---------------对话框相关----------------

const openFeedback = () => {    
    feedBackDialogVisible.value = true;
};

const openLog = () => {
    logDialogVisible.value = true;
};

const dialogClose = () => {
    logDialogVisible.value = false;
    feedBackDialogVisible.value = false;
};
</script>

<style>

</style>
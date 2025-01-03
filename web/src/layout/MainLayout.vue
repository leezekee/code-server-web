<template>
    <Space direction="vertical" :style="{ width: '100%', height: '100vh' }" :size="[0, 48]">
        <Layout>
            <LayoutHeader :style="headerStyle">
                <navbar 
                    @click:info="handleInfoClick"
                    @click:refresh="handleRefreshClick"
                    @click:mode="handleModeClick"
                    @click:leave="handleLeaveClick"
                    @click:exit="handleExitClick"
                />
            </LayoutHeader>
            <LayoutContent :style="contentStyle">
                <HorizonContainer>
                    <template #l>
                        <slot name="left"></slot>
                    </template>
                    <template #r>
                        <slot name="right"></slot>
                    </template>
                </HorizonContainer>
            </LayoutContent>
        </Layout>
    </Space>
</template>

<script setup>
import HorizonContainer from './components/HorizonContainer.vue';
import Navbar from './components/Navbar.vue';

import { Layout, LayoutHeader, LayoutContent, Space } from 'ant-design-vue';

const emit = defineEmits(['click:info', 'click:refresh', 'click:mode', 'click:leave', 'click:exit']);

const handleInfoClick = () => {
    emit('click:info');
};

const handleRefreshClick = () => {
    emit('click:refresh');
};

const handleModeClick = () => {
    emit('click:mode');
};

const handleLeaveClick = () => {
    emit('click:leave');
};

const handleExitClick = () => {
    emit('click:exit');
};


const headerStyle = {
    textAlign: 'center',
    color: '#525252',
    height: '48px',
    paddingInline: 50,
    lineHeight: '48px',
    backgroundColor: '#e2e2e2',
};

const contentStyle = {
    minHeight: 'calc(100vh - 48px)',
    color: '#525252',
    backgroundColor: '#e2e2e2',
};
</script>

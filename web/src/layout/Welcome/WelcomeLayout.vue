<!-- 总体布局 -->
<template>
    <Space direction="vertical" :style="{ width: '100%' }" :size="[0, 48]">
        <!-- 上下布局 -->
        <Layout>
            <LayoutHeader :style="headerStyle">
                <slot name="header"></slot>
            </LayoutHeader>
            <LayoutContent :style="contentStyle">
                <slot name="content"></slot>
            </LayoutContent>
            <LayoutFooter :style="footerStyle">
                <slot name="footer"></slot>
            </LayoutFooter>
        </Layout>
    </Space>
</template>

<script setup>
import { Layout, LayoutHeader, LayoutContent, Space, LayoutFooter } from 'ant-design-vue';

const headerStyle = {
    textAlign: 'center',
    color: '#525252',
    height: '48px',
    paddingInline: 50,
    lineHeight: '48px',
    backgroundColor: '#fff',
    fontSize: '24px',
    fontFamily: '-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif',
    fontWeight: 600,
};

const contentStyle = {
    minHeight: 'calc(100vh - 72px)',
    color: '#525252',
    backgroundColor: '#fff',
    display: 'flex',

};

const footerStyle = {
    color: '#525252',
    height: '24px',
    lineHeight: '24px',
    backgroundColor: '#e2e2e2',
    fontSize: '12px',
    padding: 0
};
</script>

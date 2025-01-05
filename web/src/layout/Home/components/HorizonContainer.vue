<template>
    <!-- 横向布局 -->
    <Splitter class="container-wrapper">
        <SplitterPanel class="container"> 
            <slot name="l"></slot> 
        </SplitterPanel>
        <SplitterPanel class="container">
            <slot name="r"></slot>
        </SplitterPanel>
    </Splitter>
</template>

<script setup>
import Splitter from 'primevue/splitter';
import SplitterPanel from 'primevue/splitterpanel';

</script>

<style>
.container-wrapper {
    @apply flex flex-row bg-mygray place-content-between px-3 !important
}

.container {
    @apply rounded-md bg-white ch-f w-1/2 !important
}

.p-splitter-gutter {
    @apply relative flex rounded-lg justify-center items-center bg-mygray hover:bg-myblue hover:cursor-ew-resize mx-1 ch-f w-1 !important
}

.p-splitter-gutter {
    @apply after:content-[''] after:block after:w-1 after:h-10 after:bg-gray-500 after:absolute after:left-1/2 after:top-1/2 after:-translate-x-1/2 after:-translate-y-1/2 after:rounded-lg hover:after:bg-none
}
</style>
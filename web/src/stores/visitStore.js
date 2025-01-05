import { defineStore } from "pinia";
import { ref } from "vue";

/**
 * 记录用户是否第一次访问，用于控制引导界面的显示
 * 
 * TODO: 未来可以考虑将用户访问记录存储到数据库中，以便在用户更换设备时保持用户访问记录
 */
const useVisitStore = defineStore("visit", () => {
    const firstVisited = ref(true);
    const setFirstVisited = (value) => {
        firstVisited.value = value;
    }
    const getFirstVisited = () => {
        return firstVisited.value;
    }
    return {
        firstVisited,
        setFirstVisited,
        getFirstVisited
    }
}, {
    persist: true,
});

export default useVisitStore;
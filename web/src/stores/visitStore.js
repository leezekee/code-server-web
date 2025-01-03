import { defineStore } from "pinia";
import { ref } from "vue";

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
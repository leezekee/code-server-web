import { defineStore } from "pinia";
import { ref, computed } from "vue";
import config from "@/config";

const useModeStore = defineStore("mode", () => {
    const mode = ref(config.SYS_MODE.FREE);

    const SetMode = (newMode) => {
        if (newMode !== config.SYS_MODE.FREE && newMode !== config.SYS_MODE.LIMITED) {
            console.error("Invalid mode: ", newMode);
            return;
        }
        mode.value = newMode;
    };
    const getMode = () => computed(() => mode.value);
    const toggleMode = () => {
        mode.value = mode.value === config.SYS_MODE.FREE ? config.SYS_MODE.LIMITED : config.SYS_MODE.FREE;
    };
    const isFreeMode = computed(() => mode.value === config.SYS_MODE.FREE);
    const isLimitedMode = computed(() => mode.value === config.SYS_MODE.LIMITED);
    return {
        mode,
        getMode,
        SetMode,
        toggleMode,
        isFreeMode,
        isLimitedMode
    };
});

export default useModeStore;
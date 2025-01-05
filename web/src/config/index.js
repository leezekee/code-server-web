/**
 * @fileoverview 全局配置文件
 * @module config
 * @version 1.0.0
 * 
 * 定义全局配置常量，与后端配置保持一致
 * SYS_MODE: 系统运行模式
 * CODE: 接口返回码
 * 
 * TODO: 将来可以考虑将这些配置项放到后端配置中，前端通过接口获取
 */
export default {
    SYS_MODE: {
        LIMITED: 100000,
        FREE: 100001,
    }
}
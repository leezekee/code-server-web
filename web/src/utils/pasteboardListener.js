const tokens = ["include", "for", "int", "double", "namespace", "float", "memset", "stdio.h",
    "print", "scanf", "cin", "cout", "struct", "scanf", "void", "if", "break",
    "string", "list", "map", "String", "Object",
    "return", "continue", "def", "__name__", "import"];
const warnText = "提示：VSCode已经将你剪切板的代码自动删除，请手动编写代码！";
const ListenType = {
    COPY: 'copy', 
    CUT: 'cut',
    PASTE: 'paste',
}
/**
 * 新建一个监听器，监听剪贴板内容，监听器可以挂载多个元素
 * 监听器mount方法接受两个参数，一个是元素，一个是处理函数
 * 监听器start方法开始监听
 * @returns {any} 返回一个新的监听器，调用start方法开始监听
 */
function newListener() {
    const queue = [];
    const listener = {
        mount: (ele, type, han = null) => {
            if (type === ListenType.PASTE && han === null) {
                throw new Error('当监听类型为剪切板读取时，处理函数不能为空');
            }
            const handler = han || defaultReplaceWhenCopyOrCut;
            const listenObj = {
                element: ele,
                handler: handler,
                type: type
            }
            queue.push(listenObj);
            console.log('监听器挂载成功', listenObj);
        },
        start: () => {
            queue.forEach((obj) => {
                obj.element.addEventListener(obj.type, obj.handler);
            });
        }
    }
    return listener;
}

/**
 * 默认的剪切板监听处理函数，当剪切板内容包含tokens中的关键字时，将剪切板内容替换为warnText
 * @param {*} event 
 */
function defaultReplaceWhenCopyOrCut(event) {
    const clipboardData = event.clipboardData || window.clipboardData;
    const text = clipboardData.getData('text/plain');
    const textArray = text.split(/\s+/);
    for (let i = 0; i < textArray.length; i++) {
        if (tokens.includes(textArray[i])) {
            event.preventDefault();
            event.clipboardData.setData('text/plain', warnText);
            break;
        }
    }
}

export {
    ListenType,
    newListener
}
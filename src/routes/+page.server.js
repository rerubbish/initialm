
const jsonImports = import.meta.glob('/mao/assets/rules/*.json', { eager: true });
export async function load() {
    const rulesJsonData = Object.entries(jsonImports).map(([path, module]) => {
        // 从路径中提取文件名
        const filename = path.split('/').pop();
        return {
            filename,
            data: module.default
        };
    });
    return { rulesJsonData };
}
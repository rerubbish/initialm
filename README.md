# initialm

结合了WebAssembly 和 Svelte 的一个可以简单使用的项目脚手架模板,只需要编写js与golang就可以在不使用后端服务的情况下可以做到相近功能。这个项目由[gendk](https://github.com/langbiantianya/gendk)演化,回归js前端换来更好的定制性。

## todo

- [ ] 基础界面
- [ ] 基础渲染模板
- [ ] 示例项目

## 如何开发

### 环境

- golang 1.25.3+
- nodejs 22 | 23
- make

### 命令

#### 安装依赖

```shell
npm install && cd mao && go mod tidy && cd -
```

#### 预览

```shell
npm run dev
```

#### 编译

```shell
npm run build
```

## 部署

这部分参考[svelte adapters](https://svelte.dev/docs/kit/adapters)

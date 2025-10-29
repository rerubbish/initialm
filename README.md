# initialm

这个项目由[gendk](https://github.com/langbiantianya/gendk)演化,回归js前端换来更好的定制性。结合了WebAssembly 和 Svelte 的一个可以简单使用的项目脚手架模板,只需要编写json规则和一点点的js与golang就可以在不使用后端服务的情况下可以做到相近功能。

## todo

- [x] 初版规则约定
- [ ] 基础前端渲染
- [x] 基础模板渲染
- [ ] 示例项目
- [ ] 使用浏览器运行js，在规则配置文件中处理复杂字段。

## 如何使用

目前只需要添加配置文件与模板即可，配置文件参考[Rules.md](Rules.md)。

## 如何开发

### 环境

- Golang 1.25.3+
- Nodejs 22 | 23
- GNU Make

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

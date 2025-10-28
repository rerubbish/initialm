# 约定

## tips

- golang打包go模板时，里面存在go.mod会忽略整个文件夹，需要改名。

## 目录

### 规则目录

### 模板目录

## 规则

### 文件名与模板名称

### 规则文件

#### 文件名\路径名替换

#### 变量替换

## 前端

### 组件

#### 校验

##### 自定义校验

```json
{
     "type": "text",
     "key": "Hello",
     "validate": {
         "trigger": ["change", "blur"],
         "required": true,
         "message":"必填",
         "ValidatorExpr":"typeof str !== 'string' && str.trim() === '' && isNaN(Number(value))"
     }
}
```

```js
// 表达式直接返回校验结果（布尔值）
const rule = {
  customValidatorExpr: `typeof str !== 'string' && str.trim() === '' && isNaN(Number(value))`
};
// validator 是一个直接返回校验结果的函数
const validator = new Function('value', 'formValues', `return ${rule.customValidatorExpr};`);

// 直接调用 validator 即可得到结果
console.log(validator('123', { password: '123' })); // true
console.log(validator('456', { password: '123' })); // false

```

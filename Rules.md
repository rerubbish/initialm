# 约定

## tips

- golang打包go模板时，里面存在go.mod会忽略整个文件夹，需要改名。

## 目录

### 规则目录

[rules](mao/assets/rules)目录中存放了所有的自定义规则文件，每个文件都是一个json文件，文件名称就是模板的文件夹名称。

### 模板目录

[templates](mao/assets/template)目录中存放了所有的模板项目，每个文件夹就是一个模板项目，文件夹名称就是模板的名称。

## 规则

### 文件名与模板名称

mao/assets/rules/{模板名称}.json 规则文件
mao/assets/template/{模板名称} 模板目录

### 规则文件

#### 文件名\路径名替换

##### 例子

这个例子会将模板项目中的go.modd关键字替换为go.mod关键字，实际就是修改了最后的文件名称。支持关键字也支持完整路径。完整路径由对应模板项目的路径开始例如`go.modd`,就是对应模板项目的根目录下的`go.modd`文件。

```json
    "path": [
        {
            "source": "go.modd",
            "target": "go.mod"
        }
    ]
```

###### source

模板项目的完整路径、文件名称、关键字

###### target

输出的时候会将source中的关键字替换为target中的关键字

#### 文件中变量替换

这个参考golang的[template](https://pkg.go.dev/text/template)标准库

##### 例子

- filePath 模板项目中的文件路径，例如`main.go`
- substitution
  - key 变量的关键字，例如`Hello`
  - default 默认值，例如`默认值`
  - required 是否必填，例如`true`

```json
"variable": [
        {
            "filePath": "main.go",
            "substitution": [
                {
                    "key": "Hello",
                    "default": "默认值",
                    "required": true
                }
            ]
        }
    ]
```

### 前端

#### 组件

目前前端组件都是输入相关的。  

- text
- file
- select
- radio
- checkbox

#### 规则

##### 例子

```json
"components": [
        {
            "label": "Hello?",
            "tips": "你好占位符",
            "type": "text",
            "key": "Hello",
            "validate": {
                "required": true,
                "ValidatorExpr": {
                    "type": "pattern",
                    "rule": "^[^0-9]*",
                    "message": "Hello 不可为数字"
                }
            }
        }
    ]
```

##### 生成逻辑

生成顺序与components数组顺序一致

- key 是向后端传递的参数名
- label 是前端显示的名称
- tips 是前端显示的提示信息
- type 是前端组件的类型

##### 校验

###### 简单校验

- required 是必填项校验，例如`true`

###### 自定义校验

目前只支持正则校验，也就是 `type` = "pattern"，message 是校验失败时显示的提示信息。

```json
"validate": {
                "required": true,
                "ValidatorExpr": {
                    "type": "pattern",
                    "rule": "^[^0-9]*",
                    "message": "Hello 不可为数字"
                }
            }
```

## 记录

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

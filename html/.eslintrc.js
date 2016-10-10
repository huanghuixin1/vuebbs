module.exports = {
    "env": {
        "browser": true,
        "commonjs": true,
        "es6": true,
        "node":true
    },
    "globals": {
        "window.console": true,
        "process": true,
        "console":true
    },
    "ignorePath ":".eslintignore",
    // "extends": "eslint:recommended",
    "parserOptions": {
        "sourceType": "module"
    },
    "rules": {
        "indent": [//必须用4个空格缩进
            2,
            4
        ],
        "quotes": [ //字符串必须用双引号
            "error",
            "double"
        ],
        "semi": [ // 必须有分号
            2,
            "always"
        ],
        "space-infix-ops":[ //必须有空格
            2
        ],
        "no-use-before-define":[
            1
        ]
    }
};
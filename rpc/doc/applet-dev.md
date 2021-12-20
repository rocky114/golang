# applet

### 前置要求
- applet-dev-studio [文档](http://viper.pages.gitlab.sz.sensetime.com/luacvruntime/manual/01-Getting-Started.md.html)
- Ubuntu >= 16.04 or CentOS >= 7.4
- lua [5.1.5]
- LuaRocks [2.2.2] // lua包管理工具
- docker [版本>=18]
- NVidia GPU Driver (or drivers for other accelerators) and nvidia-docker
---

### luarocks配置
- 设置luarocks私有仓库配置
```shell
   $ cd ~/.luarocks && vim config-5.1.lua 
```
```text 
   rocks_servers = {
       "http://luarocks.opencloud.sensetime.com"
   }
```

### 系统环境变量配置
```sh
export ADELA_MODEL_USER="kestrel-robot"
export ADELA_MODEL_PASS="602bcdab78074fe0bbffd33262b6d917"

export LUAROCKS_SERVER="http://luarocks.opencloud.sensetime.com"
export KESTREL_PRIVATE_TOKEN=vsHrcChsBY56XEfC5EWz

export DEVCENTER_MODEL_USER="model_download_robot"
export DEVCENTER_MODEL_PASS="downloader"
```

### 下载最新的证书

```shell
   $ cd ~/.applet-dev-studio && ./script/license.sh
```

### 创建应用
```shell
   $ cd ~
   $ applet-dev-studio init algo-my-app1 //初始化
   $ applet-dev-studio package /home/aie/algo-my-app1 . //打包[必需绝对路径]
   $ applet-dev-studio run com.example.app1.v10000.1c8180a.algo # replace file name with yours
```

---
### 常见问题
1. fatal: Needed a single revision
```text
    applet-dev-studio init algo-my-app1 初始化需要修改默认选项
```

2. 执行命令 applet-dev-studio package /home/aie/algo-my-app1 . 报geometry2d包缺失
```text
    不需要处理，实际luarocks会自动下载依赖; 手动执行make app会看到依赖包下载
```

3. license is expired
```shell
    $  cd ~/.applet-dev-studio && ./script/license.sh
```

4. lua依赖包找不到
```text
   配置lua私有仓库，见配置luarocks 
```


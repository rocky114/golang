Pallas Packing
======

# GIT

## git-lfs

### 1. 下载安装 LFS:

```sh
# 下载
$ wget -O git-lfs-linux-amd64-v2.8.0.tar.gz http://file.kestrel.sensetime.com/kestrel_env/git-lfs-linux-amd64-v2.8.0.tar.gz

# 解压
$ tar zxvf git-lfs-linux-amd64-v2.8.0.tar.gz

# 安装
$ sudo ./install.sh
```

### 2. 配置git-lfs

```sh
$ git config --global lfs.url "http://devcenter.bj.sensetime.com/api/v1/lfs/"

$ git config --global credential.helper store
```

> 在第一次 clone 带有 LFS 文件的项目时，会询问账号与秘钥. 此处账号名为公司 LDAP 的用户名，秘钥为 devcenter 的 secret，登陆 [DevCenter](http://devcenter.sensetime.com/) 后，点击左上角 Menu->Profile->ApiSecret->Show 可显示. 如果没有生成过，点击按钮生成.

> 在第一次登录之后，相关信息会存在 <HOME>/.git-credentials 文件内，如果需要更换用户或者更新 secret，可以修改此文件或者删除相应内容后重复上述操作.

### 3. confluence 参考文档:

* [git lfs 配置](https://confluence.sensetime.com/pages/viewpage.action?pageId=234364417)

* [LFS 大文件](https://confluence.sensetime.com/pages/viewpage.action?pageId=157713720)


## git submodule

### 1. 添加子模块

以product-demo项目为例:

> product-demo.url = git@gitlab.sz.sensetime.com:aie-eeg/product-demo.git

```sh
# 添加aie-stpu-lt-install子模块，路径变更为 ansible/roles/aie-stpu-lt
$ git submodule add git@gitlab.sz.sensetime.com:aie-eeg/aie-stpu-lt-install.git ansible/roles/aie-stpu-lt


# 查看 .gitmodules 文件
$ cat .gitmodules
-------------------------------------------------------------------
[submodule "ansible/roles/aie-stpu-lt"]
        path = ansible/roles/aie-stpu-lt
        url = git@gitlab.sz.sensetime.com:aie-eeg/aie-stpu-lt-install.git



$ git submodule
------------------------------------------------------------------
33037faebb8eca9bf8f557a6796f21a7b3cabf2e ansible/roles/aie-stpu-lt (remotes/origin/HEAD)
```

### 2. 修改子模块url

> 由于 product-demo 与 aie-stpu-lt-install 在同一组下, URL使用相对路径更加简洁 (infra要求必须使用相对路径)

```sh
$ vi .gitmodules
-------------------------------------------------------------------
[submodule "ansible/roles/aie-stpu-lt"]
        path = ansible/roles/aie-stpu-lt
        url = ../aie-stpu-lt-install.git
```

### 3. 拉取完整子模块内容

```sh
# 拉取完整的子模块代码
git submodule update --init --recursive
```

### 4. 参考文档

* [Git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules)


## git merge

> 由于可能需要合并 viper 的 product-demo，需要进行跨项目的分支合并


# Conan

## 源码项目中集成conan

### 1. 在源码目录下添加 `conanfile.py` , 以 aie-stpu-lt为例:

```sh
$ tree -L 1
------------------------------------------
.
├── api
├── cmd
├── conanfile.py
├── config
├── dlv
├── doc
├── go.mod
├── go.sum
├── interface_json
├── Makefile
├── makePkg.py
├── merge_swagger.rb
├── monit
├── pb
├── pkg
├── README.md
├── regen_docker.sh
├── regen_new.sh
├── regen.sh
├── regen_swagger_docker.sh
├── regen_swagger.sh
├── scripts
├── vendor
└── version


### 2. 根据实际情况, 修改conanfile.py内容

$ vi conanfile.py
-----------------------------------------------------
import os
from conans import ConanFile, tools

class AieStpuLtConan(ConanFile):
    name = "aie-stpu-lt"  # 每个代码仓库名
    version = ""
    settings = "arch_target"
    options = {"device": ["pallas"]}
    default_options = {"device": "pallas"}
    description = "Package for Pallas"
    url = "https://gitlab.sz.sensetime.com/aie-eeg/aie-stpu-lt"
    license = "SenseTime Corporation"
    author = "IDEA-ET"
    generators = "make"

    def package(self):
        self.copy("*", dst="build", src="build")

    def set_version(self):
        git = tools.Git()
        branch = os.getenv('CI_COMMIT_REF_NAME')
        if branch is None:
            branch = git.get_branch()
        self.version = "%s-%s" % (branch, git.get_revision()[:7])
```

### 3. 修改Makefile

```sh
BUILD := `git rev-parse --short HEAD`
TARGETS := aie-stpu-lt
SUPPORTEDTAGLIST = pallas

CHANNEL := $(or $(CI_COMMIT_REF_NAME), $(shell git rev-parse --abbrev-ref HEAD))

conan:
	@make build
	conan user  -p uayaeCh8 -r viper viper-robot
	mkdir -p build/bin
	cp -rf scripts/shell ./build/
	cp -rf scripts/sql ./build/
	cp config/aie-stpu-lt.conf ./build/
	cp config/config.json ./build/
	cp config/server.crt ./build/
	cp config/server.crt.bk ./build/
	cp config/server.key ./build/
	cp config/server.key.bk ./build/
	@for device in $(SUPPORTEDTAGLIST); do \
	cp $(TARGETS)_arm64_$${device} ./build/bin/; \
	conan export-pkg -f conanfile.py viper/$(CHANNEL) -s arch_target=aarch64 -o device=$${device} || exit 1; \
	rm -f ./build/bin/*; \
	done
	conan upload $(TARGETS)/$(CHANNEL)-$(BUILD)@viper/$(CHANNEL) -r viper --all --confirm
```

遵循以下规范：

1. 可执行文件放在./build/bin下
2. 数据库相关文件或数据迁移相关的文件放在./build/database下
3. 其他种类文件，可以放在./build/<kind>对应的目录下

> 这里我遵循的不是很好，需要调整~ 不过最重要的是和之后的install项目协同好.

之后就可以通过执行 `make conan` 命令 进行包上传


### 4. 添加并修改 `.gitlab-ci.yml`

<!-- TODO -->...


### 5. conan的部分使用命令

```sh
$ conan search aie-stpu-lt/master-3d7d7a8@viper/master
-----------------------------------------------------------------------------
Existing packages for recipe aie-stpu-lt/master-3d7d7a8@viper/master:

    Package_ID: c49a1362ea68cf68fb840e9afbb55c6c83df4403
        [options]
            device: pallas
        [settings]
            arch_target: aarch64
        Outdated from recipe: False


$ conan info aie-stpu-lt/master-3d7d7a8@viper/master -s arch_target=aarch64 -o device=pallas
aie-stpu-lt/master-3d7d7a8@viper/master
    ID: c49a1362ea68cf68fb840e9afbb55c6c83df4403
    BuildID: None
    Remote: viper=http://conan.kestrel.sensetime.com/artifactory/api/conan/viper
    URL: https://gitlab.sz.sensetime.com/aie-eeg/aie-stpu-lt
    License: SenseTime Corporation
    Author: IDEA-ET
    Description: Package for Pallas
    Provides: aie-stpu-lt
    Recipe: Cache
    Binary: Cache
    Binary remote: viper
    Creation date: 2021-10-27 20:02:08 UTC
```

### 6. confluence 参考文档:

* [conan 集成](https://confluence.sensetime.com/pages/viewpage.action?pageId=230563718)


# aie-stpu-lt[ansible] 产品打包

aie-stpu-lt的pallas打包，使用到了
`aie-stpu-lt`,
`aie-stpu-lt-install`,
`srs-install`,
`license`,
`product-demo`
这些项目，它们大都可以在 [这个项目组下](https://gitlab.sz.sensetime.com/aie-eeg) 找到.

> license 项目，目前直接作为role文件，放在了 product-demo 中

下面分别介绍它们~

## aie-stpu-lt

这个就是源码项目，按照上文的内容，已经集成了conan~

## aie-stpu-lt-install

这个项目本来的目的，是用来下载 aie-stpu-lt 项目中上传到conan的包, 并接入infra~ 但是初期打包为了简单，我把`web包`也放在了这里

```sh
$ tree -L 2
-----------------------------------
.
├── ansible
│   ├── convention.yml
│   └── playbook.yml
├── build
├── common
│   ├── build
│   ├── defaults
│   ├── files
│   ├── molecule
│   ├── tasks
│   ├── templates
│   └── vars
├── defaults
│   └── main.yml
├── files
│   └── web.zip
├── Makefile
├── README.md
├── tasks
│   ├── config.yml
│   ├── install.yml
│   ├── main.yml
│   └── variables.yml
├── templates
│   └── config.yml.j2
└── vars
    └── main.yml
```

## product-demo

### 1. license

```sh
$ ree -L 3
--------------------------------------
.
├── files
│   └── license
│       ├── active_code.code
│       ├── client.lic
│       ├── client.pem
│       ├── cluster.lic
│       └── J202110140003.zip
├── tasks
│   └── main.yml
└── templates
    └── license-init.sh.j2
```

### 2. srs

```sh
$ tree -L 2
---------------------------
.
├── ansible
│   ├── convention.yml
│   └── playbook.yml
├── defaults
│   └── main.yml
├── files
│   ├── srs
│   └── srs.conf
├── tasks
│   ├── config.yml
│   ├── install.yml
│   ├── main.yml
│   └── variables.yml
...
```

# infra && container

## infra

### infra相关权限

1. Harbor 权限

   [Harbor](https://registry.sensetime.com/harbor/projects)

2. 需要向viper申请权限的仓库列表

    * https://gitlab.sz.sensetime.com/viper/engine-image-ingress-service
    * https://gitlab.sz.sensetime.com/viper/engine-static-feature-db
    * https://gitlab.sz.sensetime.com/viper/engine-model-packages
    * https://gitlab.sz.sensetime.com/viper/gosdkwrapper

3. DevCenter 权限

   [DevCenter](http://devcenter.sensetime.com/) 访问：http://devcenter.sensetime.com/

   之后需要DevCenter中的secret密钥.

4. confluence 参考文档[当然也需要confluence访问权限]:

   [confluence文档-第五部分](https://confluence.sensetime.com/pages/viewpage.action?pageId=247907005)


## container

### 1. 镜像拉取

1. 登录私有镜像仓库(registry.sensetime.com) Harbor权限

```sh
$ docker login registry.sensetime.com
```

2. 拉取 viperlite/infra 镜像，注意标签，可以上[Harbor]查看.

> 镜像的TAG可以查看 Makefile文件里的 VERSION 值

此处以 TAG:v1.2.2 为例

```sh
$ docker pull registry.sensetime.com/viperlite/infra:v1.2.2
```

拉取完成后 可以在本地镜像列表中查找, 之后可以直接使用

```sh
$ docker image ls
----------------------------------------------------------------------------------------------
REPOSITORY                                      TAG      IMAGE ID       CREATED         SIZE
multiarch/qemu-user-static                      latest   113c42ae5ac7   5 weeks ago     299MB
registry.sensetime.com/viperlite/infra          v1.2.2   4413877eb0e5   2 months ago    1.5GB
aiop/ubuntu                                     latest   2c9296a77f38   2 months ago    1.3GB
alpine                                          latest   14119a10abf4   2 months ago    5.6MB
...
```

### 2. 启动容器

> 启动容器的时候，最好不要在项目目录下. [见常见打包问题]()

```sh
# 使用 -it 参数  目录会在容器启动后，直接进入容器内；  
# 使用 --rm 参数 让容器停止时，自动删除
# -v 参数 将本地目录映射到容器内
$ docker run --rm --privileged=true -itd -v /mnt/sensetime01/xcell/repos/product-demo:/mnt/data/product_demo -e CONAN_VIPER_ROBOT_PASSWORD=uayaeCh8 registry.sensetime.com/viperlite/infra:v1.2.2 bash
```

### 3. 项目构建

```sh
# 与容器交互
$ docker exec -it [CONTAINER-ID] bash

# 进入项目目录
root@071910bca077:/#  cd /mnt/data/product_demo

# 可以看到映射到这个目录下的本地项目文件
root@071910bca077:/mnt/data/product_demo# ls
---------------------------------------------------------------------------------
Makefile  README.md  ansible  build  infra  molecule  service_check.sh

# 清理
root@071910bca077:/mnt/data/product_demo# make clean

# 构建
root@071910bca077:/mnt/data/product_demo# make build-pallas && make release-pallas

# 退出容器
root@071910bca077:/mnt/data/product_demo# exit
```

等待构建完成后，可以去取打包完成的输出文件了

```sh
$ cd /mnt/sensetime01/xcell/repos/product-demo/build/output

$ ls -lh
--------------------------------------------------------------------------------------
-rw-r--r-- 1 root root 1.1G Oct 28 11:12 app_nebula-aie_v1.2.6_ota_s100box.tar.gz
-rw-r--r-- 1 root root 3.8G Oct 28 17:12 nebula-aie-pallas-2.4.9.2-20211028090914.tar.gz
-rw-r--r-- 1 root root 2.8G Oct 28 11:17 nebula-aie_v1.2.6+b1b003d_pallas_image.tar.gz
-rw-r--r-- 1 root root 1.1G Oct 28 11:17 nebula-aie_v1.2.6+b1b003d_pallas_ota_image.tar
-rw-r--r-- 1 root root  212 Oct 28 11:17 ota_info
```


# 其它参考文档

1. [产品线INFRA开发流程](https://confluence.sensetime.com/pages/viewpage.action?pageId=247907005)

2. [gitee LFS](https://gitee.com/help/articles/4235#article-header0)
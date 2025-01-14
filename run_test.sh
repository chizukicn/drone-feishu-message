#!/bin/bash

export DRONE_SYSTEM_HOST=http://127.0.0.1
export DRONE_REPO_NAMESPACE=game
export DRONE_REPO_NAME=game-web-frontent
export DRONE_REPO=game/game-web-frontent
export DRONE_BRANCH=dev
export DRONE_COMMIT_SHA=abcdefghijklmn
export DRONE_COMMIT_AUTHOR_EMAIL=chizukicn@outlook.com
export DRONE_COMMIT_MESSAGE="\n这\n是一条很长很长很长的测试消息，\n大概\n有\n25个字符吧\n"
export DRONE_COMMIT_LINK=http://127.0.0.1/game/game-web-frontent/-/commit/abcdefghijklmn
export DRONE_COMMIT_AUTHOR=keepchen
export DRONE_BUILD_STATUS=failure # failure
export DRONE_BUILD_LINK=http://127.0.0.1/game/game-web-frontent/1
export PLUGIN_TOKEN="0b15b26a-8e0d-4fda-90f5-04637c0c7028" #飞书的webhook token值
export PLUGIN_SECRET=zB1adl7OfkGifmHCqJwKLc # 飞书的签名校验secret
export PLUGIN_CARD_TITLE="Custom Title" # 卡片消息标题
export PLUGIN_SUCCESS_IMG_KEY= # 构建成功图片
export PLUGIN_FAILURE_IMG_KEY= # 构建失败图片
export PLUGIN_POWERED_BY_IMG_KEY= # 版权logo
export PLUGIN_POWERED_BY_IMG_ALT= # 版权logo的alt提示文字

# go build -o drone-feishu
# ./drone-feishu

go run .

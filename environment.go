package main

import (
	"log"
	"os"
)

type PluginEnvironment struct {

	//以下是通过settings设置的环境变量
	//
	//settings:
	//	token:
	//		from_secret: xxx
	//	card_title: xxx
	//	success_img_key: xxx
	//	failure_img_key: xxx
	//	powered_by_img_key: xxx
	//	powered_by_img_alt: xxx
	PluginSecret          string //授权token
	PluginToken           string //授权token
	PluginCardTitle       string //通知卡片标题
	PluginSuccessImgKey   string //构建成功图片imgKey
	PluginFailureImgKey   string //构建失败图片imgKey
	PluginPoweredByImgKey string //powered by图片imgKey
	PluginPoweredByImgAlt string //powered by图片alt信息
}

// GetEnv 获取环境变量
func GetPluginEnv() PluginEnvironment {
	pluginEnv := PluginEnvironment{}
	pluginEnv.PluginToken = os.Getenv("PLUGIN_TOKEN")
	pluginEnv.PluginSecret = os.Getenv("PLUGIN_SECRET")
	pluginEnv.PluginCardTitle = os.Getenv("PLUGIN_CARD_TITLE")
	pluginEnv.PluginSuccessImgKey = os.Getenv("PLUGIN_SUCCESS_IMG_KEY")
	pluginEnv.PluginFailureImgKey = os.Getenv("PLUGIN_FAILURE_IMG_KEY")
	pluginEnv.PluginPoweredByImgKey = os.Getenv("PLUGIN_POWERED_BY_IMG_KEY")
	pluginEnv.PluginPoweredByImgAlt = os.Getenv("PLUGIN_POWERED_BY_IMG_ALT")

	if pluginEnv.PluginToken == "" {
		log.Println("feishu webhook access token can not be empty")
		os.Exit(1)
	}

	if pluginEnv.PluginSecret == "" {
		log.Println("feishu sign secret can not be empty")
		os.Exit(1)
	}

	return pluginEnv
}

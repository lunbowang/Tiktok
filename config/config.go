package config

import "time"

// 存放相关配置

const GO_STARTER_TIME = "2006-01-02 15:04:05"

// 视频模块相关配置
const (
	VIDEO_NUM_PER_REFRESH     = 6
	VIDEO_INIT_NUM_PER_AUTHOR = 10
	// 阿里 OSS 相关配置
	OSS_ACCESS_KEY_ID     = ""
	OSS_ACCESS_KEY_SECRET = ""
	OSS_BUCKET_NAME       = "durlim-tiktok"
	OSS_ENDPOINT          = "oss-cn-beijing.aliyuncs.com"
	CUSTOM_DOMAIN         = "durlim-tiktok.oss-cn-beijing.aliyuncs.com"
	OSS_VIDEO_DIR         = "tiktok"
	PLAY_URL_PREFIX       = CUSTOM_DOMAIN + OSS_VIDEO_DIR
	COVER_URL_SUFFIX      = "?x-oss-process=video/snapshot,t_2000,m_fast"
	OSS_USER_AVATAR_DIR   = "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3"
)

const LIKE = 1

const BG_IMAGE = "https://tse3-mm.cn.bing.net/th/id/OIP-C.449_jFJDOg6VBfOwJU_hrQHaEK?w=309&h=180&c=7&r=0&o=7&dpr=1.1&pid=1.7&rm=3"

const SIGNATURE = "我想休息"

const SECRETE = "Durlim"

const (
	ExpireTime = 86400 * time.Second
)

var LatestRequestTime = make(map[string]time.Time, 100)

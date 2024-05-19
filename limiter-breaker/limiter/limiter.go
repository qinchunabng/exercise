package limiter

import (
	"sync"
	"time"
)

//限流器
//1.对某个接口单位时间内的访问量的限制作用
//2.拒绝超过服务器承载能力的流量

//限流器实现逻辑
//1.创建一个固定大小的桶（令牌）
//2.按一定速率对桶里的令牌进行补充
//3.每进行一次访问，减少桶里现有的令牌数
//4.令牌最多补充到桶的大小

type Limiter struct {
	tb *TokenBuket
}

type TokenBuket struct {
	//桶
	mu sync.Mutex
	//桶大小
	size int
	//当前token数
	count int
	//填充速率，即每隔多久补充一个token
	rateLimit time.Duration
	//最后成功请求的时间
	lastRequestTime time.Time
}

func (tb *TokenBuket) fillToken() {
	tb.count += tb.getFillTokenCount()
}

func (tb *TokenBuket) getFillTokenCount() int {
	//当桶中令牌的数量大于桶的大小，无需填充
	if tb.count > tb.size {
		return 0
	}
	//首次请求，无需填充
	if !tb.lastRequestTime.IsZero() {
		duration := time.Now().Sub(tb.lastRequestTime)
		count := int(duration / tb.rateLimit)
		//如果桶中剩余大小大于count
		if tb.size-tb.count >= count {
			return count
		}
		//如果桶中剩余大小小于count，填充数据为剩余空间
		return tb.size - tb.count
	}
	return 0
}

func (tb *TokenBuket) allow() bool {
	//填充
	tb.fillToken()
	if tb.count > 0 {
		tb.count--
		tb.lastRequestTime = time.Now()
		return true
	}
	return false
}

// 初始化限流器
func NewLimiter(r time.Duration, size int) *Limiter {
	return &Limiter{
		tb: &TokenBuket{
			rateLimit: r,
			size:      size,
			count:     size,
		},
	}
}

func (l *Limiter) Allow() bool {
	l.tb.mu.Lock()
	defer l.tb.mu.Unlock()
	//计算补充token数
	//当前token是否满足本次消耗
	return l.tb.allow()
}

package breaker

import (
	"errors"
	"sync"
	"time"
)

const (
	STATE_CLOSE = iota
	STATE_OPEN
	STATE_HALF_OPEN
)

type Breaker struct {
	mu sync.Mutex
	//状态
	state int
	//失败次数阈值
	failureThreshold int
	//成功次数阈值
	successThreshold int
	//半开状态最多请求次数
	halfMaxRequest int
	//半开状态周期内目前请求次数
	halfCycleReqCount int
	//时间周期（正常请求时间周期和断开状态超时时间）
	timeout time.Duration
	//连续失败次数
	failureCount int
	//连续成功次数
	sucessCount int
	//周期起始时间
	cycleStartTime time.Time
}

func NewBreaker(failureThreshold, successThreshold, halfMaxRequest int, timeout time.Duration) *Breaker {
	return &Breaker{
		state:            STATE_CLOSE,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		halfMaxRequest:   halfMaxRequest,
		timeout:          timeout,
	}
}

func (b *Breaker) Exec(f func() error) error {
	b.before()
	if b.state == STATE_OPEN {
		return errors.New("断路器处于打开状态无法访问服务")
	}
	if b.state == STATE_CLOSE {
		err := f()
		b.after(err)
		return err
	}
	//半开状态
	if b.state == STATE_HALF_OPEN {
		//半开请求数量小于半开请求总数，执行请求
		if b.halfCycleReqCount < b.halfMaxRequest {
			err := f()
			b.after(err)
			return err
		} else {
			return errors.New("断路器处于半开状态，单位时间内请求次数太多，请稍后再试")
		}
	}
	return nil
}

func (b *Breaker) before() {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch b.state {
	case STATE_OPEN:
		//如果周期开始+超时时间在当前时间之前，表示超过一个周期了
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			//状态改为半开状态
			b.state = STATE_HALF_OPEN
			//重新计数
			b.reset()
			return
		}
	case STATE_HALF_OPEN:
		//如果连续成功次数大于阈值
		if b.sucessCount >= b.successThreshold {
			//修改状态为闭合状态
			b.state = STATE_CLOSE
			b.reset()
			return
		}
		//如果已经超过一个周期
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			//重置周期开始时间
			b.cycleStartTime = time.Now()
			//半开请求次数重置
			b.halfCycleReqCount = 0
			return
		}
	case STATE_CLOSE:
		//超过一个周期，重置
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.reset()
		}
	}
}

func (b *Breaker) after(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err == nil {
		b.onSuccess()
	} else {
		b.onFailure()
	}
}

func (b *Breaker) onSuccess() {
	//连续失败次数清零
	b.failureCount = 0
	//只有关闭状态和半开状态会进入到这里
	//关闭状态无需任何处理
	//半开状态增加成功请求次数和半开状态请求次数
	if b.state == STATE_HALF_OPEN {
		b.sucessCount++
		b.halfCycleReqCount++
		//如果成功请求次数大于阈值
		if b.sucessCount >= b.successThreshold {
			//状态设置为关闭并重置
			b.state = STATE_CLOSE
			b.reset()
		}
	}
}

func (b *Breaker) onFailure() {
	b.sucessCount = 0
	b.failureCount++
	if b.state == STATE_HALF_OPEN || (b.state == STATE_CLOSE && b.failureCount >= b.failureThreshold) {
		//半开状态或者关闭状态，连续失败次数已经大于阈值
		//将状态设置为开启并重置
		b.state = STATE_OPEN
		b.reset()
		return
	}
}

// 重置重新计算周期
func (b *Breaker) reset() {
	b.sucessCount = 0
	b.failureCount = 0
	b.halfCycleReqCount = 0
	b.cycleStartTime = time.Now()
}

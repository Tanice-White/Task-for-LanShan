package pubsub

import (
	"sync"
	"time"
)

type (
	Subscriber chan interface{}         //订阅者为一个通道
	TopicFunc  func(v interface{}) bool //主题为一个过滤器
)

//Publisher 发布者对象
type Publisher struct {
	M           sync.RWMutex             //读写锁
	Buffer      int                      //订阅队列的缓存大小
	Timeout     time.Duration            //发布超时时间
	Subscribers map[Subscriber]TopicFunc //订阅者信息
}

//NewPublisher 构建发布者对象设置发布超时时间和缓存的队列长度
//参数 发布超时时间 缓存队列的长度
//返回值 对象的地址
func NewPublisher(PublishTimeOut time.Duration, buffer int) *Publisher {
	return &Publisher{
		Buffer:      buffer,
		Timeout:     PublishTimeOut,
		Subscribers: make(map[Subscriber]TopicFunc),
	}
}

//SubscribeTopic 增加一个订阅者 订阅过滤器筛选后的主题
//返回值 通道
func (p *Publisher) SubscribeTopic(topic TopicFunc) chan interface{} {
	ch := make(chan interface{}, p.Buffer)
	p.M.Lock()
	p.Subscribers[ch] = topic
	p.M.Unlock()
	return ch
}

//Subscribe 增加一个订阅者 订阅全部主题
//返回值 通道
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

//Evict 退出订阅
//参数 退出订阅的用户
func (p *Publisher) Evict(sub chan interface{}) {
	p.M.Lock()
	defer p.M.Unlock()
	delete(p.Subscribers, sub)
	close(sub)
}

//sendTopic 发送主题，可以容忍一定的超时
//参数 订阅者 主题 主题内容 waitGroup指针
func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select { //没有default语句则被阻塞知道某一情况发生再立即返回
	case sub <- v:
	case <-time.After(p.Timeout):
	}
}

//Publish 发布一个主题
//参数 主题内容
func (p *Publisher) Publish(v interface{}) {
	p.M.Lock()
	defer p.M.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.Subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//Close 关闭发布者对象 同时关闭所有通道
func (p *Publisher) Close() {
	p.M.Lock()
	defer p.M.Lock()
	for sub := range p.Subscribers {
		delete(p.Subscribers, sub)
		close(sub)
	}
}

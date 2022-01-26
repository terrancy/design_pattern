package golang

import "sync"

//单例模式
//单例模式采用饿汉模式和懒汉模式.懒汉模式支持延迟加载,但只是把冷却时间加到第一次使用,而且不可避免的需要加锁

type Singleton struct {
}

// 饿汉模式
var hungerSingleton *Singleton

func init() {
    hungerSingleton = &Singleton{}
}
func GetHungerInstance() *Singleton {
    return hungerSingleton
}

// 懒汉模式
var lazySingleton *Singleton
var once = &sync.Once{}

func GetLazyInstance() *Singleton {
    if nil == lazySingleton {
        once.Do(func() {
            lazySingleton = &Singleton{}
        })
    }
    return lazySingleton
}

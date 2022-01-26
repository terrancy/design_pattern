package test

import (
    "design/pattern/golang"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetHunger(t *testing.T) {
    assert.Equal(t, golang.GetHungerInstance(), golang.GetHungerInstance())
}

//
//  BenchmarkGetHungerInstanceParallel
//  @Description: 恶汉模式.并行测试
//  @param b
//
func BenchmarkGetHungerInstanceParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            if golang.GetHungerInstance() != golang.GetHungerInstance() {
                b.Errorf("hunger fail...")
            }
        }
    })
}

//
//  BenchmarkGetLazyInstanceParallel
//  @Description: 懒汉模式
//  @param b
//
func BenchmarkGetLazyInstanceParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            if golang.GetLazyInstance() != golang.GetHungerInstance() {
                b.Errorf("lasy fail...")
            }
        }
    })
}

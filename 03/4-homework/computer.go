package main

import (
	"fmt"
)

//=======抽象层=======

type AbstractCard interface {
	Display()
}
type AbstractMemory interface {
	Storage()
}
type AbstractCpu interface {
	Calculate()
}

//AbstractFactory 抽象工厂-产品等级结构
type AbstractFactory interface {
	CreateCard() AbstractCard
	CreateMemory() AbstractMemory
	CreateCpu() AbstractCpu
}

//=======实现层=======

//Intel厂商产品族

type IntelCard struct{}

func (it *IntelCard) Display() {
	fmt.Println("Intel 显卡具有显示功能")
}

type IntelMemory struct{}

func (it *IntelMemory) Storage() {
	fmt.Println("Intel 内存具有存储功能")
}

type IntelCpu struct{}

func (it *IntelCpu) Calculate() {
	fmt.Println("Intel Cpu具有计算功能")
}

type IntelFactory struct{}

func (it *IntelFactory) CreateCard() AbstractCard {
	var abstractCard AbstractCard
	abstractCard = new(IntelCard)
	return abstractCard
}
func (it *IntelFactory) CreateMemory() AbstractMemory {
	var abstractMemory AbstractMemory
	abstractMemory = new(IntelMemory)
	return abstractMemory
}
func (it *IntelFactory) CreateCpu() AbstractCpu {
	var abstractCpu AbstractCpu
	abstractCpu = new(IntelCpu)
	return abstractCpu
}

//Nvidia厂商产品族

type NvidiaCard struct{}

func (nd *NvidiaCard) Display() {
	fmt.Println("Nvidia 显卡具有显示功能")
}

type NvidiaMemory struct{}

func (nd *NvidiaMemory) Storage() {
	fmt.Println("Nvidia 内存具有存储功能")
}

type NvidiaCpu struct{}

func (nd *NvidiaCpu) Calculate() {
	fmt.Println("Nvidia Cpu具有计算功能")
}

type NvidiaFactory struct{}

func (nd *NvidiaFactory) CreateCard() AbstractCard {
	var abstractCard AbstractCard
	abstractCard = new(NvidiaCard)
	return abstractCard
}
func (nd *NvidiaFactory) CreateMemory() AbstractMemory {
	var abstractMemory AbstractMemory
	abstractMemory = new(NvidiaMemory)
	return abstractMemory
}
func (nd *NvidiaFactory) CreateCpu() AbstractCpu {
	var abstractCpu AbstractCpu
	abstractCpu = new(NvidiaCpu)
	return abstractCpu
}

//Kingston厂商产品族

type KingstonCard struct{}

func (kg *KingstonCard) Display() {
	fmt.Println("Kingston 显卡具有显示功能")
}

type KingstonMemory struct{}

func (kg *KingstonMemory) Storage() {
	fmt.Println("Kingston 内存具有存储功能")
}

type KingstonCpu struct{}

func (kg *KingstonCpu) Calculate() {
	fmt.Println("Kingston Cpu具有计算功能")
}

type KingstonFactory struct{}

func (kg *KingstonFactory) CreateCard() AbstractCard {
	var abstractCard AbstractCard
	abstractCard = new(KingstonCard)
	return abstractCard
}
func (kg *KingstonFactory) CreateMemory() AbstractMemory {
	var abstractMemory AbstractMemory
	abstractMemory = new(KingstonMemory)
	return abstractMemory
}
func (kg *KingstonFactory) CreateCpu() AbstractCpu {
	var abstractCpu AbstractCpu
	abstractCpu = new(KingstonCpu)
	return abstractCpu
}

//Computer 电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口
type Computer struct {
	Cpu    AbstractCpu
	Card   AbstractCard
	Memory AbstractMemory
}

func main() {
	//创建Intel工厂
	var intelFactory AbstractFactory
	intelFactory = new(IntelFactory)
	//组装两台电脑
	//1台（Intel的CPU，Intel的显卡，Intel的内存）
	computer1 := Computer{
		Cpu:    intelFactory.CreateCpu(),
		Card:   intelFactory.CreateCard(),
		Memory: intelFactory.CreateMemory(),
	}
	computer1.Cpu.Calculate()
	computer1.Card.Display()
	computer1.Memory.Storage()
	//创建Nvidia工厂
	var nvidiaFactory AbstractFactory
	nvidiaFactory = new(NvidiaFactory)
	//创建Kingston工厂
	var kingstonFactory AbstractFactory
	kingstonFactory = new(KingstonFactory)
	// 1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	computer2 := Computer{
		Cpu:    intelFactory.CreateCpu(),
		Card:   nvidiaFactory.CreateCard(),
		Memory: kingstonFactory.CreateMemory(),
	}
	computer2.Cpu.Calculate()
	computer2.Card.Display()
	computer2.Memory.Storage()
}

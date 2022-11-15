package Pipeline

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Container"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/IPipeline"
)

type Pipeline struct {
	container IContainer.IContainer

	passable interface{}

	pipes []IContainer.MagicalFunc
}

var PipeArgumentError = errors.New("pipe parameters must have a return value")

type Callback func(stack IPipeline.Pipe, next IContainer.MagicalFunc) IPipeline.Pipe

func New(container IContainer.IContainer) IPipeline.Pipeline {
	return &Pipeline{
		container: container,
	}
}

func (this *Pipeline) Send(passable interface{}) IPipeline.Pipeline {
	this.passable = passable
	return this
}

func (this *Pipeline) Through(pipes ...interface{}) IPipeline.Pipeline {
	for _, item := range pipes {
		pipe, isStaticFunc := item.(IContainer.MagicalFunc)
		if !isStaticFunc {
			pipe = Container.NewMagicalFunc(item)
		}
		if pipe.NumOut() != 1 {
			panic(PipeArgumentError)
		}
		this.pipes = append(this.pipes, pipe)
	}
	return this
}

func (this *Pipeline) Then(destination interface{}) interface{} {
	return this.ArrayReduce(
		this.reversePipes(), this.carry(), this.prepareDestination(destination),
	)(this.passable)
}

func (this *Pipeline) carry() Callback {
	return func(stack IPipeline.Pipe, next IContainer.MagicalFunc) IPipeline.Pipe {
		return func(passable interface{}) interface{} {
			return this.container.StaticCall(next, passable, stack)[0]
		}
	}
}

func (this *Pipeline) ArrayReduce(pipes []IContainer.MagicalFunc, callback Callback, current IPipeline.Pipe) IPipeline.Pipe {
	for _, magicalFunc := range pipes {
		current = callback(current, magicalFunc)
	}
	return current
}

func (this *Pipeline) reversePipes() []IContainer.MagicalFunc {
	for from, to := 0, len(this.pipes)-1; from < to; from, to = from+1, to-1 {
		this.pipes[from], this.pipes[to] = this.pipes[to], this.pipes[from]
	}
	return this.pipes
}

func (this *Pipeline) prepareDestination(destination interface{}) IPipeline.Pipe {
	pipe, isStaticFunc := destination.(IContainer.MagicalFunc)
	if !isStaticFunc {
		pipe = Container.NewMagicalFunc(destination)
	}
	if pipe.NumOut() != 1 {
		panic(PipeArgumentError)
	}
	return func(passable interface{}) interface{} {
		return this.container.StaticCall(pipe, passable)[0]
	}
}

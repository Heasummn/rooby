# goby (rooby)

[![Join the chat at https://gitter.im/rooby-lang/Lobby](https://badges.gitter.im/rooby-lang/Lobby.svg)](https://gitter.im/rooby-lang/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![Build Status](https://travis-ci.org/goby-lang/goby.svg?branch=master)](https://travis-ci.org/goby-lang/goby)
[![Code Climate](https://codeclimate.com/github/goby-lang/goby/badges/gpa.svg)](https://codeclimate.com/github/goby-lang/goby)
[![GoDoc](https://godoc.org/github.com/goby-lang/goby?status.svg)](https://godoc.org/github.com/goby-lang/goby)
[![Go Report Card](https://goreportcard.com/badge/github.com/goby-lang/goby)](https://goreportcard.com/report/github.com/goby-lang/goby)
[![codecov](https://codecov.io/gh/goby-lang/goby/branch/master/graph/badge.svg)](https://codecov.io/gh/goby-lang/goby)
[![BCH compliance](https://bettercodehub.com/edge/badge/goby-lang/goby?branch=master)](https://bettercodehub.com/)
[![Readme Score](http://readme-score-api.herokuapp.com/score.svg?url=goby-lang/goby)](http://clayallsopp.github.io/readme-score?url=goby-lang/goby)

goby is a Ruby-like object oriented language written in Go. You can think it as a simplified, compilable Ruby for now.
   
**We're looking for contributors now, join us and make this project better!**

## Goal

I want to build a language that focuses on developing microservices. Which should be performant and easy to write. This is why goby has Ruby's user friendly syntax and is written in Go.

## Questions

A lot people have questions about `goby` since it's a new language and you may get confused by the way I describe it (sorry for that 😢). Here's a list of [frequently asked questions](https://github.com/goby-lang/goby/wiki/Frequently-asked-questions).

## Supported features
- **Can be compiled into bytecode (with `.robc` extension)**
- **Can evaluate bytecode directly**
- Everything is object
- Support comment 
- Object and Class
    - Top level main object
    - Constructor
    - Support class method
    - Support inheritance
    - Support instance variable
    - Support self
- Variables
    - Constant
    - Local variable
    - Instance variable
- Method
    - Support evaluation with arguments
    - Support evaluation without arguments
    - Support evaluation with block (closure)
- BuiltIn Data Types (All of them are classes 😀)
    - Class
    - Integer
    - String
    - Boolean
    - nil
    - Hash
    - Array
- Flow control
    - If statement
    - while statement
- Import other files
    - require_relative
- IO
    - `puts`
    - `ARGV`
    
**(You can open an issue for any feature request)**
    
## TODO & WIP

Checkout this [issue](https://github.com/goby-lang/goby/issues/72) for what we will work on before first release.

Also see [huboard](https://huboard.com/goby-lang/goby)

## Install

1. You must have Golang installed
2. You must have set $GOPATH
3. Add your $GOPATH/bin into $PATH
4. Run following command 

```
$ go get github.com/goby-lang/goby
```

## Usage

**Execute goby file using VM**

(might see errors on sample-6 since vm hasn't support block yet)
``` 
$ goby ./samples/sample-1.ro
#=> 16
```

**Compile goby code**

```
$ goby -c ./samples/sample-1.ro
```

You'll see `sample-1.robc` in `./samples`

**Execute bytecode**

```
$ goby ./samples/sample-1.robc
```


## Try it!
(See sample directory)
```
$ goby ./samples/sample-1.ro
$ goby ./samples/sample-2.ro
$ goby ./samples/sample-3.ro
$ goby ./samples/sample-4.ro
$ goby .....
```
## Development & Contribute

See the [guideline](https://github.com/goby-lang/goby/blob/master/CONTRIBUTING.md)

## References

I can't build this project without these resources, and I highly recommend you to check them out if you're interested in building your own languages

- [Write An Interpreter In Go](https://interpreterbook.com)
- [Nand2Tetris II](https://www.coursera.org/learn/nand2tetris2/home/welcome)
- [Ruby under a microscope](http://patshaughnessy.net/ruby-under-a-microscope)
- [YARV's instruction table](http://www.atdot.net/yarv/insnstbl.html)

## Maintainers

- @st0012
- @janczer
- @adlerhsieh

##  Sample snippet.

```ruby
class User
  def initialize(name, age)
    @name = name
    @age = age
  end

  def name
    @name
  end

  def age
    @age
  end

  def say_hi(user)
    puts(@name + " says hi to " + user.name)
  end

  def self.sum_age(user1, user2)
    user1.age + user2.age
  end
end

stan = User.new("Stan", 22)
john = User.new("John", 40)
puts(User.sum_age(stan, john)) #=> 62
stan.say_hi(john) #=> Stan says hi to John
```

#### Build a stack using goby

```ruby
class Stack
  def initialize
    @data = []
  end
    
  def push(x)
    @data.push(x)
  end
    
  def pop
    @data.pop
  end
    
  def top
    @data[@data.length - 1]
  end
end

s = Stack.new
s.push(1)
s.push(2)
s.push(3)
s.push(4)
s.push(10)
puts(s.pop) #=> 10
puts(s.top) #=> 4
```

#### Block support

```ruby
class Car
  def initialize
    yield(self)
  end
  
  def color=(c)
    @color = c
  end
  
  def color
    @color
  end
  
  def doors=(ds)
    @doors = ds
  end
  
  def doors
    @doors
  end
end
 
car = Car.new do |c|
  c.color = "Red"
  c.doors = 4
end
 
puts("My car's color is " + car.color + " and it's got " + car.doors.to_s + " doors.")

```

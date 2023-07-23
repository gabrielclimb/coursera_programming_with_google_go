### Week 1
- Advantages of go
	 - code runs fast 
	 - garbage collector

	![[software_translation.png]]
- Compiled Language:
	- a compiled language is a language where the translation from high-level language to machine code happens **one time before you execute** the code, before you deploy the code, happens one time
	- So, the idea behind a compiled languages, the key thing we want to bring out anyway is the fact that this translation occurs once, it doesn't occur while you're running the code
- Interpreted Language:
	- So, the idea behind a compiled languages, the key thing we want to bring out anyway is the fact that this translation occurs once, it doesn't occur while you're running the code.

Go language is object-oriented, but let's say, weakly object-oriented.

	![[objects_in_go.png]]

- Concurrency
	- Parallellism
		- when do start and stop?
		- conflict memory
	- Concurrent
		- Concurrency is the management of multiple tasks at the same time.
		- Concurrent programming enables parallelism
			- communication between tasks
			- sync between tasks
			- management of tasks execution
	- In Go
		- Goroutines
		- Channels


### Week 2
- Basic data types
- Pointers
	- is a address to data in memory
	- **&** operator returns the address of a variable/function
	- * operator returns data at an address (deferencing)  

```go
var x int = 1
var y int
var ip *int // ip is pointer to int

ip = &x // ip now points to x
y = *int // y is now 1
```

- Variable Scope
	- The places in code where a variable can be accessed.
	- ![[Variable Scope.png]] in the right we gonna have problem
- Blocks
	- ![[blocks_1.png]]
	- ![[blocks_2.png]]
- Lexical Scoping
	- b is a block 
	- ![[Lexical Scoping.png]]
- Deallocating Memory
	- ![[Deallocating_1.png]]
	- Stack and Heap
		- ![[Pasted image 20230723135852.png]]
		- Stack
			- the stack are the local variables for a function
			- So every time you call a function there can be variables that you define in that function and generally they go into the stack. They are allocated in the stack area of the memory and they're deallocated.
			- If they're allocated in the stack, they are deallocated automatically when the function completes.
		- Heap
			- **The heap** on the other hand, is a persistent region of memory where when you allocate something on the heap it doesn't go away just because the function that allocated it is complete, that heap memory it you have to explicitly deallocate it somehow in another language, like in C. Go dont do this
-  Garbage Collection
	- ![[Pasted image 20230723140228.png]]
	- ![[Pasted image 20230723142213.png]]
	- ![[Pasted image 20230723142535.png]]
		- This is a really helpful thing. Now, there is a downside because the act of garbage collection does take some time. Right. So, there's a performance hit but it's a pretty efficient implementation and garbage collection is so darn useful is probably worth it to put it in Go. So, that's a trade-off that Go makes. It slows things down a little bit but it is a great advantage because it makes programming a lot easier. You don't have to go as far as using a full-on interpreter like you would in an interpreted language.
- Types
	- comments are ignored by the compiler
	- Integers
		- int8, int16, int32, int64
		- The most common thing to do, is just to declare it as an int and leave it to the compiler figure out
	- Type conversion
		- ![[Pasted image 20230723143708.png]]
	- iota - iota is a function used to generate constants
		- You use this when you have to represent some property or some property that has several different distinct possible value. So, this is also known as one-hot. So, basically if you have a variable and you know it's going to be one-hot coded.
		- ![[iota1.png]]
		- ![[iota2.png]]
	- Control Flow
		- If / else
		```go
		if x > 5 {
			fmt.Printf("Yup, %s", x)
		}
	   ```
		- For
			- ![[for.png]]
		- Switch Case
			- ![[switch.png]]
		- Tagless Switch
			- ![[tagless_switch.png]]
		- Break and Continue
			- ![[break.png]]
			- ![[continue.png]]
		- Scan
			- reads user input
			- takes a pointer as an argument
			- ![[scan.png]]
			- 
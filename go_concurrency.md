---
marp: true
theme: gaia
class: lead
paginate: true
backgroundColor: #fff
---

# Go Concurrency 
## Doing More at Once
**Presented by:** Mehta Param

---

# What is Go?

---

# Developed by Google (2009)
Statically-typed & Compiled.

---

# The Goal
Efficient Compilation + Fast Execution + Ease of Use.

---

# The Problem: 
## Sequential Programming

---

# Task A ➡️ Task B ➡️ Task C
If A takes 5 seconds...
B and C are completely blocked.

---

# Modern CPUs have multiple cores.
Sequential code wastes them.

---

# The Solution: Concurrency

---

# Concurrency 
*Dealing* with multiple things at once.
Structuring the program.

---

# Parallelism
*Doing* multiple things at once.
Executing simultaneously.

---

# In Go...
Concurrency is a first-class citizen. 
It is built into the syntax.

---

# 1. Goroutines
The Heart of Go Concurrency.

---

# Traditional OS Thread
🪨 
Heavy. Takes ~1 MB of memory just to start.

---

# Goroutine
🪶
Lightweight. Takes ~2 KB of memory.

---

# CPU Core Handling OS Threads
🧠 ➡️ 🪨 (Struggling)

---

# CPU Core Handling Goroutines
🧠 ➡️ 🪶🪶🪶🪶🪶🪶🪶🪶🪶🪶 (Effortless)

---

# The Go Scheduler
Multiplexes thousands of Goroutines onto a few OS threads.

---

# How to use it?
Just type `go`

---

```go
// Normal
processTask()
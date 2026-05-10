````md id="kw8x2v"
# Distributed File Transfer & MapReduce using Go

## Overview

This project demonstrates a simple Distributed System using Go.

The system contains:

- File Sender
- File Receiver
- MapReduce Word Count

The project uses:
- Goroutines
- Channels
- HTTP
- Multipart Upload
- Concurrent Processing

---

# Project Structure

```text
project/
│
├── sender.go
├── receiver.go
├── mapreduce.go
├── uploads/
└── README.md
````

---

# Sender

## Purpose

The sender sends files from one computer to another using HTTP POST requests.

---

## Features

* Send multiple files
* Concurrent file reading
* Multipart upload
* Uses Channels and Goroutines

---

## Important Code

### Goroutine

```go id="jlwm6s"
go readFile(filePath, ch)
```

Used for concurrent execution.

---

### Channel

```go id="jlwm7g"
ch := make(chan []byte)
```

Transfers file data between Goroutines.

---

### Multipart Upload

```go id="4fjlwm"
writer.CreateFormFile("Files", fileName)
```

Creates upload form data.

---

# Sender Flow

```text
Read File
   ↓
Goroutine
   ↓
Channel
   ↓
HTTP POST
   ↓
Receiver
```

---

# Receiver

## Purpose

Receives uploaded files and saves them in:

```text
uploads/
```

---

## Endpoint

```text
/save
```

---

## Receiver Flow

```text
HTTP Request
     ↓
Parse Multipart Form
     ↓
Extract Files
     ↓
Save Files
```

---

# MapReduce

## Purpose

Processes:

```text
gene.fna
```

The system calculates:

* Total words
* Words repeated more than 10 times

---

# MapReduce Architecture

```text
Input File
    ↓
Scanner
    ↓
Lines
    ↓
Mapper Goroutines
    ↓
Channel
    ↓
Reducer
    ↓
Final Result
```

---

# Mapper

The mapper:

* Reads line
* Splits words
* Counts occurrences

Example:

```text
go go map
```

Result:

```text
go : 2
map : 1
```

---

# Reducer

The reducer:

* Receives mapper results
* Merges counts
* Produces final output

---

# Concurrency

## Goroutines

```go id="jlwm4p"
go mapper(...)
```

Used for parallel processing.

---

## WaitGroup

```go id="jlwm7t"
var wg sync.WaitGroup
```

Waits for all mappers.

---

## Channels

Used for safe communication between Goroutines.

---

# Running Receiver

```bash id="jlwm8k"
go run receiver.go
```

Receiver listens on:

```text
:9080
```

---

# Running Sender

Update receiver IP:

```go id="jlwm4g"
http://192.168.8.8:9080/save
```

Then run:

```bash id="jlwm6z"
go run sender.go
```

---

# Running MapReduce

```bash id="jlwm1x"
go run mapreduce.go
```

---

# Network Requirements

Both devices must:

* Be on same Wi-Fi/LAN
* Allow port 9080 through firewall

---

# Example IPs

Sender PC:

```text
192.168.8.15
```

Receiver PC:

```text
192.168.8.8
```

---

# Example Output

## Receiver

```text
Request received
Saved: gene.fna
```

---

## MapReduce

```text
Total Words: 50000

Words Count > 10:

gene : 120
dna : 55
sequence : 42
```

---

# Concepts Learned

This project demonstrates:

* Distributed Systems
* File Transfer
* HTTP Communication
* Goroutines
* Channels
* Parallel Processing
* Synchronization
* MapReduce

---

# Future Improvements

* Encryption
* Compression
* Multiple Nodes
* Fault Tolerance
* Distributed Storage
* Hadoop-style Distribution

---

# Conclusion

This project combines:

* Networking
* Concurrency
* Distributed Processing
* MapReduce

using Go.

```
```

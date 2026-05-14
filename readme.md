````markdown
# Distributed MapReduce System using Go

## Overview

This project demonstrates a simple Distributed MapReduce System using Go.

The system contains:

- Master Node
- Worker Nodes
- Distributed Processing
- HTTP Communication
- Concurrent Requests
- MapReduce Character Counting

The project uses:

- Goroutines
- Channels
- HTTP APIs
- JSON Communication
- Concurrent Processing
- Distributed Computing Concepts

---

# Project Architecture

```text
                    MASTER NODE
                          |
        ---------------------------------------
        |                 |                   |
        |                 |                   |
     Worker-1          Worker-2           Worker-3
    genome1.fa        genome2.fa         genome3.fa
````

---

# How The System Works

1. The Master Node sends HTTP requests to all Workers.
2. Each Worker processes its local genome file.
3. Workers count genome characters:

   * A
   * T
   * G
   * C
4. Each Worker sends its partial result back to the Master.
5. The Master combines all results.
6. Final result is displayed.

---

# MapReduce Flow

```text
Distributed Files
       |
       v
+-------------------+
|   Worker Nodes    |
| Local Processing  |
+-------------------+
       |
       v
Partial Results
       |
       v
+-------------------+
|    MASTER NODE    |
|   Reduce Phase    |
+-------------------+
       |
       v
 Final Combined Result
```

---

# Technologies Used

* Golang
* Goroutines
* Channels
* HTTP Server
* HTTP Client
* JSON
* Concurrent Programming
* Distributed Systems

---

# Project Structure

```text
project/
│
├── master/
│   └── main.go
│
├── worker/
│   └── main.go
│
├── data/
│   ├── genome1.fa
│   ├── genome2.fa
│   └── genome3.fa
│
└── README.md
```

---

# Worker Responsibilities

Each Worker:

* Stores a part of the dataset
* Runs local processing
* Counts genome characters
* Sends results to Master

Example:

```text
Worker-1 -> genome1.fa
Worker-2 -> genome2.fa
Worker-3 -> genome3.fa
```

---

# Master Responsibilities

The Master Node:

* Sends requests to Workers
* Receives all partial results
* Combines final counts
* Displays final output

---

# Example Final Output

```text
=================================
 FINAL RESULT
=================================

a : 120000
t : 118000
g : 115000
c : 121000

Total : 474000

MASTER FINISHED SUCCESSFULLY
```

---

# Features

* Distributed Processing
* Parallel Execution
* HTTP Communication
* Concurrent Workers
* Real-Time Logging
* JSON Data Exchange
* MapReduce Simulation
* Scalable Architecture

---

# Running The Project

## 1. Start Workers

Run on each Worker machine:

```bash
go run main.go
```

---

## 2. Configure Worker IPs

Example:

```go
workers := []string{
    "http://192.168.1.10:8081/count",
    "http://192.168.1.11:8082/count",
    "http://192.168.1.12:8083/count",
}
```

---

## 3. Run Master

```bash
go run main.go
```

---

# Example Worker Log

```text
[WORKER] Request Received From Master
[WORKER] Starting Genome Processing...
[WORKER] Finished Processing
[WORKER] Result Sent Successfully
```

---

# Example Master Log

```text
Connecting To: http://192.168.1.10:8081/count
Data Received From: http://192.168.1.10:8081/count
Result Received From Worker
```

---

# Concepts Demonstrated

* Distributed Systems
* MapReduce
* Parallel Computing
* Network Communication
* Client-Server Architecture
* Concurrent Programming
* Data Aggregation

---

# Future Improvements

* Fault Tolerance
* Dynamic Worker Discovery
* Heartbeat Monitoring
* Load Balancing
* Distributed File Storage
* Web Dashboard
* Docker Deployment

---

# Conclusion

This project demonstrates a simplified Distributed MapReduce System inspired by large-scale distributed processing systems such as Google MapReduce and Hadoop.

The system distributes computation across multiple Worker Nodes and combines results using a central Master Node.

```
```

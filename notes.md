12 Threads (Hardware threads)   3500 OS threads
                            M:N 
OS threads vs green threads
Context Switching
Scheduling
Preemptive Vs Cooperative Scheduling
8 mb 1 byte 8 mb
Thread pooling 

----------------------------------------------
CSP - go routines
- Users cannot create threads, there is not library to create a thread in Go
- can create a go routine (similar to green thread)
- context switching, scheduling , memory management is handled by runtime
- few thousand+ go routines can be created very easily
- very little memory footprint 2kb (can grow based on demand)


-----------------------------------------------------------------------
                      12 Threads(12 P --> logical process)

Global Queue
Local Queue
Work Stealing

# binder

binder分为用户空间的`libbinder`和内核里面的驱动`binder driver`，文件节点为`/dev/binder`，`libbinder`通过`ioctl`与`binder driver`交互。


## mmap

mmap(0, BINDER_VM_SIZE, PROT_READ, MAP_PRIVATE | MAP_NORESERVE, mDriverFD, 0);

1. 开始地址
2. 结束地址
3. 其他进程的读写权限
4. MAP的类型
5. fd
6. 内存的初始值

## History
mmap and associated systems calls were designed as part of the Berkeley Software Distribution (BSD) version of Unix. Their API was already described in the 4.2BSD System Manual, even though it was neither implemented in that release, nor in 4.3BSD.[1] Sun Microsystems had implemented this very API, though, in their SunOS operating system. The BSD developers at U.C. Berkeley requested Sun to donate its implementation, but these talks never led to any transfer of code; 4.3BSD-Reno was shipped instead with an implementation based on the virtual memory system of Mach.[2]

## File-backed and anonymous
File-backed mapping maps an area of the process's virtual memory to files; i.e. reading those areas of memory causes the file to be read. It is the default mapping type.

Anonymous mapping maps an area of the process's virtual memory not backed by any file. The contents are initialized to zero.[3] In this respect an anonymous mapping is similar to malloc, and is used in some malloc(3) implementations for certain allocations. However, anonymous mappings are not part of the POSIX standard, though implemented by almost all operating systems by the `MAP_ANONYMOUS` and `MAP_ANON` flags.

## Memory visibility
If the mapping is shared (the `MAP_SHARED` flag is set), then it is preserved across a fork(2) system call. This means that writes to a mapped area in one process are immediately visible in all related (parent, child or sibling) processes. If the mapping is shared and backed by a file (not `MAP_ANONYMOUS`) the underlying file media is only guaranteed to be written after it is msync(2)'ed.

If the mapping is private (the `MAP_PRIVATE` flag is set), the changes will neither be seen by other processes nor written to the file.

A process reading from or writing to the underlying file will not always see the same data as a process that has mapped the file, since the segment of the file is copied into RAM and periodically flushed to disk. Synchronization can be forced with the msync system call.

mmap(2)ing files can significantly reduce memory overhead for applications accessing the same file; they can share the memory area the file encompasses, instead of loading the file for each application that wants access to it. This means that mmap(2) is sometimes used for Interprocess Communication (IPC). On modern operating systems mmap(2) is typically preferred to the System V IPC Shared Memory facility.

The main difference between System V shared memory (shmem) and memory mapped I/O (mmap) is that SystemV shared memory is persistent: unless explicitly removed by a process, it is kept in memory and remains available until the system is shut down. mmap'd memory is not persistent between application executions (unless it is backed by a file).

# Linux 进程线程

## LWP(轻量级进程)

* pid 进程号
* lwp 线程号
* tgid 线程组id,等于pid

一个`lwp`对应一个内核线程，用户级线程绑定到`lwp`，从而被内核调度执行。

## 调度方式
0.NORMAL(普通进程CFS)
1.FIFO(进程主动放弃)
2.RR(时间片轮转)
3.BATCH(批处理)

1. **静态优先级(NI)** 不会随着时间而改变，内核不会修改它，只能通过系统调用nice去修改
2. **动态优先级(PRI)** 增加或者减小进程静态优先级的值来奖励IO小的进程或者惩罚cpu消耗型的进程，调整后的优先级称为动态优先级

```
cat /proc/915/schedstat 
193191802 1422708 529
```
* 第一个：该进程拥有的cpu的时间
* 第二个：在对列上的等待时间，即睡眠时间
* 第三个：被调度的次数

## 互斥实现
`CAS`实现互斥(乐观锁)
`Java Synchronized`关键字会挂起线程(悲观锁)，即把线程加入挂起队列(每个资源都会有一个挂起队列)，等资源释放的时候，再把线程加入全局就绪队列等待
调度，有点像订阅模式。

## 进程结构
cat /proc/17064/maps
```
00400000-01ade000 r-xp 00000000 fd:00 409599                             /usr/sbin/mysqld
01cdd000-01dcc000 r--p 016dd000 fd:00 409599                             /usr/sbin/mysqld
01dcc000-01e77000 rw-p 017cc000 fd:00 409599                             /usr/sbin/mysqld
01e77000-01f36000 rw-p 00000000 00:00 0 
02075000-02096000 rw-p 00000000 00:00 0                                  [heap]
02096000-02bda000 rw-p 00000000 00:00 0                                  [heap]
7fb438000000-7fb43829a000 rw-p 00000000 00:00 0 
7fb474d1c000-7fb474d5d000 rw-p 00000000 00:00 0                          [stack:17979]
7fb474d5d000-7fb474d5e000 ---p 00000000 00:00 0 
7fb474de1000-7fb474df7000 r-xp 00000000 fd:00 123351                     /usr/lib64/libresolv-2.17.so
7fb474df7000-7fb474ff6000 ---p 00016000 fd:00 123351                     /usr/lib64/libresolv-2.17.so
7fb474ff6000-7fb474ff7000 r--p 00015000 fd:00 123351                     /usr/lib64/libresolv-2.17.so
7fb474ff7000-7fb474ff8000 rw-p 00016000 fd:00 123351                     /usr/lib64/libresolv-2.17.so
7fb474ff8000-7fb474ffa000 rw-p 00000000 00:00 0 
7fb474ffa000-7fb474ffb000 ---p 00000000 00:00 0 
7fb474ffb000-7fb4757fb000 rw-p 00000000 00:00 0                          [stack:17072]
7fb4757fb000-7fb4757fc000 ---p 00000000 00:00 0 
7fb48c04d000-7fb48c052000 r-xp 00000000 fd:00 100769770                  /usr/lib64/mysql/plugin/validate_password.so
7fb48c052000-7fb48c252000 ---p 00005000 fd:00 100769770                  /usr/lib64/mysql/plugin/validate_password.so
7fb48c252000-7fb48c253000 r--p 00005000 fd:00 100769770                  /usr/lib64/mysql/plugin/validate_password.so
7fb48c253000-7fb48c254000 rw-p 00006000 fd:00 100769770                  /usr/lib64/mysql/plugin/validate_password.so
7fb48c254000-7fb48c47a000 rw-p 00000000 00:00 0 
7fb48c47a000-7fb48c47b000 ---p 00000000 00:00 0 
7fb48c47b000-7fb48cc7b000 rw-p 00000000 00:00 0                          [stack:17057]
7fb48cc7b000-7fb48cc7c000 ---p 00000000 00:00 0 
7fb4a5787000-7fb4a5789000 r-xp 00000000 fd:00 59459                      /usr/lib64/libfreebl3.so;5c89ece0 (deleted)
7fb4a5789000-7fb4a5988000 ---p 00002000 fd:00 59459                      /usr/lib64/libfreebl3.so;5c89ece0 (deleted)
7fb4a5988000-7fb4a5989000 r--p 00001000 fd:00 59459                      /usr/lib64/libfreebl3.so;5c89ece0 (deleted)
7fb4a5989000-7fb4a598a000 rw-p 00002000 fd:00 59459                      /usr/lib64/libfreebl3.so;5c89ece0 (deleted)
7fb4a71df000-7fb4a71e3000 rw-p 00000000 00:00 0 
7fb4a71e3000-7fb4a7205000 r-xp 00000000 fd:00 85                         /usr/lib64/ld-2.17.so
7fb4a720b000-7fb4a720c000 rw-p 00000000 00:00 0 
7fb4a7224000-7fb4a73f3000 rw-p 00000000 00:00 0 
7fb4a73f3000-7fb4a73fb000 rw-p 00000000 00:00 0 
7fb4a73fb000-7fb4a73fc000 rw-s 00000000 00:0a 5496615                    /[aio] (deleted)
7fb4a73fc000-7fb4a73ff000 rw-s 00000000 00:0a 5496606                    /[aio] (deleted)
7fb4a73ff000-7fb4a7402000 rw-s 00000000 00:0a 5496605                    /[aio] (deleted)
7fb4a7402000-7fb4a7403000 rw-s 00000000 00:0a 5496604                    /[aio] (deleted)
7fb4a7403000-7fb4a7404000 rw-p 00000000 00:00 0 
7fb4a7404000-7fb4a7405000 r--p 00021000 fd:00 85                         /usr/lib64/ld-2.17.so
7fb4a7405000-7fb4a7406000 rw-p 00022000 fd:00 85                         /usr/lib64/ld-2.17.so
7fb4a7406000-7fb4a7407000 rw-p 00000000 00:00 0 
7ffd3e5b1000-7ffd3e5d2000 rw-p 00000000 00:00 0                          [stack]
7ffd3e5f0000-7ffd3e5f2000 r-xp 00000000 00:00 0                          [vdso]
ffffffffff600000-ffffffffff601000 r-xp 00000000 00:00 0                  [vsyscall]

```
可以看到，`[text]/usr/sbin/mysqld` -> `[heap]` -> `[stack:17979]` -> `[mmap]/usr/lib64/libresolv-2.17.so` -> `[stack]` -> `[vdso]` -> `[vsyscall]`

[stack:17979]代表`LWP`的`stack`

## 进程创建的方式
1. `fork()`
2. `clone()`

`fork()`基于`clone()`修改，只是通过不同的参数决定是创建进程还是创建线程。






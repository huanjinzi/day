# java lock

```
monitorenter
...//临界区
monitorexit
```

## 自旋锁(synchronized)
硬件锁住内存总线，其他线程不能访问内存，所以只会有一个线程拿到锁(互斥锁)

## 可重入锁(ReentrantLock)
进入临界区需要记录线程ID，表明锁的归属，还用一个计数器来记录锁的重入次数；释放的时候依次释放，直到计数器为0才真正释放锁。

## 信号量(Semaphore)

## 读写锁(ReadWriteLock)

## ContetionList

## EntryList

## WaitSet

## wait()/notify()和park()/unpark()的区别
`wait()` 加入`WaitSet`，`notify()`加入`EntryList`

`park()` 设置`counter=0`,`unpark()` 设置`counter=1`;当`counter=1`时，调用`park()`会立即返回，同时设置`counter=0`





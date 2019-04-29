# GC

## GC算法

### 复制 copy
内存一分为二

### 标记清除 mark-swipe

### 标记整理 mark-compact

### 分代回收


## 分代

### Young
复制
#### Eden
#### Suvivor1
#### Survivor2


### Older
标记整理

### Perm(JVM规范叫方法区)


## 垃圾回收器

### 串行垃圾回收器
### 并行垃圾回收器

### 并发标记清除垃圾回收器

## Android GC

### Collector
```
MS    mark-swipe
CMS    concurrent-mark-swipe
SS    Semi-space mark-sweep hybrid, enables compaction
GSS    generational SS
MC   mark-compact
CC    concurrent copying

// A homogeneous space compaction collector used in background transition
// when both foreground and background collector are CMS.
HomogeneousSpaceCompact 
```

### Allocator
```
BumpPointer    Use BumpPointer allocator, has entrypoints.
TLAB    Thread Local Allocation Buffer.
RosAlloc    Use RosAlloc allocator, has entrypoints.
DlMalloc    Use dlmalloc allocator, has entrypoints.
NonMoving    Special allocator for non moving objects, doesn't have entrypoints.
LOS    Large object space, also doesn't have entrypoints.
Region
RegionTLAB
```

### GC Cause
```
Alloc
NativeAlloc
Background    A background GC trying to ensure there is free memory ahead of allocations.
Explicit    explicit System.gc() call.
CollectorTransition    collector transition.
HomogeneousSpaceCompact    background transition when both foreground and background collector are CMS.
ClassLinker    guard filling art methods with special values.
```






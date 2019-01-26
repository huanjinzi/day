# MessageQueue和Message

代码位置:
```
frameworks/base/core/java/android/os/MessageQueue.java
frameworks/base/core/java/android/os/Message.java
```

作为一个`MessageQueue`的基本方法：`入队(enqueue)`和`出队(dequeue)`：
```
// 入队
boolean enqueueMessage(Message msg, long when)
// 出队
Message next()
```
既然涉及到队列，一般有两种实现方式：`链表(linklist)`和`数组(array)`

`Message`作为`MessageQueue`中的元素，而且`MessageQueue`采用`链表(linklist)`实现。所以，
`Message`内部会有`next`指向下一个`Message`对象：
```
class Message {
    Message next;
}
```

## MessageQueue入队操作
`MessageQueue`的入队方法`enqueueMessage(Message msg, long when)`多了一个参数`long when`。请注意，`when`表示`Message`进入队列的时间，而不是
`Message`被工作线程处理的时间，请在这儿停顿并思考5秒钟。队列中的`Message`根据`when`由小到大排序，`when`越小，在队列中的位置越靠前面。这里的`when`为`android.os.SystemClock.uptimeMillis()`。所以，
正常情况下，`when > 0`，当然也存在`when = 0`的情况，当`msg`的`when == 0`的时候，`msg`会成为队列的`head`节点。

当调用`Handler.sendMessageAtFrontOfQueue(Message msg)`时候，会调用`MessageQueue.enqueueMessage(Message msg, long when)`，
这里传递的`when=0`，所以这里的`msg`会成为队列的`head`节点。

`enqueueMessage`方法解析：
```
boolean enqueueMessage(Message msg, long when) {
        if (msg.target == null) { //msg.target是消息的目的地，其实就是Handler
            throw new IllegalArgumentException("Message must have a target.");
        }
        if (msg.isInUse()) {
            throw new IllegalStateException(msg + " This message is already in use.");
        }

        synchronized (this) { // this为MessageQueue对象
            if (mQuitting) {
                IllegalStateException e = new IllegalStateException(
                        msg.target + " sending message to a Handler on a dead thread");
                Log.w(TAG, e.getMessage(), e);
                msg.recycle();
                return false;
            }

            msg.markInUse();
            msg.when = when;
            Message p = mMessages;// mMessages是链表的head
            boolean needWake;
            /*1111111111111111111111111111111111111111111*/
            if (p == null || when == 0 || when < p.when) {
                // New head, wake up the event queue if blocked.
                /*
                    这里有三个条件:
                        1.p == null //p(p = mMessages)是head节点，head为null,说明链表为空，这时进入链表的节点就直接成为head节点(New Head)
                        2.when == 0 //调用sendMessageAtFrontOfQueue时，when=0，表示进来的msg需要成为head结点(New Head)
                        3.when < p.when //新进入链表的msg.when < head.when，所以新进入的消息成为head，head.when是整个链表中最小的(New Head)
                 */
                msg.next = p;
                mMessages = msg;
                
                /* 
                 * mBlocked为true时，说明此时msg都被处理完了，head结点插入后，需要叫醒工作线程，通知有新的msg要处理。
                 */
                needWake = mBlocked;
            } 
            /*1111111111111111111111111111111111111111111*/
            
            
            /*222222222222222222222222222222222222222222*/
            else {
                /*
                    msg链表插入
                
                 */
                // Inserted within the middle of the queue.  Usually we don't have to wake
                // up the event queue unless there is a barrier at the head of the queue
                // and the message is the earliest asynchronous message in the queue.
                needWake = mBlocked && p.target == null && msg.isAsynchronous();
                Message prev;
                for (;;) {
                    prev = p;
                    p = p.next;
                    if (p == null || when < p.when) {
                        /*
                           1. p==null,说明msg被插在队列的最后面
                           2. when < p.when,找到when的插入位置
                         */
                        break;
                    }
                    if (needWake && p.isAsynchronous()) {
                        needWake = false;
                    }
                }
                msg.next = p; // invariant: p == prev.next
                prev.next = msg;
            }
            /*222222222222222222222222222222222222222222*/


            /*33333333333333333333333333333333333333333333333333*/
            // We can assume mPtr != 0 because mQuitting is false.
            if (needWake) {
                nativeWake(mPtr);
            }
            /*33333333333333333333333333333333333333333333333333*/
        }
        return true;
    }
```
从代码结构上面来看，代码大体分为3个部分：`/*111*/`，`/*222*/`和`/*333*/`，其中`/*111*/`、`/*222*/`是队列的插入操作，也就是说有新的`Message`
进入队列，`/*333*/`决定是否唤醒工作线程。

## MessageQueue出队操作
```
    Message next() {
        // Return here if the message loop has already quit and been disposed.
        // This can happen if the application tries to restart a looper after quit
        // which is not supported.
        final long ptr = mPtr;
        if (ptr == 0) {
            return null;
        }

        int pendingIdleHandlerCount = -1; // -1 only during first iteration
        int nextPollTimeoutMillis = 0;
        for (;;) {
            if (nextPollTimeoutMillis != 0) {
                Binder.flushPendingCommands();
            }

            nativePollOnce(ptr, nextPollTimeoutMillis);

            synchronized (this) {
                // Try to retrieve the next message.  Return if found.
                final long now = SystemClock.uptimeMillis();
                Message prevMsg = null;
                Message msg = mMessages;
                if (msg != null && msg.target == null) {
                    // Stalled by a barrier.  Find the next asynchronous message in the queue.
                    do {
                        prevMsg = msg;
                        msg = msg.next;
                    } while (msg != null && !msg.isAsynchronous());
                }
                if (msg != null) {
                    if (now < msg.when) {
                        // Next message is not ready.  Set a timeout to wake up when it is ready.
                        nextPollTimeoutMillis = (int) Math.min(msg.when - now, Integer.MAX_VALUE);
                    } else {
                        // Got a message.
                        mBlocked = false;
                        if (prevMsg != null) {
                            prevMsg.next = msg.next;
                        } else {
                            mMessages = msg.next;
                        }
                        msg.next = null;
                        if (DEBUG) Log.v(TAG, "Returning message: " + msg);
                        msg.markInUse();
                        return msg;
                    }
                } else {
                    // No more messages.
                    nextPollTimeoutMillis = -1;
                }

                // Process the quit message now that all pending messages have been handled.
                if (mQuitting) {
                    dispose();
                    return null;
                }

                // If first time idle, then get the number of idlers to run.
                // Idle handles only run if the queue is empty or if the first message
                // in the queue (possibly a barrier) is due to be handled in the future.
                if (pendingIdleHandlerCount < 0
                        && (mMessages == null || now < mMessages.when)) {
                    pendingIdleHandlerCount = mIdleHandlers.size();
                }
                if (pendingIdleHandlerCount <= 0) {
                    // No idle handlers to run.  Loop and wait some more.
                    mBlocked = true;
                    continue;
                }

                if (mPendingIdleHandlers == null) {
                    mPendingIdleHandlers = new IdleHandler[Math.max(pendingIdleHandlerCount, 4)];
                }
                mPendingIdleHandlers = mIdleHandlers.toArray(mPendingIdleHandlers);
            }

            // Run the idle handlers.
            // We only ever reach this code block during the first iteration.
            for (int i = 0; i < pendingIdleHandlerCount; i++) {
                final IdleHandler idler = mPendingIdleHandlers[i];
                mPendingIdleHandlers[i] = null; // release the reference to the handler

                boolean keep = false;
                try {
                    keep = idler.queueIdle();
                } catch (Throwable t) {
                    Log.wtf(TAG, "IdleHandler threw exception", t);
                }

                if (!keep) {
                    synchronized (this) {
                        mIdleHandlers.remove(idler);
                    }
                }
            }

            // Reset the idle handler count to 0 so we do not run them again.
            pendingIdleHandlerCount = 0;

            // While calling an idle handler, a new message could have been delivered
            // so go back and look again for a pending message without waiting.
            nextPollTimeoutMillis = 0;
        }
    }
```

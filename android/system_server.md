# SystemServer

## RemoteService
对第三方提供的系统服务
### ActivityManagerService
onStart()
addAppLocked()
startProcessLocked()
Process.start()

## LocalService
内部使用的服务

## Watchdog
`Watchdog`是一个线程，首先在`SystemServer`的`startOtherServices()`中调用`init()`，接着在`ActivityManagerService`的
`systemReady()`的回调方法中`start()`方法，启动`Watchdog`线程。

```
final Watchdog watchdog = Watchdog.getInstance();
watchdog.init(context, mActivityManagerService);
```

1.scheduleCheck() 记录开始时间StartTime
2.post(new HandlerChecker()) 到被监控的线程，被监控的线程依次执行`Monitor.monitor()`
3.标记HandlerChecker Complete

### HandlerChecker


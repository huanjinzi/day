# SystemServer

## RemoteService
对第三方提供的系统服务
### ActivityManagerService
onStart()
addAppLocked()
startProcessLocked()
Process.start()

## LocalService
```
SystemServiceManager
```

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


## 启动过程
1.将时间设置为`1970年`
2.关闭`Heap`大小限制，`clearGrowthLimit()`，并且调节`TargetHeapUtilization(0.8f)`，也就是说，当`Heap`内存利用率达到`80%`时候，分配新的内存。
3.配置`ro.build.fingerprint`属性
4.准备`Looper.prepareMainLooper()`
5.加载`libandroid_servers.so`

6.`startBootstrapServices`
7.`startCoreServices`
8.`startOtherServices()`
9.`systemReady()`
10.`Looper.loop()`

## 服务类型
```
SystemServiceManager SystemService
ServiceManager/LocalServices IBinder
```

## 6.startBootstrapServices
1.`Installer`
```
    Installer installer = mSystemServiceManager.startService(Installer.class)
```
2.`ActivityManagerService`
```
mActivityManagerService = mSystemServiceManager.startService(ActivityManagerService.Lifecycle.class).getService();
mActivityManagerService.setSystemServiceManager(mSystemServiceManager);
mActivityManagerService.setInstaller(installer);
```

3.`PowerManagerService`
```
 mPowerManagerService = mSystemServiceManager.startService(PowerManagerService.class);
 mActivityManagerService.initPowerManagement();
```

4.`LightService`
```
mSystemServiceManager.startService(LightsService.class);
```
5.`DisplayManagerService`
```
mDisplayManagerService = mSystemServiceManager.startService(DisplayManagerService.class);
```

6.`PackageManagerService`
```
mSystemServiceManager.startBootPhase(SystemService.PHASE_WAIT_FOR_DEFAULT_DISPLAY);
String cryptState = SystemProperties.get("vold.decrypt");
if (ENCRYPTING_STATE.equals(cryptState)) {
    Slog.w(TAG, "Detected encryption in progress - only parsing core apps");
    mOnlyCore = true;
} else if (ENCRYPTED_STATE.equals(cryptState)) {
    Slog.w(TAG, "Device encrypted - only parsing core apps");
    mOnlyCore = true;
}

// Start the package manager.
traceBeginAndSlog("StartPackageManagerService");
mPackageManagerService = PackageManagerService.main(mSystemContext, installer,
mFactoryTestMode != FactoryTest.FACTORY_TEST_OFF, mOnlyCore);
mFirstBoot = mPackageManagerService.isFirstBoot();
mPackageManager = mSystemContext.getPackageManager();
Trace.traceEnd(Trace.TRACE_TAG_SYSTEM_SERVER);
```

7.`OtaDexoptService`
```
if (!mOnlyCore) {
    boolean disableOtaDexopt = SystemProperties.getBoolean("config.disable_otadexopt",false);
    if (!disableOtaDexopt) {
        traceBeginAndSlog("StartOtaDexOptService");
            try {
                OtaDexoptService.main(mSystemContext, mPackageManagerService);
            } catch (Throwable e) {
                reportWtf("starting OtaDexOptService", e);
            } finally {
                Trace.traceEnd(Trace.TRACE_TAG_SYSTEM_SERVER);
            }
    }
}
```

8.`UserManagerService`
```
mSystemServiceManager.startService(UserManagerService.LifeCycle.class);
```

9.`SensorService`
```
startSensorService();
```

## 7.startCoreServices
1.`BatteryService`
```
mSystemServiceManager.startService(BatteryService.class);
```
2.`UsageStatsService`
```
mSystemServiceManager.startService(UsageStatsService.class);
mActivityManagerService.setUsageStatsManager(
LocalServices.getService(UsageStatsManagerInternal.class));
```

3.`WebViewUpdateService`
```
mWebViewUpdateService = mSystemServiceManager.startService(WebViewUpdateService.class);
```

## 8.startOtherServices
1.`SchedulingPolicyService`
```
ServiceManager.addService("scheduling_policy", new SchedulingPolicyService());
```

2.`TelephonyRegistry`
```
telephonyRegistry = new TelephonyRegistry(context);
ServiceManager.addService("telephony.registry", telephonyRegistry);
```

3.`EntropyMixer`
```
mEntropyMixer = new EntropyMixer(context);
```

4.`CameraService`
```
if (!disableCameraService) {
    Slog.i(TAG, "Camera Service");
    mSystemServiceManager.startService(CameraService.class);
}
```

5.`ACCOUNT_SERVICE`
```
mSystemServiceManager.startService(ACCOUNT_SERVICE_CLASS);
```

6.`CONTENT_SERVICE`
```
mSystemServiceManager.startService(CONTENT_SERVICE_CLASS);
```

7.`SystemProviders`
```
mActivityManagerService.installSystemProviders();
```

8.`VibratorService`
```
vibrator = new VibratorService(context);
```
9.`ConsumerIrService`
```
if (!disableConsumerIr) {
    consumerIr = new ConsumerIrService(context);
    ServiceManager.addService(Context.CONSUMER_IR_SERVICE, consumerIr);
}
```

10.`AlarmManagerService`
```
mSystemServiceManager.startService(AlarmManagerService.class);
```

11.`Watchdog`
```
final Watchdog watchdog = Watchdog.getInstance();
watchdog.init(context, mActivityManagerService);
```

12.`InputManagerService`
```
inputManager = new InputManagerService(context);
ServiceManager.addService(Context.INPUT_SERVICE, inputManager);
```

13.`WindowManagerService`
```
wm = WindowManagerService.main(context, inputManager,mFactoryTestMode != FactoryTest.FACTORY_TEST_LOW_LEVEL,!mFirstBoot, mOnlyCore);
ServiceManager.addService(Context.WINDOW_SERVICE, wm);
```

14.`VrManagerService`
```
if (!disableVrManager) {
    traceBeginAndSlog("StartVrManagerService");
    mSystemServiceManager.startService(VrManagerService.class);
}
```

```
mActivityManagerService.setWindowManager(wm);  
inputManager.setWindowManagerCallbacks(wm.getInputMonitor());
inputManager.start();

// TODO: Use service dependencies instead.
mDisplayManagerService.windowManagerAndInputReady();
```

15.`BluetoothService`
```
if (isEmulator) {
    Slog.i(TAG, "No Bluetooth Service (emulator)");
} else if (mFactoryTestMode == FactoryTest.FACTORY_TEST_LOW_LEVEL) {
    Slog.i(TAG, "No Bluetooth Service (factory test)");
    } else if (!context.getPackageManager().hasSystemFeature(PackageManager.FEATURE_BLUETOOTH)) {
    Slog.i(TAG, "No Bluetooth Service (Bluetooth Hardware Not Present)");
} else if (disableBluetooth) {
    Slog.i(TAG, "Bluetooth Service disabled by config");
} else {
    mSystemServiceManager.startService(BluetoothService.class);
}
```

16.`MetricsLoggerService`
```
mSystemServiceManager.startService(MetricsLoggerService.class);
```

17.`IpConnectivityMetrics`
```
mSystemServiceManager.startService(IpConnectivityMetrics.class);
```

18.`PinnerService`
```
mSystemServiceManager.startService(PinnerService.class);
```

19.`InputMethodManagerService`
```
mSystemServiceManager.startService(InputMethodManagerService.Lifecycle.class);
```

20.`AccessibilityManagerService`
```
ServiceManager.addService(Context.ACCESSIBILITY_SERVICE,new AccessibilityManagerService(context));
```

```
wm.displayReady();
```

21.`MountService`
```
if (mFactoryTestMode != FactoryTest.FACTORY_TEST_LOW_LEVEL) {
    if (!disableStorage && !"0".equals(SystemProperties.get("system_init.startmountservice"))) {
        try {
            /*
            * NotificationManagerService is dependant on MountService,
            * (for media / usb notifications) so we must start MountService first.
            */
            mSystemServiceManager.startService(MOUNT_SERVICE_CLASS);
            mountService = IMountService.Stub.asInterface(ServiceManager.getService("mount"));
        } catch (Throwable e) {
            reportWtf("starting Mount Service", e);
        }
    }
}
```

22.`UiModeManagerService`
```
// We start this here so that we update our configuration to set watch or television
// as appropriate.
mSystemServiceManager.startService(UiModeManagerService.class);
```

23.`LockSettingsService`
```
mSystemServiceManager.startService(LOCK_SETTINGS_SERVICE_CLASS);
lockSettings = ILockSettings.Stub.asInterface(ServiceManager.getService("lock_settings"));
```

24.`PersistentDataBlockService`
```
if (!SystemProperties.get(PERSISTENT_DATA_BLOCK_PROP).equals("")) {
    mSystemServiceManager.startService(PersistentDataBlockService.class);
}
```

25.`DeviceIdleController`
```
mSystemServiceManager.startService(DeviceIdleController.class);
```

26.`DevicePolicyManagerService`
```
// Always start the Device Policy Manager, so that the API is compatible with
// API8.
 mSystemServiceManager.startService(DevicePolicyManagerService.Lifecycle.class);
```

27.`StatusBarManagerService`
```
statusBar = new StatusBarManagerService(context, wm);
ServiceManager.addService(Context.STATUS_BAR_SERVICE, statusBar);
```

28.`ClipboardService`
```
ServiceManager.addService(Context.CLIPBOARD_SERVICE,new ClipboardService(context));
```

29.`NetworkManagementService`
```
networkManagement = NetworkManagementService.create(context);
ServiceManager.addService(Context.NETWORKMANAGEMENT_SERVICE, networkManagement);
```

30.`TextServicesManagerService`
```
mSystemServiceManager.startService(TextServicesManagerService.Lifecycle.class);
```

31.`NetworkScoreService`
```
networkScore = new NetworkScoreService(context);
ServiceManager.addService(Context.NETWORK_SCORE_SERVICE, networkScore);
```

32.`NetworkStatsService`
```
networkStats = NetworkStatsService.create(context, networkManagement);
ServiceManager.addService(Context.NETWORK_STATS_SERVICE, networkStats);
```

33.`NetworkPolicyManagerService`
```
networkPolicy = new NetworkPolicyManagerService(context,mActivityManagerService, networkStats, networkManagement);
ServiceManager.addService(Context.NETWORK_POLICY_SERVICE, networkPolicy);
```

34.`WIFINanService`
```
mSystemServiceManager.startService(WIFI_NAN_SERVICE_CLASS);
```

35.`WIFIP2PService`
```
mSystemServiceManager.startService(WIFI_P2P_SERVICE_CLASS);
```

36.`WIFIService`
```
mSystemServiceManager.startService(WIFI_SERVICE_CLASS);
```

37.`WifiScanningService`
```
mSystemServiceManager.startService("com.android.server.wifi.scanner.WifiScanningService");
```

38.`RttService`
```
mSystemServiceManager.startService("com.android.server.wifi.RttService");
```

39.`EthernetService`
```
if (mPackageManager.hasSystemFeature(PackageManager.FEATURE_ETHERNET) ||
    mPackageManager.hasSystemFeature(PackageManager.FEATURE_USB_HOST)) {
    mSystemServiceManager.startService(ETHERNET_SERVICE_CLASS);
}
```

40.`ConnectivityService`
```
connectivity = new ConnectivityService(
    context, networkManagement, networkStats, networkPolicy);
ServiceManager.addService(Context.CONNECTIVITY_SERVICE, connectivity);
networkStats.bindConnectivityManager(connectivity);
networkPolicy.bindConnectivityManager(connectivity);
```

41.`NsdService`
```
serviceDiscovery = NsdService.create(context);
ServiceManager.addService(Context.NSD_SERVICE, serviceDiscovery);
```

42.`UpdateLockService`
```
ServiceManager.addService(Context.UPDATE_LOCK_SERVICE,new UpdateLockService(context));
```

43.`RecoverySystemService`
```
mSystemServiceManager.startService(RecoverySystemService.class);
```

44.`NotificationManagerService`
```
mSystemServiceManager.startService(NotificationManagerService.class);
notification = INotificationManager.Stub.asInterface(
ServiceManager.getService(Context.NOTIFICATION_SERVICE));
networkPolicy.bindNotificationManager(notification);
```

45.`DeviceStorageMonitorService`
```
mSystemServiceManager.startService(DeviceStorageMonitorService.class);
```

46.`LocationManagerService`
```
location = new LocationManagerService(context);
ServiceManager.addService(Context.LOCATION_SERVICE, location);
```

47.`CountryDetectorService`
```
countryDetector = new CountryDetectorService(context);
ServiceManager.addService(Context.COUNTRY_DETECTOR, countryDetector);
```

48.`SearchManagerService`
```
mSystemServiceManager.startService(SEARCH_MANAGER_SERVICE_CLASS);
```

49.`DropBoxManagerService`
```
mSystemServiceManager.startService(DropBoxManagerService.class);
```

50.`WallPaperService`
```
mSystemServiceManager.startService(WALLPAPER_SERVICE_CLASS);
```

60.`AudioService`
```
mSystemServiceManager.startService(AudioService.Lifecycle.class);
```

61.`DockObserver`
```
mSystemServiceManager.startService(DockObserver.class);
```

62.`ThermalObserver`
```
mSystemServiceManager.startService(THERMAL_OBSERVER_CLASS);
```

63.`MIDIService`
```
mSystemServiceManager.startService(MIDI_SERVICE_CLASS);
```

64.`USBService`
```
mSystemServiceManager.startService(USB_SERVICE_CLASS);
```

65.`SerialService`
```
serial = new SerialService(context);
ServiceManager.addService(Context.SERIAL_SERVICE, serial);
```

66.`HardwarePropertiesManagerService`
```
hardwarePropertiesService = new HardwarePropertiesManagerService(context);
ServiceManager.addService(Context.HARDWARE_PROPERTIES_SERVICE,hardwarePropertiesService);
```

67.`TwilightService`
```
mSystemServiceManager.startService(TwilightService.class);
```

68.`NightDisplayService`
```
mSystemServiceManager.startService(NightDisplayService.class);
```

69.`JobSchedulerService`
```
mSystemServiceManager.startService(JobSchedulerService.class);
```

70.`SoundTriggerService`
```
mSystemServiceManager.startService(SoundTriggerService.class);
```

71.`BackupManagerService`
```
mSystemServiceManager.startService(BACKUP_MANAGER_SERVICE_CLASS);
```

72.`AppWidgetService`
```
mSystemServiceManager.startService(APPWIDGET_SERVICE_CLASS);
```

73.`VoiceRecognitionManagerService`
```
mSystemServiceManager.startService(VOICE_RECOGNITION_MANAGER_SERVICE_CLASS);
```

74.`GestureLauncherService`
```
mSystemServiceManager.startService(GestureLauncherService.class);
```

75.`SensorNotificationService`
```
mSystemServiceManager.startService(SensorNotificationService.class);
```

76.`ContextHubSystemService`
```
mSystemServiceManager.startService(ContextHubSystemService.class);
```

77.`DiskStatsService`
```
ServiceManager.addService("diskstats", new DiskStatsService(context));
```

78.`SamplingProfilerService`
```
ServiceManager.addService("samplingprofiler",new SamplingProfilerService(context));
```

79.`NetworkTimeUpdateService`
```
 networkTimeUpdater = new NetworkTimeUpdateService(context);
ServiceManager.addService("network_time_update_service", networkTimeUpdater);
```

80.`CommonTimeManagementService`
```
commonTimeMgmtService = new CommonTimeManagementService(context);
ServiceManager.addService("commontime_management", commonTimeMgmtService);
```

81.`EmergencyAffordanceService`
```
// EmergencyMode sevice
mSystemServiceManager.startService(EmergencyAffordanceService.class);
```

82.`DreamManagerService`
```
// Dreams (interactive idle-time views, a/k/a screen savers, and doze mode)
mSystemServiceManager.startService(DreamManagerService.class);
```

83.`AssetAtlasService`
```
atlas = new AssetAtlasService(context);
ServiceManager.addService(AssetAtlasService.ASSET_ATLAS_SERVICE, atlas);
```

84.`GraphicsStatsService`
```
ServiceManager.addService(GraphicsStatsService.GRAPHICS_STATS_SERVICE,new GraphicsStatsService(context));
```

85.`PrintManagerService`
```
mSystemServiceManager.startService(PRINT_MANAGER_SERVICE_CLASS);
```

86.`RestrictionsManagerService`
```
mSystemServiceManager.startService(RestrictionsManagerService.class);
```

87.`MediaSessionService`
```
mSystemServiceManager.startService(MediaSessionService.class);
```

88.`HdmiControlService`
```
mSystemServiceManager.startService(HdmiControlService.class);
```

89.`TvInputManagerService`
```
mSystemServiceManager.startService(TvInputManagerService.class);
```

90.`MediaResourceMonitorService`
```
mSystemServiceManager.startService(MediaResourceMonitorService.class);
```

91.`TvRemoteService`
```
mSystemServiceManager.startService(TvRemoteService.class);
```

92.`MediaRouterService`
```
mediaRouter = new MediaRouterService(context);
ServiceManager.addService(Context.MEDIA_ROUTER_SERVICE, mediaRouter);
```

93.`TrustManagerService`
```
mSystemServiceManager.startService(TrustManagerService.class);
```

94.`FingerprintService`
```
mSystemServiceManager.startService(FingerprintService.class);
```

95.`BackgroundDexOptService`
```
BackgroundDexOptService.schedule(context);
```

96.`ShortcutService`
```
// LauncherAppsService uses ShortcutService.
mSystemServiceManager.startService(ShortcutService.Lifecycle.class);
```

97.`LauncherAppsService`
```
mSystemServiceManager.startService(LauncherAppsService.class);
```

98.`MediaProjectionManagerService`
```
mSystemServiceManager.startService(MediaProjectionManagerService.class);
```

99.`WEAR_BLUETOOTH_SERVICE`
```
mSystemServiceManager.startService(WEAR_BLUETOOTH_SERVICE_CLASS);
```

100.`WEAR_WIFI_MEDIATOR_SERVICE`
```
mSystemServiceManager.startService(WEAR_WIFI_MEDIATOR_SERVICE_CLASS);
```

101.`WEAR_CELLULAR_MEDIATOR_SERVICE`
```
mSystemServiceManager.startService(WEAR_CELLULAR_MEDIATOR_SERVICE_CLASS);
```

102.`WEAR_TIME_SERVICE`
```
mSystemServiceManager.startService(WEAR_TIME_SERVICE_CLASS);
```

103.`MmsServiceBroker`
```
// MMS service broker
mmsService = mSystemServiceManager.startService(MmsServiceBroker.class);
```

104.`RetailDemoModeService`
```
mSystemServiceManager.startService(RetailDemoModeService.class);
```


# slf4j

The motivation for the SLF4J Android project was to ease using existing libraries which use SLF4J as their logging framework on the Google Android platform.

This project is basically a (i) repackaging of the SLF4J API part, together with (ii) a very lightweight binding implementation that simply forwards all SLF4J log requests to the logger provided on the Google Android platform. The API part is compiled from the same code basis of the standard distribution. This is the reason why we decided to keep the version numbering in sync with the standard SLF4J releases in order to reflect the code basis from which it was built.

## Usage
Assuming that you use the Eclipse ADT plugin, simply add slf4j-android-<version>.jar to your project classpath.
There is no further configuration required.
Use loggers as usual:
Declare a logger 
```
private static final Logger logger = LoggerFactory.getLogger(MyClass.class);
```
Invoke logging methods, e.g., 
```
logger.debug("Some log message. Details: {}", someObject.toString());
```
## Log level mapping
The following table shows the mapping from SLF4J log levels to log levels in the Android platform, implemented by the logger binding.
```
SLF4J	Android
----------------
TRACE	VERBOSE
DEBUG	DEBUG
INFO	INFO
WARN	WARN
ERROR	ERROR
```


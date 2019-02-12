# Window兼容ID

## Microsoft OS String Descriptor

The first time an USB device is plugged in, the Microsoft USB port driver (`usbport.sys`) will read the standard USB Descriptors, which includes the standard Device, Configuration, Interface and String Descriptors. It will also attempt to read an additional String Descriptor located at index `0xEE`. This String Descriptor, which is not mandated by the USB specifications, is a Microsoft extension called an OS String Descriptor and it is the first of two elements that establish whether a device is WCID.

When Windows checks for this String Descriptor, one of the first thing that happens is the creation of a registry entry, under`HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\usbflags`, that is the concatenation of `VID+PID+BCD_RELEASE_NUMBER` (**This key is never deleted**). In this new section, a `osvc` 2 byte `REG_BINARY` key is created, that indicates whether an appropriate Microsoft OS String Descriptor was found.

    If the `0xEE` descriptor doesn't exist, or if it doesn't match the Microsoft signature (`"MSFT100"`), then `osvc` is set to `0x0000`, which, in our case, means that the device is not WCID.
    If the `0xEE` descriptor exists and matches the expected organization of a Microsoft OS String Descriptor, then `oscv` is set to `0x01##`, where `01` indicates that the device satisfies the MS OS Vendor extensions and where `##` is the Vendor Code byte value (see below).

The following table details how exactly the Microsoft OS String Descriptor should be set in your firmware:

|Value      |Type       |Description|
|:-:        |:-:        |:--        |
|`0x12`     |BYTE       |Descriptor length (18 bytes)   |
|`0x03`     |BYTE       |Descriptor type (3 = String)   |
|`0x4D`, `0x00`, `0x53`, `0x00`,`0x46`, `0x00`, `0x54`, `0x00`,`0x31`, `0x00`, `0x30`, `0x00`,`0x30`, `0x00`  |7 WORDS Unicode String (LE)  |Signature: "MSFT100"   |
|`0x##`     |BYTE       |Vendor Code|
|`0x00`     |BYTE       |Padding    |



# nio

## buffer
```
0 <= mark <= length <= limit <= capacity
```

## select
```
while(true) {
    selector.select();
    Iterator<SelectionKey> iterator = selector.selectedKeys().iterator();
    while (iterator.hasNext()) {
        SelectionKey key = iterator.next();
        iterator.remove();  // Why remove it?
        process(key);
    } 
}
```

## process
```
void process(SelectionKey event) {
    buffer.clear();
    String eventType = (String) event.attachment();
    if (event.isReadable()) {
        SocketChannel channel = (SocketChannel) event.channel();
        channel.read(buffer);

        byte[] ret = new byte[buffer.flip().limit()];
        buffer.get(ret).clear();
        Log.d(TAG, new String(ret));
        Log.d(TAG, Arrays.toString(ret));

        if (Arrays.equals(ret, Command.CONNECTACCEPT)) {
            Log.d(TAG, "connect success.");
        } else if (new String(ret).contains(new String(Command.DISCIN))) {
                Log.d(TAG, "disk in");
        }
    }
}
```

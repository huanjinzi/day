# Java

## 修改Final字段
```

      Field final_o =  surfaceViewCls.getDeclaredField("modifiers");
      final_o.setAccessible(true);
      try {
            final_o.setInt(field,field.getModifiers() & ~Modifier.FINAL);
          } catch (IllegalAccessException e) {
            e.printStackTrace();
       }
```

## try-catch-finally
```
try {
            throwEx();
            return;
        } catch (Exception e) {
            System.out.println("catch");
            return;
        } finally {
            System.out.println("finally");
        }
        
        // System.out.println("method"); 编译错误
```
在`catch`代码段的最后添加`return`语句，`finally`的语句能够执行

## java7增强try语句
* 被自动关闭的资源必须实现`Closeable`或`AutoCloseable`接口。（`Closeable`是`AutoCloseable`的子接口，`Closeeable`接口里的`close()`方法声明抛出了`IOException`;`AutoCloseable`接口里的`close()`方法声明抛出了`Exception`）
* 被关闭的资源必须放在`try`语句后的圆括号中声明、初始化。如果程序有需要自动关闭资源的try语句后可以带多个catch块和一个finally块。
```

```

## Java String split
```
//Cation:head match and end match
String s = "/xx";
String[] ret = s.split("[/]");
// ret = ["","xx"]
```

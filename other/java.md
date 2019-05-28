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

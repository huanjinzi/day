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

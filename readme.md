#MyWhatsapp

*Proper Readme coming soon*


---
Android Regex for breaking apart chats:

```(?P<datetime>\d{1,2}\/\d{1,2}\/\d{2}, \d{1,2}:\d{1,2} ((?i)[ap]m))(?: - )(?P<name>.*?)(?::) (?P<message>.+)```


IOs: 
```(?P<datetime>\d{4}\-\d{2}\-\d{2}, \d{1,2}:\d{1,2}:\d{2} ((?i)[ap]m))(?:\]) (?P<name>.*?)(?::) (?P<message>.+)```




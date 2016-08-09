You use when you want to easily send a log.

# Usage

```
type ServiceLog struct {
 UID  string `json:"uid"`
 Action string `json:"action"`
}
serviceLog := &ServiceLog{UID: "test", Acton: "send log"}

t := tdlog.NewTDLog(endpoint, apikey)
err := t.Send(serviceLog)
```

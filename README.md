# go-pubsub
  
![Test](https://github.com/stefanoschrs/go-pubsub/workflows/Test/badge.svg)  
  
### Example
```
// Sub
s, err := Create("id")
if err != nil {
  log.Fatal(err)
}

data := s.Sub()

// Pub
s, err := Get("id")
if err != nil {
  log.Fatal(err)
}

s.Pub([]byte("Hello World"))
s.Close()
```

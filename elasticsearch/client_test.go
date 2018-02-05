package elasticsearch

var data = `
{
  "index": {
    "_index": "zipkin:span-2018-02-05",
    "_type": "span"
  }
}
{
  "timestamp_millis": 1517811812494,
  "_q": [
    "http.path",
    "http.path=/a/b/c.jpg"
  ],
  "traceId": "abcdefghijklmnopq",
  "id": "abcdefghijklmnopq",
  "kind": "SERVER",
  "name": "get",
  "timestamp": 1517811812494100,
  "duration": 30143,
  "localEndpoint": {
    "serviceName": "openzipkin-example",
    "ipv4": "127.0.0.1"
  },
  "remoteEndpoint": {
    "ipv4": "127.0.0.1",
    "port": 59513
  },
  "tags": {
    "http.path=/a/b/c.jpg"
  }
}
`

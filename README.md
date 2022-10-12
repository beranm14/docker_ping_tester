# Docker ping tester


Have this ever happen to you?

You needed to see if your service running on Cloud Run, Heroku, ... uses external IP address you allocated at the cloud.

You would appriciate a small app which could call a remote server where you can check in the logs what was the call IP address.

Look no more, that's the purpose of this thingy.

## Variables

 * `REMOTE_URL` - url you would like to call, default is `https://example.com`
 * `PORT` - port on which should be this service running

## Example on localhost

1) Run `go build main.go` and `./main`, runtime is mostly silent.
1) Check with `curl`:
```
curl http://localhost:3000/
Hello 🐱 from "/"
Called https://example.com, got 1256 bytes
```

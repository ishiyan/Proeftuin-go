# Web Authentication, Encryption, JWT, HMAC, & OAuth With Go

## Git version tagging

```bash
git add -S
git commit -m "message"
git push

git tag 0.1.0
git push --tags
```

## Git stash

```bash
git stash
git stash drop
```

## Using Curl in JSON decoding examples

Use [curlbuild.com](https://curlbuilder.com/) to build a command line.

```bash
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"}]' localhost:8080/decode
```

```bash
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"},{"First":"James"}]' localhost:8080/decode
```

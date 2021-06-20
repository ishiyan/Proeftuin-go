# Notes

On the server, decode this JSON into the appropriate data structure:

```json
[{"First":"Jenny"},{"First":"James"}]
```

```bash
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"}]' localhost:8080/decode
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"},{"First":"James"}]' localhost:8080/decode
```

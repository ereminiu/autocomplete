# Autocomplete System
REST API for autocomplete system aka "Top K(5) most searched queries" via Trie datastructure to get most popular queries. Server is running on port 6000 by default.

## Handlers
[POST] /search?q=kat&add=false

The "add" param is using only at the end of the search query to insert into database one.

Response:
```json
{
    "requests": [
        "kate",
        "katrina",
        "katie",
        "kathy",
        "kathleen"
    ]
}
```

[POST] /rebuild is using to rebuild trie with queries and their frequencies from database. 

Response:
```json
{
    "message": "Trie has been reinited from database"
}
```

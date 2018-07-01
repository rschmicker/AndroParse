#!/bin/bash
curl -X GET "http://localhost:9200/apks/_search?scroll=1m" -H 'Content-Type: application/json' -d'
{
    "query": {
        "match_all": {}
    }
}
' > dump.json

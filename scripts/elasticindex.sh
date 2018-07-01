#!/bin/bash
curl -X DELETE "http://localhost:9200/apks"
curl -X PUT "http://localhost:9200/apks"
curl -X PUT "http://localhost:9200/apks/_mapping/_doc" -H 'Content-Type: application/json' -d'
{
      "properties": {
        "Apis": {
          "type": "text"
        },
        "Date": {
          "type": "date",
          "format": "date_hour_minute_second"
        },
        "FileSize": {
          "type": "integer"
        },
        "Intents": {
          "type": "text"
        },
        "Malicious": {
          "type": "text"
        },
        "Md5": {
          "type": "text"
        },
        "PackageName": {
          "type": "text"
        },
        "PackageVersion": {
          "type": "text"
        },
        "Sha1": {
          "type": "text"
        },
        "Sha256": {
          "type": "text"
        },
        "Strings": {
          "type": "text"
        },
        "Permissions": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
}
'
go run bulkimport.go
time go run query_elastic.go

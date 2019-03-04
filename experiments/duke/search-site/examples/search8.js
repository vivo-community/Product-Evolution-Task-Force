import { client } from './config'
import { inspectAgg, error } from './search'

//https://stackoverflow.com/questions/48784733/aggregation-on-keyword-field-fails

// facets=aggregation, bucketing etc...
client.search({
  //size: "0",
  index: ["people"],
  //index: "people",
  //type: 'person',
  body: {
    query: {
      "query_string": { query: "*" }
    },
    aggs: {
      "types" : { "terms": {"field" : "type.label" }},
      "keywords": {
         "nested": {
             "path": "keywordList"
         },
         "aggs": {
           "keyword" : { "terms" : { "field": "keywordList.label.keyword" } }
         }
      },
      "affliliations": {
         "nested": {
             "path": "affiliationList"
         },
         "aggs": {
           //"title": { "terms" : { "field": "affiliationList.label.title" } },
           "department": { "terms" : { "field": "affiliationList.organization.label.dept" } }
         }
      }
    }
  }
}).then(inspectAgg, error);




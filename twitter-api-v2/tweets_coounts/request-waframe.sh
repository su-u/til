curl -X GET --location "https://api.twitter.com/2/tweets/counts/recent?query=warframe&granularity=day" \
    -H "Authorization: Bearer ${TWITTER_BEARER_TOKEN}" \
    -H "Accept: application/json" | jq -r '.data[] | [.start, .end, .tweet_count] | @csv' > warframe.csv

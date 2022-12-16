curl -X GET --location "https://api.twitter.com/2/tweets/counts/recent?query=twitter" \
    -H "Authorization: Bearer ${TWITTER_BEARER_TOKEN}" \
    -H "Accept: application/json" | jq -r '.'

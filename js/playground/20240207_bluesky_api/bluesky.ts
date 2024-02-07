import { BskyAgent } from '@atproto/api'
require('dotenv').config();

export const main = async () => {

  const agent = new BskyAgent({ service: 'https://bsky.social' });

  console.log(process.env.BSKY_EMAIL);

  const res = await agent.login({
    identifier: process.env.BSKY_EMAIL,
    password: process.env.BSKY_PASSWORD,
  });
  console.log(res);

  const follows = await agent.getFollows({
    actor: 'su-u.bsky.social',
    limit: 100,
  });
  console.log(follows.data);

  const axios = require('axios');

  let config = {
    method: 'get',
    maxBodyLength: Infinity,
    url: 'https://bsky.social/xrpc/app.bsky.feed.searchPosts?q=socialdog',
    headers: {
      'Accept': 'application/json',
      'Authorization': 'Bearer eyJhbGciOiJFUzI1NksifQ.eyJzY29wZSI6ImNvbS5hdHByb3RvLmFjY2VzcyIsInN1YiI6ImRpZDpwbGM6aDJ5Mmh0aGIydzN1emplZ3J4NXIyajNrIiwiaWF0IjoxNzA3MzE0NDY1LCJleHAiOjE3MDczMjE2NjUsImF1ZCI6ImRpZDp3ZWI6Ymxld2l0LnVzLXdlc3QuaG9zdC5ic2t5Lm5ldHdvcmsifQ.k8kdIf6MoPiabofv4S5H4al3IyxJlHfJtmvSEqTs7NrZeWFjl4_TKNvGIwawu5hXQM_KwR6XKmMC3aPH7_LLiQ'
    }
  };

  axios(config)
    .then((response) => {
      console.log(JSON.stringify(response.data));
    })
    .catch((error) => {
      console.log(error);
    });

}
import { serve } from 'https://deno.land/std@0.157.0/http/server.ts';

const handler = async (req: Request): Promise<Response> => {
  const reqUrl = new URL(req.url);
  const num = reqUrl.searchParams.get('number');

  const url = `https://pokeapi.co/api/v2/pokemon/${num}/`;
  const res = await fetch(url, {
    headers: {
      acccept: 'application/json',
    },
  });

  let body = '404';
  if (res.status == 200) {
    const json = await res.json();
    body = `${json['id']}<br>${json['name']}`;
  }

  return new Response(body, {
    status: res.status,
    headers: {
      'content-type': 'text/html',
    },
  });
};

serve(handler);
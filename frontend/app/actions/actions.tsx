'use server';

import { revalidatePath } from 'next/cache';

export async function getResult(prevState: any, formdata: FormData) {
  // console.log(formdata);
  const data = { 'day': parseInt(formdata.get('day') as string, 10), 'input': formdata.get('textInput') };
  const baseUrl = process.env.API_URL;
  const response = await fetch(`${baseUrl}`, {
    cache: 'no-store',
    method: 'POST',
    mode: 'cors',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  if (response.ok) {
    // const res = await response.text();
    // console.log(res);
    const { p1, p2 } = await response.json();
    return {
      p1: p1,
      p2: p2,
      //   message: res,
    };
  } else {
    console.error('Error: response not Ok', response.status, response.statusText);
  }
}

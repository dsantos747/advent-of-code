import Image from 'next/image';

const fetchData = async () => {
  const baseUrl = process.env.API_URL;
  const response = await fetch(`${baseUrl}/api/hello`);
  if (response.ok) {
    const data = await response.text();
    return data;
  } else {
    console.error('Error: response not Ok', response.status, response.statusText);
    return `ERROR - Code ${response.status}`;
  }
};

export default function Home() {
  // const test = fetchData();
  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      <div className='z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex'>{/* <div>{test}</div> */}</div>
    </main>
  );
}

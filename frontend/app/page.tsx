import Grid from './components/colourGrid';

const fetchData = async () => {
  // const baseUrl = process.env.API_URL;
  const response = await fetch(`https://europe-west9-advent-of-code-408219.cloudfunctions.net/hello-world-test`, {
    method: 'POST',
  });
  // console.log(baseUrl);

  if (response.ok) {
    const data = await response.text();
    return data;
  } else {
    console.error('Error: response not Ok', response.status, response.statusText);
    return `ERROR - Code ${response.status}`;
  }
};

export default function Home() {
  const submit = async (formdata: FormData) => {
    'use server';

    const baseUrl = process.env.API_URL;
    const response = await fetch(`${baseUrl}/api/hello`);
    if (response.ok) {
      const data = await response.text(); ///////// But now how do I update the value on the screen with this?
      ///////// Is it even possible to do this in a server form?
      // return data;
    } else {
      console.error('Error: response not Ok', response.status, response.statusText);
      // return `ERROR - Code ${response.status}`;
    }
  };

  const testVal = fetchData();

  const dayCount = 17;

  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      <div className='z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex'>
        <div className='text-left'>
          <form id='challengeForm' action={submit} className='flex flex-col gap-2'>
            <div>
              <label>Choose Challenge Day</label>
              <div className='-z-10 m-1'>
                <Grid length={25}></Grid>
              </div>
              {/* <div className='grid grid-cols-5 rounded-lg'>
                {Array.from({ length: 25 }, (_, i) => i + 1).map((item, index) => {
                  return (
                    <div key={index} className='bg-blue-600 aspect-square' hidden={item > dayCount ? true : false}>
                      <input id={`radio_${item}`} type='radio' name='day' value={item} className='hidden peer'></input>
                      <label
                        htmlFor={`radio_${item}`}
                        className='flex justify-center items-center m-0.5 px-2 aspect-square rounded-sm bg-black peer-checked:bg-[radial-gradient(rgb(243,227,124)_70%,rgb(243,211,74)_100%)] border border-gray-400'>
                        {item}
                      </label>
                    </div>
                  );
                })}
              </div> */}
            </div>
            <div className='flex flex-col'>
              <label htmlFor='textInput'>Paste your input below:</label>
              <textarea id='textInput' name='textInput'></textarea>
            </div>
            <button type='submit' className='mx-auto flex justify-center items-center p-2 border rounded-sm border-amber-400'>
              Generate Answer
            </button>
          </form>
          <div>
            <div>Answers:</div>
            {/* <div>{testVal}</div> */}
          </div>
        </div>
      </div>
    </main>
  );
}

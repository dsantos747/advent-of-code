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
    <main className='container'>
      <div className='content splash-content'>
        <div className=''>
          <form id='challengeForm' action={submit} className=''>
            <div id='challengeDay'>
              <label>Choose Challenge Day</label>
              <div className=''>
                <Grid length={25}></Grid>
              </div>
            </div>
            <div id='challengeText' className='flex flex-col'>
              <label htmlFor='textInput'>Paste your input below:</label>
              <textarea id='textInput' name='textInput'></textarea>
            </div>
            <button type='submit' className='mx-auto flex justify-center items-center p-2 border rounded-sm border-amber-400'>
              Generate Answer
            </button>
          </form>
        </div>
        <div>
          <div>Fix mouseglow carrying on when mouse leaves grid area</div>
          <div>Answers:</div>
          <div id='answer'>
            {/* {testVal} */}
            Response will go here
          </div>
        </div>
      </div>
    </main>
  );
}

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

    const data = { 'day': 1, 'input': formdata.get('textInput') };
    const baseUrl = 'http://localhost:8080/';
    const response = await fetch(`${baseUrl}`, {
      cache: 'no-store',
      method: 'POST',
      mode: 'cors',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    if (response.ok) {
      const res = await response.text();
      console.log(res);
    } else {
      console.error('Error: response not Ok', response.status, response.statusText);
    }
  };

  const testVal = fetchData();

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
            <button type='submit' className=''>
              Generate Answer
            </button>
          </form>
        </div>
        <div>
          <div>Answers:</div>
          <div id='answer'>
            {/* {testVal} */}
            Response will go here
          </div>
          <div>
            <br />
            <br />
            <br />
            <br />
            Fix mouseglow carrying on when mouse leaves grid area<br></br>
            <br></br>Maybe hide it behind all content on the screen, except background?
          </div>
        </div>
      </div>
    </main>
  );
}

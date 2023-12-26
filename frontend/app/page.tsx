import DayForm from './components/form';

const initialState = {
  p1: null,
  p2: null,
};

export default function Home() {
  return (
    <main className='container'>
      {/* <div className='content splash-content'> */}
      <DayForm></DayForm>
      {/* </div> */}
    </main>
  );
}

'use client';
import { useEffect, useState } from 'react';
import { useFormState, useFormStatus } from 'react-dom';
import { getResult } from '../actions/actions';

const debounce = (func: Function, delay: number) => {
  let timeoutId: number; //ReturnType<typeof setTimeout>;
  return function () {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(func, delay);
  };
};

type Props = {};

const initialState = {
  p1: null,
  p2: null,
};

function DayForm() {
  const [mousePos, setMousePos] = useState<{ x: number; y: number }>({ x: 0, y: 0 });
  const [mouseOver, setMouseOver] = useState<boolean>(false);
  const { pending } = useFormStatus();
  const [formState, formAction] = useFormState(getResult, initialState);

  const handleMouseMove: any = (event: React.MouseEvent<Window, MouseEvent>) => {
    const { clientX, clientY } = event;
    setMousePos({ x: clientX, y: clientY });
  };

  const length = 25;

  return (
    <div className='content splash-content'>
      <div className=''>
        <form id='challengeForm' action={formAction} className=''>
          <div id='challengeDay'>
            <label>Choose Challenge Day:</label>
            <div className=''>
              <div
                className='colourGrid'
                onMouseEnter={() => setMouseOver(true)}
                onMouseLeave={() => setMouseOver(false)}
                onMouseMove={handleMouseMove}>
                {mouseOver && <div style={{ top: mousePos.y - 16, left: mousePos.x - 16 }} className='mouseGlow'></div>}
                {Array.from({ length }, (_, i) => i + 1).map((item, index) => {
                  return (
                    <div key={index} className='colourTile' hidden={item > length ? true : false}>
                      <input id={`radio_${item}`} type='radio' name='day' value={item} className=''></input>
                      <label htmlFor={`radio_${item}`} className=''>
                        {item}
                      </label>
                    </div>
                  );
                })}
              </div>
            </div>
          </div>
          <div id='challengeText' className='flex flex-col'>
            <label htmlFor='textInput'>Paste your input below:</label>
            <textarea id='textInput' name='textInput'></textarea>
          </div>
          <button type='submit' aria-disabled={pending}>
            Generate Answer
          </button>
        </form>
      </div>
      <div>
        <div>Answers:</div>
        {/* <div> */}
        <div className='answer'>
          <p>Part 1:</p>
          <div className='answerField'>{formState?.p1}</div>
        </div>
        <div className='answer'>
          <p>Part 2:</p>
          <div className='answerField'>{formState?.p2}</div>
        </div>
        {/* </div> */}
      </div>
    </div>
  );
}

export default DayForm;
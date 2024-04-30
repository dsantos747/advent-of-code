'use client';
import { useCallback, useEffect, useState } from 'react';
import { useFormState, useFormStatus } from 'react-dom';
import { getResult } from '../actions/actions';

interface Year {
  Year: string;
  Locked: boolean;
  UnlockedDays: number;
}

const Years: Year[] = [
  { Year: '2015', Locked: false, UnlockedDays: 1 },
  { Year: '2016', Locked: true, UnlockedDays: 0 },
  { Year: '2017', Locked: true, UnlockedDays: 0 },
  { Year: '2018', Locked: true, UnlockedDays: 0 },
  { Year: '2019', Locked: true, UnlockedDays: 0 },
  { Year: '2020', Locked: true, UnlockedDays: 0 },
  { Year: '2021', Locked: true, UnlockedDays: 0 },
  { Year: '2022', Locked: true, UnlockedDays: 0 },
  { Year: '2023', Locked: false, UnlockedDays: 25 },
];

const debounce = (func: Function, delay: number) => {
  let timeoutId: number; //ReturnType<typeof setTimeout>;
  return function () {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(func, delay);
  };
};

const initialState = {
  p1: null,
  p2: null,
};

function DayForm() {
  const [mousePos, setMousePos] = useState<{ x: number; y: number }>({ x: 0, y: 0 });
  const [mouseOver, setMouseOver] = useState<boolean>(false);
  const { pending } = useFormStatus();
  const [formState, formAction] = useFormState(getResult, initialState);
  const [rect, setRect] = useState<DOMRect | null>(null);
  const [year, setYear] = useState<Year>(Years[Years.length - 1]);

  useEffect(() => {
    const grid = document.getElementsByClassName('colourGrid')[0];
    if (grid) {
      const newRect = grid.getBoundingClientRect();
      setRect(newRect);
    }
  }, []);

  const handleMouseMove: any = (event: React.MouseEvent<Window, MouseEvent>) => {
    const { clientX, clientY } = event;
    setMousePos({ x: clientX, y: clientY });
  };

  const handleYearDropdown = useCallback(
    (e: React.ChangeEvent<HTMLSelectElement>) => {
      const selectedYear = Years.find((s) => s.Year === e.target.value);
      if (selectedYear) {
        setYear(selectedYear);
      }
    },
    [year]
  );

  const maxDays = 25;

  return (
    <div className='content splash-content'>
      <div className=''>
        <form id='challengeForm' action={formAction} className=''>
          <div id='challengeYear'>
            <label htmlFor='year'>Year:</label>
            <select id='year' name='year' value={year.Year} onChange={handleYearDropdown}>
              {Years.map((y) => {
                return (
                  <option key={y.Year} disabled={y.Locked}>
                    {y.Year}
                  </option>
                );
              })}
            </select>
          </div>
          <div id='challengeDay'>
            <label>Choose Challenge Day:</label>
            <div className=''>
              <div
                className='colourGrid'
                onMouseEnter={() => setMouseOver(true)}
                onMouseLeave={() => setMouseOver(false)}
                onMouseMove={handleMouseMove}>
                {mouseOver && (
                  <div
                    style={{
                      top: (rect?.y ?? 0) + (rect?.height ?? 0) / 2 + 0.5 * (mousePos.y - ((rect?.y ?? 0) + (rect?.height ?? 0) / 2)) - 32,
                      left: (rect?.x ?? 0) + (rect?.width ?? 0) / 2 + 0.5 * (mousePos.x - ((rect?.x ?? 0) + (rect?.width ?? 0) / 2)) - 32,
                    }}
                    className='mouseGlow'></div>
                )}
                {Array.from(Array(maxDays).keys()).map((item: number) => {
                  item += 1;
                  return (
                    <div key={item} className='colourTile' hidden={item > maxDays}>
                      <input
                        id={`radio_${item}`}
                        type='radio'
                        name='day'
                        disabled={item > year.UnlockedDays}
                        value={item}
                        className=''></input>
                      <label htmlFor={`radio_${item}`} className=''>
                        {item}
                      </label>
                    </div>
                  );
                })}
              </div>
            </div>
          </div>
          <div id='challengeText'>
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

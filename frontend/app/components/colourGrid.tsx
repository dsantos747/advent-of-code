'use client';
import { useEffect, useState } from 'react';

const debounce = (func: Function, delay: number) => {
  let timeoutId: number; //ReturnType<typeof setTimeout>;
  return function () {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(func, delay);
  };
};

type Props = {
  length: number;
};

function Grid({ length }: Props) {
  const [mousePos, setMousePos] = useState<{ x: number; y: number }>({ x: 0, y: 0 });
  const [mouseOver, setMouseOver] = useState<boolean>(false);

  const handleMouseMove: any = (event: React.MouseEvent<Window, MouseEvent>) => {
    const { clientX, clientY } = event;
    setMousePos({ x: clientX, y: clientY });
  };

  return (
    <div
      className='-z-40 grid grid-cols-5 rounded-lg'
      onMouseEnter={() => setMouseOver(true)}
      onMouseLeave={() => setMouseOver(false)}
      onMouseMove={handleMouseMove}>
      {mouseOver && (
        <div
          style={{ top: mousePos.y - 10, left: mousePos.x - 10 }}
          className='transition-all absolute -z-40 w-6 h-6 bg-green-500 rounded-full blur-md'></div>
      )}
      {Array.from({ length }, (_, i) => i + 1).map((item, index) => {
        return (
          <div key={index} className='z-10 aspect-square' hidden={item > length ? true : false}>
            <input id={`radio_${item}`} type='radio' name='day' value={item} className='hidden peer'></input>
            <label
              htmlFor={`radio_${item}`}
              className='flex justify-center items-center m-0.5 px-2 aspect-square rounded-sm bg-black peer-checked:bg-[radial-gradient(rgb(243,227,124)_70%,rgb(243,211,74)_100%)] border border-gray-400'>
              {item}
            </label>
          </div>
        );
      })}
    </div>
  );
}

export default Grid;

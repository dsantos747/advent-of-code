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
  );
}

export default Grid;

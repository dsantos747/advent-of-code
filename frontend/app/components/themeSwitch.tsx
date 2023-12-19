'use client';

import { useTheme } from 'next-themes';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMoon, faSun, faLightbulb } from '@fortawesome/free-solid-svg-icons';

type Props = {};

function ThemeSwitch({}: Props) {
  const { theme, setTheme } = useTheme();

  return (
    <div className=''>
      <button onClick={() => (theme === 'dark' ? setTheme('light') : setTheme('dark'))}>
        <FontAwesomeIcon icon={faLightbulb} className='' />
      </button>
    </div>
  );
}

export default ThemeSwitch;

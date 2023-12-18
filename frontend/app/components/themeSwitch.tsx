import { Context, createContext, useState, Dispatch, SetStateAction } from 'react';
import { ReactNode } from 'react';

const getInitialTheme = () => {
  const prevColourPref = window.localStorage.getItem('color-mode');
  const hasPrevPref = typeof prevColourPref === 'string';

  if (hasPrevPref) {
    return prevColourPref;
  }

  const mqPref = window.matchMedia('(prefers-color-scheme: dark)');
  const hasMediaQueryPrefeence = typeof mqPref.matches === 'boolean';

  if (hasMediaQueryPrefeence) {
    return mqPref.matches ? 'dark' : 'light';
  }

  // Default theme
  return 'light';
};

interface Theme {
  // Add properties according to your theme structure
  primaryColor: string;
  secondaryColor: string;
  // ... other properties
}

interface ThemeContextValue {
  theme: Theme;
  setTheme: Dispatch<SetStateAction<Theme>>;
}

interface ThemeProviderProps {
  children: ReactNode;
}

export const ThemeContext = createContext<ThemeContextValue | undefined>(undefined);

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const [theme, setThemeState] = useState(getInitialTheme);

  const setTheme = (value) => {
    setThemeState(value);

    window.localStorage.setItem('color-mode', value);

    return <ThemeContext.Provider value={{ theme, setThemeState }}>{children}</ThemeContext.Provider>;
  };
};

type Props = {};

function ThemeSwitch({}: Props) {
  return <div className=''>Theme switch</div>;
}

export default ThemeSwitch;

import ThemeSwitch from './themeSwitch';

type Props = {};

function Header({}: Props) {
  return (
    <div id='header' className=''>
      <div className='content'>
        <h1>Advent of code 2023</h1>
        <ThemeSwitch></ThemeSwitch>
      </div>
    </div>
  );
}

export default Header;

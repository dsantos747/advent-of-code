import ThemeSwitch from './themeSwitch';

type Props = {};

function Header({}: Props) {
  return (
    <div id='header' className=''>
      <div className='content'>
        <div>
          <h1>Advent of Code 2023</h1>
          <p>Solutions Calculator</p>
        </div>

        <ThemeSwitch></ThemeSwitch>
      </div>
    </div>
  );
}

export default Header;

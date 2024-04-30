import ThemeSwitch from './themeSwitch';

function Header() {
  return (
    <div id='header' className=''>
      <div className='content'>
        <div>
          <h1>Advent of Code</h1>
          <p>Solutions Calculator</p>
        </div>

        <ThemeSwitch></ThemeSwitch>
      </div>
    </div>
  );
}

export default Header;

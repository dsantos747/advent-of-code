import ThemeSwitch from './themeSwitch';

type Props = {};

function Header({}: Props) {
  return (
    <div id='header' className=''>
      <div className='content'>
        <div>Advent of code 2023</div>
        <ThemeSwitch></ThemeSwitch>
      </div>
    </div>
  );
}

export default Header;

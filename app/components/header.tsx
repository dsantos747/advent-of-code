type Props = {};

function Header({}: Props) {
  return (
    <div className='w-screen fixed flex justify-center items-center py-6 px-12 md:px-20 shadow-lg z-50 bg-gray-700'>
      <div className='flex grow justify-between items-center max-w-6xl'>
        <div>Advent of code 2023</div>
      </div>
    </div>
  );
}

export default Header;

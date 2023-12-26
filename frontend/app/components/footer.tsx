type Props = {};

function Footer({}: Props) {
  return (
    <div id='footer' className=''>
      <div className='content'>
        <p>
          Created by Daniel Santos |{' '}
          <a href='https://danielsantosdev.vercel.app/' target='_blank'>
            Website
          </a>{' '}
          |{' '}
          <a href='https://github.com/dsantos747' target='_blank'>
            Github
          </a>
        </p>
      </div>
    </div>
  );
}

export default Footer;

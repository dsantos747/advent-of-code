import type { Metadata } from 'next';
import { Red_Hat_Mono } from 'next/font/google';
import './styles.css';
import Header from './components/header';
import Footer from './components/footer';
import { ThemeProvider } from './components/themeProvider';

const fontMono = Red_Hat_Mono({ weight: '400', subsets: ['latin'], fallback: ['monospace', 'system-ui'] });

export const metadata: Metadata = {
  title: 'Advent of Code 2023 - Solutions Calculator',
  description:
    'A calculator for the programming challenges in Advent of Code 2023. All solutions were written in Go and are available on my Github',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang='en' suppressHydrationWarning>
      <body className={fontMono.className}>
        <ThemeProvider attribute='class' defaultTheme='system' enableSystem>
          <Header></Header>
          {children}
          <Footer></Footer>
        </ThemeProvider>
      </body>
    </html>
  );
}

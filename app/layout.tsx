import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';
import Header from './components/header';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Advent of Code 2023 - Solutions Calculator',
  description:
    'A calculator for the programming challenges in Advent of Code 2023. All solutions were written in Go and are available on my Github',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang='en'>
      <Header></Header>
      <body className={inter.className}>{children}</body>
    </html>
  );
}

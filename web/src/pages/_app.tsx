import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { TodoContextProvider } from 'context/TodoContext';
import { SessionProvider } from 'next-auth/react';

function MyApp({ Component, pageProps: { session, ...pageProps } }: AppProps) {
  return (
    <SessionProvider session={session}>
      <TodoContextProvider>
        <Component {...pageProps} />
      </TodoContextProvider>
    </SessionProvider>
  );
}

export default MyApp;

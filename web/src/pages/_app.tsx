import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { TodoContextProvider } from 'context/TodoContext';
import { SessionProvider } from 'next-auth/react';
import { UserContextProvider } from 'context/UserContext';

function MyApp({ Component, pageProps: { session, ...pageProps } }: AppProps) {
  return (
    <SessionProvider session={session}>
      <UserContextProvider>
        <TodoContextProvider>
          <Component {...pageProps} />
        </TodoContextProvider>
      </UserContextProvider>
    </SessionProvider>
  );
}

export default MyApp;

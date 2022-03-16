import type { NextPage } from 'next';
import Head from 'next/head';
import styles from 'styles/Home.module.css';
import { UserSetting } from 'component/UserSetting';

export default function Home() {
  return (
    <div>
      <UserSetting></UserSetting>
    </div>
  );
}

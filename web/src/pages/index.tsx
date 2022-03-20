import type { NextPage } from 'next';
import Head from 'next/head';
import styles from 'styles/Home.module.css';
import { Header } from 'component/Header';
import { Footer } from 'component/Footer';
import { Body } from 'component/Body';
import { ChangeEventHandler, useEffect, useRef, useState } from 'react';
import { addDays, isSameDay, startOfDay, subMinutes } from 'date-fns';
import { formatInTimeZone } from 'date-fns-tz';

type Items = {
  id: string;
  content: string;
  completed: boolean;
  execution_date: string | null;
}[];

// apiレスポンス 削除予定
const res: { items: Items } = {
  items: [
    {
      id: 'id001',
      content: '買い物する',
      completed: false,
      execution_date: '2019-10-06 00:00:00',
    },
    {
      id: 'id002',
      content: '洗濯物を干す',
      completed: true,
      execution_date: null,
    },
    {
      id: 'id003',
      content: '※常に明日になるタスク',
      completed: true,
      execution_date: formatInTimeZone(addDays(new Date(), 1), 'UTC', 'yyyy-MM-dd HH:mm:ss'),
      // テスト用 常に明日の日付になる
    },
  ],
};

type TodosList = {
  label: string;
  color: string;
  bg: string;
  todos: Items;
}[];

const initTodosList: TodosList = [
  {
    label: '今日する',
    color: 'text-rose',
    bg: 'bg-rose',
    todos: [],
  },
  {
    label: '明日する',
    color: 'text-orange',
    bg: 'bg-orange',
    todos: [],
  },
  {
    label: '今度する',
    color: 'text-yellow',
    bg: 'bg-yellow',
    todos: [],
  },
];

const Home: NextPage = () => {
  // input関連
  const ref = useRef<HTMLInputElement>(null);
  const [text, setText] = useState('');
  const input: ChangeEventHandler<HTMLInputElement> = (e) => setText(e.target.value);
  // APIのレスポンス
  const [data, setData] = useState<Items>(res.items);
  // フロント用に加工した配列
  const [todosList, setTodosList] = useState(initTodosList);

  /*
    日付周りのメモ
    execution_dateが明日か検証して出し分ける
    明日なら「明日やる」
    明日じゃないなら「今日やる」
    ※ 過去はすべて今日やるタスク、明日以降の未来は扱わない。
    nullは「今度やる」
    基本的には日本時刻は残さず条件分岐する際にのみ利用する
  */

  // 日付の形式
  const dateFormat = 'yyyy-MM-dd HH:mm:ss';
  // 明日か判別する、nullはそもそも入れさせない。
  const validationTommorow = (date: string): boolean => {
    const argDate = new Date(date);
    const formatArgDate = subMinutes(argDate, argDate.getTimezoneOffset());
    // 現在からみて明日を取得する
    const tomorrow = startOfDay(addDays(new Date(), 1));
    // 引数がtomorrowと同じ日付か検証する
    return isSameDay(formatArgDate, tomorrow);
  };

  // タスクを追加する
  const addTask = (label: string): boolean => {
    if (text === '') return false;

    const date = (): string | null => {
      switch (label) {
        case todosList[0].label:
          return formatInTimeZone(new Date(), 'UTC', dateFormat);
        case todosList[1].label:
          return formatInTimeZone(addDays(new Date(), 1), 'UTC', dateFormat);
        default:
          return null;
      }
    };

    setData((prevData) => [
      ...prevData,
      {
        id: String(Math.random()),
        content: text,
        completed: false,
        execution_date: date(),
      },
    ]);

    setText('');
    return false;
  };

  // タスクを削除する
  const removeTask = (id: string) => {
    setData((prevData) => [...prevData.filter((item) => item.id !== id)]);
  };

  // タスクを完了・未完了にする
  const completedTask = (id: string) => {
    setData((prevData) => [
      ...prevData.map((item) => {
        if (item.id !== id) return item;
        return {
          ...item,
          completed: !item.completed,
        };
      }),
    ]);
  };

  // タスクを編集する
  const [currentTask, setCurrentTask] = useState({ active: false, id: '' });
  const editTask = (id: string, content: string) => {
    setCurrentTask({ active: true, id: id });
    ref.current?.focus();
    setText(content);
  };

  // タスクを更新する
  const updateTask = (label: string) => {
    const date = (): string | null => {
      switch (label) {
        case todosList[0].label:
          return formatInTimeZone(new Date(), 'UTC', dateFormat);
        case todosList[1].label:
          return formatInTimeZone(addDays(new Date(), 1), 'UTC', dateFormat);
        default:
          return null;
      }
    };

    setData((prevData) => [
      ...prevData.filter((item) => item.id !== currentTask.id),
      ...prevData
        .filter((item) => item.id === currentTask.id)
        .map((item) => ({
          ...item,
          content: text,
          execution_date: date(),
        })),
    ]);

    setCurrentTask({ active: false, id: '' });
    setText('');
  };

  // タスクが増減するたびに配列を作成
  useEffect(() => {
    setTodosList((prevTodosList) => [
      ...prevTodosList.map((prevTodosItem) => ({
        ...prevTodosItem,
        todos: data.filter((todo) => {
          const date = todo.execution_date;
          switch (prevTodosItem.label) {
            case prevTodosList[0].label:
              return date === null ? false : !validationTommorow(date);
            case prevTodosList[1].label:
              return date === null ? false : validationTommorow(date);
            default:
              return date === null;
          }
        }),
      })),
    ]);
  }, [data]);

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <Header></Header>

      <Body>
        <div className='grid gap-8'>
          {todosList.map((todosItem, i) => (
            <section key={i}>
              <p className={todosItem.color}>{todosItem.label}</p>
              <ul>
                {todosItem.todos.map((todo, i) => (
                  <li className={currentTask.id === todo.id ? 'bg-[#FBBF24]/10' : ''} key={i}>
                    <input onChange={() => completedTask(todo.id)} checked={todo.completed} type='checkbox' />
                    <button onClick={() => editTask(todo.id, todo.content)}>{todo.content}</button>
                    <button onClick={() => removeTask(todo.id)}>削除</button>
                  </li>
                ))}
                <li className={todosItem.todos.length ? 'hidden md:block' : ''}>
                  <button onClick={() => ref.current?.focus()}>タスクを追加する</button>
                </li>
              </ul>
            </section>
          ))}
        </div>
      </Body>

      <Footer>
        <div className='grid gap-2 py-8'>
          <input className='border border-gray' value={text} onChange={input} ref={ref} type='text' />
          <div className='grid gap-1 grid-flow-col'>
            {todosList.map((todosItem, index) => (
              <button
                className={`${todosItem.bg} text-white`}
                onClick={() => {
                  currentTask.active ? updateTask(todosItem.label) : addTask(todosItem.label);
                }}
                key={index}
              >
                {todosItem.label}
              </button>
            ))}
          </div>
        </div>
      </Footer>
    </div>
  );
};

export default Home;

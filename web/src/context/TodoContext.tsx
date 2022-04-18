import { addDays, isSameDay, startOfDay, subMinutes } from 'date-fns';
import { formatInTimeZone } from 'date-fns-tz';
import React, { ChangeEventHandler, createContext, ReactNode, useContext, useRef, useState, VFC } from 'react';

type InitialState = {
  data: Items;
  setData: React.Dispatch<React.SetStateAction<Items>>;
  todosList: TodosList;
  setTodosList: React.Dispatch<React.SetStateAction<TodosList>>;
  handleTodosList: () => void;
  text: string;
  setText: React.Dispatch<React.SetStateAction<string>>;
  ref: React.RefObject<HTMLInputElement>;
  input: React.ChangeEventHandler<HTMLInputElement>;
  dateFormat: 'yyyy-MM-dd HH:mm:ss';
  validationTommorow: (date: string) => boolean;
  addTask: (label: string) => boolean;
  removeTask: (id: string) => void;
  completedTask: (id: string) => void;
  currentTask: string;
  setCurrentTask: React.Dispatch<React.SetStateAction<string>>;
  editTask: (id: string, content: string) => void;
  updateTask: (label: string) => void;
};

type Props = {
  children: ReactNode;
};

type Items = {
  id: string;
  content: string;
  completed: boolean;
  execution_date: {
    String: string;
    valid: boolean;
  };
}[];

type TodosList = {
  label: string;
  color: string;
  todos: Items;
}[];

const initTodosList: TodosList = [
  {
    label: '今日する',
    color: 'text-rose',
    todos: [],
  },
  {
    label: '明日する',
    color: 'text-orange',
    todos: [],
  },
  {
    label: '今度する',
    color: 'text-yellow',
    todos: [],
  },
];

export const TodoContext = createContext<InitialState | null>(null);
export const useTodoContext = () => {
  const context = useContext(TodoContext);
  if (context === null) {
    throw new Error('useCount must be used within a CountProvider');
  }
  return context;
};
export const TodoContextProvider: VFC<Props> = ({ children }) => {
  // APIのレスポンス
  const [data, setData] = useState<Items>([]);
  // フロント用に加工した配列
  const [todosList, setTodosList] = useState(initTodosList);
  // フロント用に加工する処理
  const handleTodosList = () => {
    setTodosList((prevTodosList) => [
      ...prevTodosList.map((prevTodosItem) => ({
        ...prevTodosItem,
        todos: data.filter((todo) => {
          const date = todo.execution_date.String;
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
  };

  // input関連
  const [text, setText] = useState('');
  const ref = useRef<HTMLInputElement>(null);
  const input: ChangeEventHandler<HTMLInputElement> = (e) => setText(e.target.value);

  // 日付の形式
  const dateFormat = 'yyyy-MM-dd HH:mm:ss';
  // 明日ならtrue、nullはそもそも入れさせない。
  const validationTommorow = (date: string): boolean => {
    /*
      iosだと年月日の-（ハイフン）が悪さをしている模様
      参考記事:https://qiita.com/pearmaster8293/items/b5b0df28147eb049f1ea
    */
    const argDate = new Date(date.replace(/-/g, '/'));
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

  // 選択したタスクの情報
  const [currentTask, setCurrentTask] = useState('');
  // 選択したタスクを編集する
  const editTask = (id: string, content: string) => {
    ref.current?.focus();
    setCurrentTask(id);
    setText(content);
  };

  // 選択したタスクを更新する
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
      ...prevData.filter((item) => item.id !== currentTask),
      ...prevData
        .filter((item) => item.id === currentTask)
        .map((item) => ({
          ...item,
          content: text,
          execution_date: {
            String: date(),
            valid: true,
          },
        })),
    ]);

    setCurrentTask('');
    setText('');
  };

  return (
    <TodoContext.Provider
      value={{
        data,
        setData,
        todosList,
        setTodosList,
        handleTodosList,
        text,
        setText,
        ref,
        input,
        dateFormat,
        validationTommorow,
        addTask,
        removeTask,
        completedTask,
        currentTask,
        setCurrentTask,
        editTask,
        updateTask,
      }}
    >
      {children}
    </TodoContext.Provider>
  );
};

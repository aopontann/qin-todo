import React from 'react';
import { Todo } from 'component/Todo';

interface Props {
  todosList: {
    title: string;
    category: string;
    todos: string[];
  }[];
}

export const TodosList: React.VFC<Props> = ({ todosList }) => {
  return (
    <div>
      {todosList.map((todosItem, index) => (
        <section className='my-4' key={index}>
          <h2>{todosItem.title}</h2>
          <ul>
            {todosItem.todos.map((todo, index) => (
              <li key={index}>
                <Todo todo={todo} />
              </li>
            ))}
            {/* タスクが一個以上で非表示にするクラスをつける PCは常に表示 */}
            <li className={todosItem.todos.length > 0 ? 'hidden md:block' : undefined}>
              <button>タスクを追加する</button>
            </li>
          </ul>
        </section>
      ))}
    </div>
  );
};

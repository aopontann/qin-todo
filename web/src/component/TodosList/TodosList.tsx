import React from 'react';
import { useTodoContext } from 'context/TodoContext';
import { Todo } from 'component/Todo';

export const TodosList = () => {
  const { todosList, ref } = useTodoContext();

  return (
    <div className='grid gap-8'>
      {todosList.map((todosItem, i) => (
        <section className={`${todosItem.color} grid gap-2`} key={i}>
          <p className={`text-h2 font-bold text-current`}>{todosItem.label}</p>
          <ul>
            {todosItem.todos.map((todo, i) => (
              <li key={i}>
                <Todo todo={todo} />
              </li>
            ))}
            <li className={todosItem.todos.length ? 'hidden md:block' : ''}>
              <button className='text-gray' onClick={() => ref.current?.focus()}>
                タスクを追加する
              </button>
            </li>
          </ul>
        </section>
      ))}
    </div>
  );
};

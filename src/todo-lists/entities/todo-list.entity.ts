import { TodoItem } from 'src/todo-items/entities/todo-item.entity';
import {
  Column,
  Entity,
  OneToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('todo_lists')
export class TodoList {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ length: 100 })
  name: string;

  @Column({ type: 'tinyint', default: 0 })
  status: number;

  @Column({
    type: 'datetime',
    name: 'created_time',
    default: () => 'CURRENT_TIMESTAMP',
  })
  createdTime: Date;

  @Column({
    type: 'datetime',
    name: 'updated_time',
    default: () => 'CURRENT_TIMESTAMP',
    onUpdate: 'CURRENT_TIMESTAMP',
  })
  updatedTime: Date;

  @Column({ type: 'datetime', nullable: true, name: 'deleted_time' })
  deletedTime: Date;

  @Column({ nullable: true, name: 'create_user' })
  createUser: number;

  @Column({ nullable: true, name: 'update_user' })
  updateUser: number;

  @Column({ nullable: true, name: 'delete_user' })
  deleteUser: number;

  @OneToOne(() => TodoItem, (item) => item.todoList)
  todoItem: TodoItem;
}

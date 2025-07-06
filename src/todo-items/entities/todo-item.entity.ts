import { TodoList } from 'src/todo-lists/entities/todo-list.entity';
import {
  Column,
  Entity,
  JoinColumn,
  OneToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('todo_items')
export class TodoItem {
  @PrimaryGeneratedColumn()
  id: number;

  @Column('text')
  content: string;

  @Column({ default: 0 })
  position: number;

  @Column({
    type: 'tinyint',
    default: 0,
    name: 'is_completed',
    comment: '0:未完成, 1:已完成',
  })
  isCompleted: number;

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

  @Column({ type: 'datetime', name: 'deleted_time', nullable: true })
  deletedTime: Date;

  @Column({ nullable: true, name: 'create_user' })
  createUser: number;

  @Column({ nullable: true, name: 'update_user' })
  updateUser: number;

  @Column({ nullable: true, name: 'delete_user' })
  deleteUser: number;

  @OneToOne(() => TodoList, (list) => list.todoItem, {
    cascade: true,
    onDelete: 'CASCADE',
    onUpdate: 'CASCADE',
  })
  @JoinColumn({
    name: 'todo_list_id',
  })
  todoList: TodoList;
}

import { UserAuth } from 'src/user-auths/entities/user-auth.entity';
import {
  Column,
  Entity,
  OneToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity('users')
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ length: 50 })
  username: string;

  @Column({ length: 200, nullable: true })
  avatar: string;

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

  @OneToOne(() => UserAuth, (userAuth) => userAuth.user)
  userAuth: UserAuth;
}

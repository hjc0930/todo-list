import { User } from 'src/users/entities/user.entity';
import {
  Column,
  Entity,
  JoinColumn,
  OneToOne,
  PrimaryGeneratedColumn,
  Unique,
} from 'typeorm';

@Entity('user_auths')
@Unique(['identifier'])
export class UserAuth {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ length: 20, name: 'identity_type' })
  identityType: string;

  @Column({ length: 100 })
  identifier: string;

  @Column({ length: 200 })
  credential: string;

  @Column({ type: 'tinyint', default: 0 })
  verified: number;

  @OneToOne(() => User, (user) => user.userAuth, {
    cascade: true,
    onDelete: 'CASCADE',
    onUpdate: 'CASCADE',
  })
  @JoinColumn({
    name: 'user_id',
  })
  user: User;

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

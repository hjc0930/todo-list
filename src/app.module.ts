import { Module, ValidationPipe } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UsersModule } from './users/users.module';
import { UserAuthsModule } from './user-auths/user-auths.module';
import { TodoListsModule } from './todo-lists/todo-lists.module';
import { TodoItemsModule } from './todo-items/todo-items.module';
import { ConfigModule } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { AuthModule } from './auth/auth.module';
import { APP_FILTER, APP_GUARD } from '@nestjs/core';
import { AuthGuard } from './common/auth-guard.guard';
import { GlobalExceptionFilter } from './common/global-exception.filter';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      envFilePath: ['.env', '.env.local'],
      cache: true,
    }),
    JwtModule.register({
      global: true,
      secret: 'cheng',
      signOptions: {
        expiresIn: '1800s',
      },
    }),
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'localhost',
      port: 3306,
      username: 'root',
      password: '123456',
      database: 'todo_list',
      entities: [__dirname + '/**/*.entity{.ts,.js}'],
      logging: true,
      synchronize: true,
    }),
    UsersModule,
    UserAuthsModule,
    TodoListsModule,
    TodoItemsModule,
    AuthModule,
  ],
  controllers: [AppController],
  providers: [AppService, {
    provide: APP_GUARD,
    useClass: AuthGuard
  }, {
      provide: APP_FILTER,
      useValue: new ValidationPipe({
        transform: true,        // 自动类型转换（如字符串转数字）
        whitelist: true,         // 过滤 DTO 未定义的属性
        forbidNonWhitelisted: true, // 禁止额外属性，返回 400 错误
      })
    }, {
      provide: APP_FILTER,
      useClass: GlobalExceptionFilter,
    }],
})
export class AppModule { }

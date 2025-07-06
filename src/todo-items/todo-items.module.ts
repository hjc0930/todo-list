import { Module } from '@nestjs/common';
import { TodoItemsService } from './todo-items.service';

@Module({
  providers: [TodoItemsService],
})
export class TodoItemsModule {}

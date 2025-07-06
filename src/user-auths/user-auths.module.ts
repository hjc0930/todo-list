import { Module } from '@nestjs/common';
import { UserAuthsService } from './user-auths.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserAuth } from './entities/user-auth.entity';

@Module({
  imports: [TypeOrmModule.forFeature([UserAuth])],
  providers: [UserAuthsService],
  exports: [UserAuthsService],
})
export class UserAuthsModule {}

import { Injectable } from '@nestjs/common';
import { CreateUserAuthDto } from './dto/create-user-auth.dto';
import { UpdateUserAuthDto } from './dto/update-user-auth.dto';
import { getConnection, Repository } from 'typeorm';
import { UserAuth } from './entities/user-auth.entity';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class UserAuthsService {
  constructor(
    @InjectRepository(UserAuth)
    private readonly userAuthRepository: Repository<UserAuth>,
  ) { }
  create(createUserAuthDto: CreateUserAuthDto) {
    return 'This action adds a new userAuth';
  }

  findAll() {
    return `This action returns all userAuths`;
  }

  findOne(id: number) {
    return this.userAuthRepository.findBy({
      id,
    });
  }

  findUserAuthByIdentity(identifier: string) {
    return this.userAuthRepository.findOneBy({ identifier });
  }

  update(id: number, updateUserAuthDto: UpdateUserAuthDto) {
    return `This action updates a #${id} userAuth`;
  }

  async remove(id: number) {
    await this.userAuthRepository.update(id, {
      deleteUser: 1,
      deletedTime: new Date(),
    })
    return `This action removes a #${id} userAuth`;
  }
}

import { Inject, Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { Repository } from 'typeorm';
import { User } from './entities/user.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { UserAuthsService } from 'src/user-auths/user-auths.service';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User)
    private readonly userRepository: Repository<User>,
    @Inject(UserAuthsService)
    private readonly userAuthsService: UserAuthsService,
  ) { }

  create = (createUserDto: CreateUserDto) => {
    return this.userRepository.create(createUserDto);
  };

  findAll() {
    return `This action returns all users`;
  }

  findOne(id: number) {
    return `This action returns a #${id} user`;
  }

  async findUserByIdentity(identity: string) {
    const userAuth =
      await this.userAuthsService.findUserAuthByIdentity(identity);
    if (!userAuth) {
      return null;
    }
    const userId = userAuth.user.id;
    const currentUser = await this.userRepository.findOne({
      where: { id: userId },
      relations: ['userAuth'],
    });

    return currentUser;
  }

  update(id: number, updateUserDto: UpdateUserDto) {
    return `This action updates a #${id} user`;
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}

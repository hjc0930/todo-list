import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { RegisterDto } from './dto/register.dto';
import { UsersService } from 'src/users/users.service';
import { IdentityTypeEnum } from 'src/common/todo-list-enum';
import * as bcrypt from 'bcrypt';
import { BusinessException } from 'src/common/business.expection';
import { BusinessStatusEnum } from 'src/common/business-status-enum';
import { LoginDto } from './dto/login.dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly jwtService: JwtService,
    private readonly userService: UsersService,
  ) { }
  private readonly SALT_ROUNDS = 10; // 加密强度，推荐值 10-12[1](@ref)

  public async register(register: RegisterDto) {
    const { identityType } = register;

    switch (identityType) {
      case IdentityTypeEnum.EMAIL: {
        await this.registerForEmail(register);
        break;
      }
      case IdentityTypeEnum.PHONE: {
        this.registerForPhone(register);
        break;
      }
      case IdentityTypeEnum.WEICHAT: {
        this.registerForWeiChat(register);
        break;
      }
    }
    return 'Successfully registered';
  };

  public async login(loginDto: LoginDto) {
    const { identityType } = loginDto;
    switch (identityType) {
      case IdentityTypeEnum.EMAIL: {
        await this.loginForEmail(loginDto);
        break;
      }
      case IdentityTypeEnum.PHONE: {
        await this.loginForPhone(loginDto);
        break;
      }
      case IdentityTypeEnum.WEICHAT: {
        await this.loginForWeiChat(loginDto);
        break;
      }
    }
    return 'Login successful';
  };

  private registerForEmail = async (register: RegisterDto) => {
    const currentUser = await this.userService.findUserByIdentity(
      register.identity,
    );

    if (currentUser) {
      throw new BusinessException(
        BusinessStatusEnum.PARAM_ERROR.getCode(),
        BusinessStatusEnum.PARAM_ERROR.getMsg(),
      );
    }
    const hashedPassword = await this.hashPassword(register.identity);

    this.userService.create({
      username: register.username,
      identityType: IdentityTypeEnum.EMAIL,
      identity: register.identity,
      credential: hashedPassword,
    });
  };

  private registerForPhone = (register: RegisterDto) => {
    // TODO
  };

  private registerForWeiChat = (register: RegisterDto) => {
    // TODO
  };

  private loginForEmail = async (loginDto: LoginDto) => {
    const currentUser = await this.userService.findUserByIdentity(
      loginDto.identity,
    );

    if (!currentUser) {
      throw new BusinessException(
        BusinessStatusEnum.Incorrect_USERNAME_OR_PASSWORD.getCode(),
        BusinessStatusEnum.Incorrect_USERNAME_OR_PASSWORD.getMsg(),
      );
    }
    const isPasswordValid = await this.validatePassword(
      loginDto.credential,
      currentUser.userAuth.credential,
    );

    if (!isPasswordValid) {
      throw new BusinessException(
        BusinessStatusEnum.Incorrect_USERNAME_OR_PASSWORD.getCode(),
        BusinessStatusEnum.Incorrect_USERNAME_OR_PASSWORD.getMsg(),
      );
    }
    const payload = {
      userId: currentUser.id,
      userName: currentUser.username,
    };
    const token = await this.jwtService.signAsync(payload);
    return token;
  };
  private loginForPhone = async (loginDto: LoginDto) => {
    // TODO
  };
  private loginForWeiChat = async (loginDto: LoginDto) => {
    // TODO
  };

  private hashPassword = (password: string) => {
    return bcrypt.hash(password, this.SALT_ROUNDS);
  };

  private validatePassword = (
    plainPassword: string,
    hashedPassword: string,
  ): Promise<boolean> => {
    return bcrypt.compare(plainPassword, hashedPassword); // 返回布尔值
  };
}

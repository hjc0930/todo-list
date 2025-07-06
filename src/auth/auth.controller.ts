import { Controller, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { Public } from 'src/common/public.decorator';
import { RegisterDto } from './dto/register.dto';
import { LoginDto } from './dto/login.dto';

@Public()
@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) { }

  @Post('register')
  public register(register: RegisterDto) {
    return this.authService.register(register);
  }

  @Post('login')
  public login(loginDto: LoginDto) {
    return this.authService.login(loginDto);
  }
}

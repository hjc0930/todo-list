import { IsString, IsNotEmpty } from 'class-validator';
import { IdentityTypeEnum } from 'src/common/todo-list-enum';

export class RegisterDto {
  @IsString()
  username: string;
  @IsNotEmpty()
  identityType: IdentityTypeEnum;
  @IsNotEmpty()
  identity: string;
  @IsNotEmpty()
  credential: string;
}

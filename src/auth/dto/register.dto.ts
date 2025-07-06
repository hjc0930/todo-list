import { IdentityTypeEnum } from 'src/common/todo-list-enum';

export class RegisterDto {
  username: string;
  identityType: IdentityTypeEnum;
  identity: string;
  credential: string;
}

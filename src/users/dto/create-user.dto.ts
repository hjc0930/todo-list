import { IdentityTypeEnum } from 'src/common/todo-list-enum';

export class CreateUserDto {
  username: string;
  identityType: IdentityTypeEnum;
  identity: string;
  credential: string;
}

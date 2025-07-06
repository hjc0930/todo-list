import { IdentityTypeEnum } from 'src/common/todo-list-enum';

export class LoginDto {
  identityType: IdentityTypeEnum; // Identity type, e.g., 'email', 'phone'
  identity: string; // User's identity, e.g., email or phone number
  credential: string; // User's password
}

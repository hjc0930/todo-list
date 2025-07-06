export class BusinessStatusEnum {
  static readonly SUCCESS = new BusinessStatusEnum(
    'SUCCESS',
    '20000',
    'Success',
  );
  static readonly PARAM_ERROR = new BusinessStatusEnum(
    'PARAM_ERROR',
    '40000',
    'Parameter Error',
  );
  static readonly UNAUTHORIZED = new BusinessStatusEnum(
    'UNAUTHORIZED',
    '40001',
    'Unauthorized',
  );

  static readonly NOT_FOUND = new BusinessStatusEnum(
    'NOT_FOUND',
    '40002',
    'Not Found',
  );

  static readonly SYSTEM_ERROR = new BusinessStatusEnum(
    'SYSTEM_ERROR',
    '50000',
    'System Error',
  );

  static readonly Incorrect_USERNAME_OR_PASSWORD = new BusinessStatusEnum(
    'Incorrect_USERNAME_OR_PASSWORD',
    '40003',
    'Incorrect username or password',
  );

  constructor(
    private readonly name: string,
    private readonly code: string,
    private readonly msg: string,
  ) {}

  public getName(): string {
    return this.name;
  }
  public getCode(): string {
    return this.code;
  }

  public getMsg(): string {
    return this.msg;
  }
}

import { HttpException, HttpStatus } from '@nestjs/common';

export class BusinessException extends HttpException {
  constructor(
    public readonly errorCode: string, // 业务错误码（如 10001）
    public readonly errorMessage: string, // 业务错误信息
    statusCode: HttpStatus = HttpStatus.PARTIAL_CONTENT, // HTTP 状态码（默认206）
  ) {
    super({ errorCode, errorMessage }, statusCode);
  }
}

import {
  ArgumentsHost,
  ExceptionFilter,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { Catch } from '@nestjs/common';
import { Request, Response } from 'express';
import { BusinessException } from './business.expection';
import { BusinessStatusEnum } from './business-status-enum';

@Catch()
export class GlobalExceptionFilter implements ExceptionFilter {
  catch(exception: unknown, host: ArgumentsHost) {
    const ctx = host.switchToHttp();

    const request = ctx.getRequest<Request>();
    const response = ctx.getResponse<Response>();

    let status: number, message: string, errorCode: string;

    if (exception instanceof BusinessException) {
      status = exception.getStatus();
      message = exception.message;
      errorCode = exception.errorCode;
    } else if (exception instanceof HttpException) {
      // 处理内置HTTP异常
      status = exception.getStatus();
      message = exception.message;
      errorCode = status.toString(); // 默认用HTTP状态码作为错误码
    } else {
      // 处理未知错误
      status = HttpStatus.INTERNAL_SERVER_ERROR;
      message = BusinessStatusEnum.SYSTEM_ERROR.getMsg();
      errorCode = BusinessStatusEnum.SYSTEM_ERROR.getCode();
    }

    // 记录日志（生产环境需接入ELK/Sentry）
    console.error(`[${request.method}] ${request.url} - ${message}`);

    response.status(status).json({
      statusCode: status,
      errorCode,
      message,
      path: request.url,
      timestamp: new Date().toISOString(),
    });
  }
}

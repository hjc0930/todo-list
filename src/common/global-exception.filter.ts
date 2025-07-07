import { ArgumentsHost, Catch, ExceptionFilter, HttpException } from '@nestjs/common';
import { Request, Response } from 'express';

@Catch(HttpException)
export class GlobalExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const http = host.switchToHttp();
    const request = http.getRequest<Request>();
    const response = http.getResponse<Response>();

    const statusCode = exception.getStatus();

    const res = exception.getResponse() as { message: string[] };

    response.status(statusCode).json({
      statusCode,
      message: Array.isArray(res?.message) ? res?.message : exception.message,
      path: request.url,
      timestamp: new Date().toISOString(),
    })
  }
}


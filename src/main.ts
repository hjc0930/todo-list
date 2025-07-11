import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  // app.useGlobalFilters(new GlobalExceptionFilter());
  // app.useGlobalPipes(new ValidationPipe({
  //   transform: true,        // 自动类型转换（如字符串转数字）
  //   whitelist: true,         // 过滤 DTO 未定义的属性
  //   forbidNonWhitelisted: true, // 禁止额外属性，返回 400 错误
  // }))
  app.enableCors({
    origin: '*', // 允许所有来源
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE', // 允许的HTTP方法
    allowedHeaders: '*', // 允许的请求头
    credentials: true, // 是否允许携带凭证
  });
  await app.listen(process.env.PORT ?? 8080);
}
void bootstrap();

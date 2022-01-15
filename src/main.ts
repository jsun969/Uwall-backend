import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app/app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.useGlobalPipes(new ValidationPipe());

  const document = SwaggerModule.createDocument(
    app,
    new DocumentBuilder().setTitle('Uwall').build(),
  );
  SwaggerModule.setup('api-docs', app, document);

  await app.listen(3000);
}
bootstrap();

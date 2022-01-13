import { NestFactory } from '@nestjs/core';
import { AppModule } from './app/app.module';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  const document = SwaggerModule.createDocument(
    app,
    new DocumentBuilder().setTitle('Uwall').build(),
  );
  SwaggerModule.setup('api-docs', app, document);

  await app.listen(3000);
}
bootstrap();

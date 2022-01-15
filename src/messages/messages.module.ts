import { Module } from '@nestjs/common';
import { TypegooseModule } from 'nestjs-typegoose';
import { MessagesController } from './messages.controller';
import { Message } from './models/message.model';

@Module({
  imports: [TypegooseModule.forFeature([Message])],
  controllers: [MessagesController],
})
export class MessagesModule {}

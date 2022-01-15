import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { ReturnModelType } from '@typegoose/typegoose';
import { InjectModel } from 'nestjs-typegoose';
import { CreateMessageDto } from './dto/create-message.dto';
import { QueryMessageDto } from './dto/query-messages.dto';
import { Message } from './models/message.model';

@Controller('messages')
@ApiTags('messages')
export class MessagesController {
  constructor(
    @InjectModel(Message) private messageModel: ReturnModelType<typeof Message>,
  ) {}

  @Get()
  @ApiOperation({ summary: 'Get messages' })
  async getMessages(@Query() messagesQuery: QueryMessageDto) {
    return await this.messageModel
      .find(messagesQuery.type ? { type: messagesQuery.type } : {}, null, {
        skip: messagesQuery.skip,
        limit: messagesQuery.limit,
        sort: { createdAt: -1 },
      })
      .exec();
  }

  @Post()
  @ApiOperation({ summary: 'Create message' })
  async createMessage(@Body() createMessageDto: CreateMessageDto) {
    const createMessage = new this.messageModel(createMessageDto);
    return await createMessage.save();
  }
}

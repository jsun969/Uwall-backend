import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { Model } from 'mongoose';
import { CreateMessageDto } from './dto/create-message.dto';
import { MessagesQuery } from './dto/messages-query.dto';
import { MessageDocument } from './schemas/message.schema';

@Controller('messages')
@ApiTags('messages')
export class MessagesController {
  constructor(
    @InjectModel('Message') private messageModel: Model<MessageDocument>,
  ) {}

  @Get()
  @ApiOperation({ summary: 'Get messages' })
  async getMessages(@Query() messagesQuery: MessagesQuery) {
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

import { Body, Controller, Post } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { Model } from 'mongoose';
import { CreateMessageDto } from './dto/create-message.dto';
import { MessageDocument } from './schemas/message.schema';

@Controller('messages')
@ApiTags('messages')
export class MessagesController {
  constructor(
    @InjectModel('Message') private messageModel: Model<MessageDocument>,
  ) {}

  @Post()
  @ApiOperation({ summary: 'Create message' })
  async createMessage(@Body() createMessageDto: CreateMessageDto) {
    const createMessage = new this.messageModel(createMessageDto);
    return await createMessage.save();
  }
}

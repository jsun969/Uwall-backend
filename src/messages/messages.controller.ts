import { Controller, Get } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

@Controller('messages')
@ApiTags('messages')
export class MessagesController {
  @Get()
  @ApiOperation({ summary: 'Get all messages' })
  getMessages() {
    return {
      hello: 'world',
    };
  }
}

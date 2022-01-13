import { Controller, Get } from '@nestjs/common';
import { ApiOperation } from '@nestjs/swagger';

@Controller()
export class AppController {
  @Get('messages')
  @ApiOperation({ summary: 'Get all messages' })
  getMessages() {
    return [{ all: 'messages' }];
  }
}

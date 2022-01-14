import { ApiProperty } from '@nestjs/swagger';

export class MessagesQuery {
  @ApiProperty()
  skip: number;

  @ApiProperty()
  limit: number;

  @ApiProperty({
    required: false,
    enum: ['love', 'complaint', 'help', 'notice', 'expand'],
  })
  type?: string;
}

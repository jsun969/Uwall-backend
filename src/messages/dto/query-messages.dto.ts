import { ApiProperty } from '@nestjs/swagger';
import { IsIn, IsNotEmpty, IsNumberString } from 'class-validator';

export class QueryMessageDto {
  @ApiProperty()
  @IsNumberString()
  @IsNotEmpty()
  skip: number;

  @ApiProperty()
  @IsNumberString()
  @IsNotEmpty()
  limit: number;

  @ApiProperty({
    required: false,
    enum: ['love', 'complaint', 'help', 'notice', 'expand'],
  })
  @IsIn([undefined, 'love', 'complaint', 'help', 'notice', 'expand'])
  type?: string;
}

import { ApiProperty } from '@nestjs/swagger';
import { IsBoolean, IsIn, IsNotEmpty, IsString } from 'class-validator';

export class CreateMessageDto {
  @ApiProperty({
    description: 'love/complaint/help/notice/expand',
    default: 'notice',
  })
  @IsNotEmpty()
  @IsIn(['love', 'complaint', 'help', 'notice', 'expand'])
  type: string;

  @ApiProperty()
  @IsNotEmpty()
  @IsString()
  content: string;

  @ApiProperty()
  @IsNotEmpty()
  fromName: string;

  @ApiProperty({ required: false, description: 'true:girl false:boy' })
  fromGender?: boolean;

  @ApiProperty({ required: false })
  toName?: string;

  @ApiProperty({ required: false })
  toGender?: boolean;

  @ApiProperty()
  @IsNotEmpty()
  @IsBoolean()
  anonymous: boolean;
}

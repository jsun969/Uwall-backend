import { ApiProperty } from '@nestjs/swagger';

export class CreateMessageDto {
  @ApiProperty({
    description: 'love/complaint/help/notice/expand',
    default: 'notice',
  })
  type: string;

  @ApiProperty()
  content: string;

  @ApiProperty()
  fromName: string;

  @ApiProperty({ required: false, description: 'true:girl false:boy' })
  fromGender?: boolean;

  @ApiProperty({ required: false })
  toName?: string;

  @ApiProperty({ required: false })
  toGender?: boolean;

  @ApiProperty()
  anonymous: boolean;
}

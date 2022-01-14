import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type MessageDocument = Message & Document;

@Schema({ timestamps: true })
export class Message extends Document {
  @Prop()
  type: 'love' | 'complaint' | 'help' | 'notice' | 'expand';
  @Prop()
  fromName: string;
  @Prop()
  fromGender?: boolean;
  @Prop()
  toName?: string;
  @Prop()
  toGender?: boolean;
  @Prop()
  content: string;
  @Prop()
  anonymous: boolean;
  @Prop({ default: 0 })
  likes: number;
  @Prop({ default: false })
  show: boolean;
}

export const MessageSchema = SchemaFactory.createForClass(Message);

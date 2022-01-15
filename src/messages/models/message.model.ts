import { Prop } from '@typegoose/typegoose';
import { TimeStamps } from '@typegoose/typegoose/lib/defaultClasses';

export class Message extends TimeStamps {
  @Prop({ required: true })
  type!: 'love' | 'complaint' | 'help' | 'notice' | 'expand';
  @Prop({ required: true })
  fromName!: string;
  @Prop()
  fromGender?: boolean;
  @Prop()
  toName?: string;
  @Prop()
  toGender?: boolean;
  @Prop({ required: true })
  content!: string;
  @Prop({ required: true })
  anonymous!: boolean;
  @Prop({ default: 0 })
  likes: number;
  @Prop({ default: false })
  show: boolean;
}

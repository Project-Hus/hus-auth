import { Module } from '@nestjs/common';
import { RecordService } from 'utils/record/record.service';
import { ApiController } from './api.controller';
import { PredService } from './services/pred.service';

@Module({
  imports: [],
  controllers: [ ApiController ],
  providers: [ PredService, RecordService ],
})
export class ApiModule {}
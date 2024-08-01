import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { CardModule } from './card/card.module';
import {TypeOrmModule} from "@nestjs/typeorm";
import * as process from "node:process";
import {ConfigModule} from "@nestjs/config";

@Module({
  imports: [
    ConfigModule.forRoot({
      ignoreEnvFile: process.env.NODE_ENV !== 'development'
    }),
    TypeOrmModule.forRoot({
      type: 'mongodb',
      host: process.env.DATABASE_HOST,
      port: 27017,
      username: process.env.DATABASE_USERNAME,
      password: process.env.DATABASE_PASSWORD,
      database: process.env.DATABASE_NAME,
      autoLoadEntities: true,
      authSource: 'admin'
    }),
    CardModule
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}

import { Module } from '@nestjs/common';
import { CardController } from './card.controller';
import { CardService } from './card.service';
import {TypeOrmModule} from "@nestjs/typeorm";
import {Card} from "./card.entity";
import {DataSource} from "typeorm";

@Module({
  imports: [TypeOrmModule.forFeature([Card])],
  controllers: [CardController],
  providers: [CardService],
  exports: [TypeOrmModule]
})
export class CardModule {
  constructor(private dataSource: DataSource) {}
}

import {Body, Controller, Get, Post} from "@nestjs/common";
import { CardService } from "./card.service";
import {CardDto} from "./card.dto";

@Controller()
export class CardController {

    constructor(private readonly service: CardService) {}

    @Get('/api/cards')
    async getCards() {
        return await this.service.getCards();
    }
    @Post('/api/cards')
    async mockCardSave(@Body() card: CardDto){
        return await this.service.saveCard(card);
    }
}
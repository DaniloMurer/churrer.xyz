import { Injectable } from "@nestjs/common";
import {DataSource, Repository} from "typeorm";
import {InjectRepository} from "@nestjs/typeorm";
import {Card} from "./card.entity";
import {CardDto} from "./card.dto";

@Injectable()
export class CardService {
    constructor(@InjectRepository(Card) private readonly cardRepository: Repository<Card>) {
    }
    getCards(): Promise<Card[]> {
        return this.cardRepository.find();
    }
    saveCard(card: CardDto): Promise<Card> {
        const newCard = new Card();
        newCard.title = card.title;
        newCard.content = card.content;
        return this.cardRepository.save(newCard);
    }
}
import {Column, Entity, ObjectId, ObjectIdColumn} from "typeorm";

@Entity()
export class Card {
    @ObjectIdColumn()
    _id: ObjectId

    @Column()
    title: string;

    @Column()
    content: string;
}
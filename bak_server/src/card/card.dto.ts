import {IsNotEmpty, IsString} from "class-validator";

export class CardDto {
    @IsString()
    @IsNotEmpty()
    title: string;

    @IsNotEmpty()
    @IsString()
    content: string;
}
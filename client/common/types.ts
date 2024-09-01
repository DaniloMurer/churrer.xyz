export interface Experience {
	position: string;
	company: string;
	timeFrame: string;
	responsibilities: string;
}

export interface Technology {
	name: string;
	logoClass: string;
	experience: string;
	description: string;
}

export class ExperienceDto implements Experience {
	position: string;
	company: string;
	timeFrame: string;
	responsibilities: string;

	constructor(position: string, company: string, timeFrame: string, responsibilities: string) {
		this.position = position;
		this.company = company;
		this.timeFrame = timeFrame;
		this.responsibilities = responsibilities;
	}
}

export class TechnologyDto implements Technology {
	name: string;
	logoClass: string;
	experience: string;
	description: string;

	constructor(name: string, logoClass: string, experience: string, description: string) {
		this.name = name;
		this.logoClass = logoClass;
		this.experience = experience;
		this.description = description;
	}
}

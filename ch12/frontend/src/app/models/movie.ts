export class Movie {
    private name: string;
    private cover: string;
    private description: string;

    constructor(name: string, description: string, cover?: string){
        this.name = name;
        this.description = description;
        this.cover = cover ? cover : "http://via.placeholder.com/185x287";
    }

    public getName(){
        return this.name;
    }

    public getCover(){
        return this.cover;
    }

    public getDescription(){
        return this.description;
    }

    public setName(name: string){
        this.name = name;
    }

    public setCover(cover: string){
        this.cover = cover;
    }

    public setDescription(description: string){
        this.description = description;
    }
}

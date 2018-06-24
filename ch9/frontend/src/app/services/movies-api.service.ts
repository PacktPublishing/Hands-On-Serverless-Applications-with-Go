import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';

@Injectable()
export class MoviesApiService {

  constructor(private http:Http) { }

  findAll(){
    return this.http
      .get(environment.api)
      .map(res => {
        return res.json()
      })
  }

}

import { Injectable, Inject } from '@angular/core';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';
import { Movie } from '../models/movie';
import {LOCAL_STORAGE, WebStorageService} from 'angular-webstorage-service';

@Injectable()
export class MoviesApiService {

  constructor(private http:Http,
              @Inject(LOCAL_STORAGE) private storage: WebStorageService) { }

  findAll(){
    return this.http
      .get(environment.api, {headers: this.getHeaders()})
      .map(res => {
        return res.json()
      })
  }

  insert(movie: Movie){
    return this.http
      .post(environment.api, JSON.stringify(movie), {headers: this.getHeaders()})
      .map(res => {
        return res
      })
  }

  getHeaders(){
    let headers = new Headers()
    headers.append('Authorization', this.storage.get("COGNITO_TOKEN"))
    return headers
  }
}

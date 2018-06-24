import { Component, OnInit } from '@angular/core';
import { Movie } from '../../models/movie';
import { MoviesApiService } from '../../services/movies-api.service';

@Component({
  selector: 'new-movie',
  templateUrl: './new-movie.component.html',
  styleUrls: ['./new-movie.component.css']
})
export class NewMovieComponent implements OnInit {

  private movie : Movie;
  public showMsg: boolean;

  constructor(private moviesApiService: MoviesApiService) {
    this.showMsg = false;
  }

  ngOnInit() {
  }

  save(title, description, cover) {
    this.movie = new Movie(title, description, cover)
    this.moviesApiService.insert(this.movie).subscribe(res => {
      this.showMsg = true;
    }, err => {
      this.showMsg = false;
    })
  }

}

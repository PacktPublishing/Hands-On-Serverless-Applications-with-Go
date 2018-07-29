import { Component, OnInit } from '@angular/core';
import { Movie } from '../../models/movie';
import { MoviesApiService } from '../../services/movies-api.service';

@Component({
  selector: 'list-movies',
  templateUrl: './list-movies.component.html',
  styleUrls: ['./list-movies.component.css']
})
export class ListMoviesComponent implements OnInit {

  public movies: Movie[];

  constructor(private moviesApiService: MoviesApiService) {
    this.movies = []

    this.moviesApiService.findAll().subscribe(res => {
      res.forEach(movie => {
        this.movies.push(new Movie(movie.name, movie.description, movie.cover))
      })
    })
  }

  ngOnInit() {
  }

}

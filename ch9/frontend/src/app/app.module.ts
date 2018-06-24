import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';


import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { ListMoviesComponent } from './components/list-movies/list-movies.component';
import { MovieItemComponent } from './components/movie-item/movie-item.component';
import { MoviesApiService } from './services/movies-api.service';


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ListMoviesComponent,
    MovieItemComponent
  ],
  imports: [
    BrowserModule,
    HttpModule
  ],
  providers: [
    MoviesApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

import { TestBed, inject } from '@angular/core/testing';

import { MoviesApiService } from './movies-api.service';

describe('MoviesApiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [MoviesApiService]
    });
  });

  it('should be created', inject([MoviesApiService], (service: MoviesApiService) => {
    expect(service).toBeTruthy();
  }));
});

import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListMoviesComponent } from './list-movies.component';

describe('ListMoviesComponent', () => {
  let component: ListMoviesComponent;
  let fixture: ComponentFixture<ListMoviesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListMoviesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListMoviesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

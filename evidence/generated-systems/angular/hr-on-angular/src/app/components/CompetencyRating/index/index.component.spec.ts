
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCompetencyRatingComponent } from './index.component';
import { CompetencyRatingService } from '../../../services/CompetencyRating.service';

describe('IndexCompetencyRatingComponent', () => {
  let component: IndexCompetencyRatingComponent;
  let fixture: ComponentFixture<IndexCompetencyRatingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCompetencyRatingComponent
      ],
      providers: [
        CompetencyRatingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCompetencyRatingComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
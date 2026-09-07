
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCompetencyRatingComponent } from './create.component';
import { CompetencyRatingService } from '../../../services/CompetencyRating.service';
import { Router } from '@angular/router';

describe('CreateCompetencyRatingComponent', () => {
  let component: CreateCompetencyRatingComponent;
  let fixture: ComponentFixture<CreateCompetencyRatingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCompetencyRatingComponent
      ],
      providers: [
        CompetencyRatingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCompetencyRatingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
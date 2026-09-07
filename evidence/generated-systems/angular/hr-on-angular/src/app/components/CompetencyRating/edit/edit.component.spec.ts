
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditCompetencyRatingComponent } from './edit.component';
import { CompetencyRatingService } from '../../../services/CompetencyRating.service';

describe('EditCompetencyRatingComponent', () => {
  let component: EditCompetencyRatingComponent;
  let fixture: ComponentFixture<EditCompetencyRatingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditCompetencyRatingComponent
      ],
      providers: [
        CompetencyRatingService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditCompetencyRatingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
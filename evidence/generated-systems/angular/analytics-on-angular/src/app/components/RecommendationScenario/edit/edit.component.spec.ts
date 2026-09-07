
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditRecommendationScenarioComponent } from './edit.component';
import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';

describe('EditRecommendationScenarioComponent', () => {
  let component: EditRecommendationScenarioComponent;
  let fixture: ComponentFixture<EditRecommendationScenarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditRecommendationScenarioComponent
      ],
      providers: [
        RecommendationScenarioService,
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

    fixture = TestBed.createComponent(EditRecommendationScenarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
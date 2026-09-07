
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRecommendationScenarioComponent } from './create.component';
import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';
import { Router } from '@angular/router';

describe('CreateRecommendationScenarioComponent', () => {
  let component: CreateRecommendationScenarioComponent;
  let fixture: ComponentFixture<CreateRecommendationScenarioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRecommendationScenarioComponent
      ],
      providers: [
        RecommendationScenarioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRecommendationScenarioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
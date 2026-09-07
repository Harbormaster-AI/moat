
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEvaluationMetricComponent } from './create.component';
import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';
import { Router } from '@angular/router';

describe('CreateEvaluationMetricComponent', () => {
  let component: CreateEvaluationMetricComponent;
  let fixture: ComponentFixture<CreateEvaluationMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEvaluationMetricComponent
      ],
      providers: [
        EvaluationMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEvaluationMetricComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEvaluationMetricComponent } from './index.component';
import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';

describe('IndexEvaluationMetricComponent', () => {
  let component: IndexEvaluationMetricComponent;
  let fixture: ComponentFixture<IndexEvaluationMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEvaluationMetricComponent
      ],
      providers: [
        EvaluationMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEvaluationMetricComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
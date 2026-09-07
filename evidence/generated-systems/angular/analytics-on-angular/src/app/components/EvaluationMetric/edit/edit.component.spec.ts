
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditEvaluationMetricComponent } from './edit.component';
import { EvaluationMetricService } from '../../../services/EvaluationMetric.service';

describe('EditEvaluationMetricComponent', () => {
  let component: EditEvaluationMetricComponent;
  let fixture: ComponentFixture<EditEvaluationMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditEvaluationMetricComponent
      ],
      providers: [
        EvaluationMetricService,
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

    fixture = TestBed.createComponent(EditEvaluationMetricComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
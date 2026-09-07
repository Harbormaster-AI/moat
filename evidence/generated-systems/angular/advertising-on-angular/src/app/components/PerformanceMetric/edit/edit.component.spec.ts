
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPerformanceMetricComponent } from './edit.component';
import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';

describe('EditPerformanceMetricComponent', () => {
  let component: EditPerformanceMetricComponent;
  let fixture: ComponentFixture<EditPerformanceMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPerformanceMetricComponent
      ],
      providers: [
        PerformanceMetricService,
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

    fixture = TestBed.createComponent(EditPerformanceMetricComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
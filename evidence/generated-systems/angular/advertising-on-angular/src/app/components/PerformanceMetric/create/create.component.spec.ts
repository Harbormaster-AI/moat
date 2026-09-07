
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePerformanceMetricComponent } from './create.component';
import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';
import { Router } from '@angular/router';

describe('CreatePerformanceMetricComponent', () => {
  let component: CreatePerformanceMetricComponent;
  let fixture: ComponentFixture<CreatePerformanceMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePerformanceMetricComponent
      ],
      providers: [
        PerformanceMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePerformanceMetricComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
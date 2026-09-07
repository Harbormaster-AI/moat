
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRunMetricComponent } from './create.component';
import { RunMetricService } from '../../../services/RunMetric.service';
import { Router } from '@angular/router';

describe('CreateRunMetricComponent', () => {
  let component: CreateRunMetricComponent;
  let fixture: ComponentFixture<CreateRunMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRunMetricComponent
      ],
      providers: [
        RunMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRunMetricComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
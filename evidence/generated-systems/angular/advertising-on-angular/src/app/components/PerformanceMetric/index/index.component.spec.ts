
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPerformanceMetricComponent } from './index.component';
import { PerformanceMetricService } from '../../../services/PerformanceMetric.service';

describe('IndexPerformanceMetricComponent', () => {
  let component: IndexPerformanceMetricComponent;
  let fixture: ComponentFixture<IndexPerformanceMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPerformanceMetricComponent
      ],
      providers: [
        PerformanceMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPerformanceMetricComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
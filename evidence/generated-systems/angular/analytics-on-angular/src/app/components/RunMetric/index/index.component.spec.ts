
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexRunMetricComponent } from './index.component';
import { RunMetricService } from '../../../services/RunMetric.service';

describe('IndexRunMetricComponent', () => {
  let component: IndexRunMetricComponent;
  let fixture: ComponentFixture<IndexRunMetricComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexRunMetricComponent
      ],
      providers: [
        RunMetricService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexRunMetricComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
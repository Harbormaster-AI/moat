
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTimeSeriesComponent } from './index.component';
import { TimeSeriesService } from '../../../services/TimeSeries.service';

describe('IndexTimeSeriesComponent', () => {
  let component: IndexTimeSeriesComponent;
  let fixture: ComponentFixture<IndexTimeSeriesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTimeSeriesComponent
      ],
      providers: [
        TimeSeriesService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTimeSeriesComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
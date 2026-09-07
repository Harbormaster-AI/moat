
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTimeSeriesComponent } from './create.component';
import { TimeSeriesService } from '../../../services/TimeSeries.service';
import { Router } from '@angular/router';

describe('CreateTimeSeriesComponent', () => {
  let component: CreateTimeSeriesComponent;
  let fixture: ComponentFixture<CreateTimeSeriesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTimeSeriesComponent
      ],
      providers: [
        TimeSeriesService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTimeSeriesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
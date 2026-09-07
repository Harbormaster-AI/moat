
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateForecastComponent } from './create.component';
import { ForecastService } from '../../../services/Forecast.service';
import { Router } from '@angular/router';

describe('CreateForecastComponent', () => {
  let component: CreateForecastComponent;
  let fixture: ComponentFixture<CreateForecastComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateForecastComponent
      ],
      providers: [
        ForecastService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateForecastComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
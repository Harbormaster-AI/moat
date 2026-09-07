
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateForecastLineComponent } from './create.component';
import { ForecastLineService } from '../../../services/ForecastLine.service';
import { Router } from '@angular/router';

describe('CreateForecastLineComponent', () => {
  let component: CreateForecastLineComponent;
  let fixture: ComponentFixture<CreateForecastLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateForecastLineComponent
      ],
      providers: [
        ForecastLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateForecastLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
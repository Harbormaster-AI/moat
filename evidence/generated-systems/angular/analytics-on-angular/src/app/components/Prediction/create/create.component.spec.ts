
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePredictionComponent } from './create.component';
import { PredictionService } from '../../../services/Prediction.service';
import { Router } from '@angular/router';

describe('CreatePredictionComponent', () => {
  let component: CreatePredictionComponent;
  let fixture: ComponentFixture<CreatePredictionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePredictionComponent
      ],
      providers: [
        PredictionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePredictionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
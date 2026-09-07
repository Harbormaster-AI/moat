
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPredictionComponent } from './index.component';
import { PredictionService } from '../../../services/Prediction.service';

describe('IndexPredictionComponent', () => {
  let component: IndexPredictionComponent;
  let fixture: ComponentFixture<IndexPredictionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPredictionComponent
      ],
      providers: [
        PredictionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPredictionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
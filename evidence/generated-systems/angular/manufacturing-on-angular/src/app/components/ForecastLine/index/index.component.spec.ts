
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexForecastLineComponent } from './index.component';
import { ForecastLineService } from '../../../services/ForecastLine.service';

describe('IndexForecastLineComponent', () => {
  let component: IndexForecastLineComponent;
  let fixture: ComponentFixture<IndexForecastLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexForecastLineComponent
      ],
      providers: [
        ForecastLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexForecastLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
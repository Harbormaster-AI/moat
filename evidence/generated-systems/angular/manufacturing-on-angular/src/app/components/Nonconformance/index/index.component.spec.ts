
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexNonconformanceComponent } from './index.component';
import { NonconformanceService } from '../../../services/Nonconformance.service';

describe('IndexNonconformanceComponent', () => {
  let component: IndexNonconformanceComponent;
  let fixture: ComponentFixture<IndexNonconformanceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexNonconformanceComponent
      ],
      providers: [
        NonconformanceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexNonconformanceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
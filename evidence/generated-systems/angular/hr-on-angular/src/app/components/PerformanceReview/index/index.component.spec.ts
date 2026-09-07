
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPerformanceReviewComponent } from './index.component';
import { PerformanceReviewService } from '../../../services/PerformanceReview.service';

describe('IndexPerformanceReviewComponent', () => {
  let component: IndexPerformanceReviewComponent;
  let fixture: ComponentFixture<IndexPerformanceReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPerformanceReviewComponent
      ],
      providers: [
        PerformanceReviewService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPerformanceReviewComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
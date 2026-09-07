
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePerformanceReviewComponent } from './create.component';
import { PerformanceReviewService } from '../../../services/PerformanceReview.service';
import { Router } from '@angular/router';

describe('CreatePerformanceReviewComponent', () => {
  let component: CreatePerformanceReviewComponent;
  let fixture: ComponentFixture<CreatePerformanceReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePerformanceReviewComponent
      ],
      providers: [
        PerformanceReviewService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePerformanceReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
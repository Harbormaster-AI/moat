
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDispositionReviewComponent } from './index.component';
import { DispositionReviewService } from '../../../services/DispositionReview.service';

describe('IndexDispositionReviewComponent', () => {
  let component: IndexDispositionReviewComponent;
  let fixture: ComponentFixture<IndexDispositionReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDispositionReviewComponent
      ],
      providers: [
        DispositionReviewService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDispositionReviewComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
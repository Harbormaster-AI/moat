
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDispositionReviewComponent } from './create.component';
import { DispositionReviewService } from '../../../services/DispositionReview.service';
import { Router } from '@angular/router';

describe('CreateDispositionReviewComponent', () => {
  let component: CreateDispositionReviewComponent;
  let fixture: ComponentFixture<CreateDispositionReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDispositionReviewComponent
      ],
      providers: [
        DispositionReviewService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDispositionReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
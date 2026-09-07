
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBenefitEnrollmentComponent } from './index.component';
import { BenefitEnrollmentService } from '../../../services/BenefitEnrollment.service';

describe('IndexBenefitEnrollmentComponent', () => {
  let component: IndexBenefitEnrollmentComponent;
  let fixture: ComponentFixture<IndexBenefitEnrollmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBenefitEnrollmentComponent
      ],
      providers: [
        BenefitEnrollmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBenefitEnrollmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBenefitEnrollmentComponent } from './create.component';
import { BenefitEnrollmentService } from '../../../services/BenefitEnrollment.service';
import { Router } from '@angular/router';

describe('CreateBenefitEnrollmentComponent', () => {
  let component: CreateBenefitEnrollmentComponent;
  let fixture: ComponentFixture<CreateBenefitEnrollmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBenefitEnrollmentComponent
      ],
      providers: [
        BenefitEnrollmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBenefitEnrollmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBenefitPlanComponent } from './create.component';
import { BenefitPlanService } from '../../../services/BenefitPlan.service';
import { Router } from '@angular/router';

describe('CreateBenefitPlanComponent', () => {
  let component: CreateBenefitPlanComponent;
  let fixture: ComponentFixture<CreateBenefitPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBenefitPlanComponent
      ],
      providers: [
        BenefitPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBenefitPlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
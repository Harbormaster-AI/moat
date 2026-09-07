
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBenefitPlanComponent } from './index.component';
import { BenefitPlanService } from '../../../services/BenefitPlan.service';

describe('IndexBenefitPlanComponent', () => {
  let component: IndexBenefitPlanComponent;
  let fixture: ComponentFixture<IndexBenefitPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBenefitPlanComponent
      ],
      providers: [
        BenefitPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBenefitPlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
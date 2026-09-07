
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePricingPlanComponent } from './create.component';
import { PricingPlanService } from '../../../services/PricingPlan.service';
import { Router } from '@angular/router';

describe('CreatePricingPlanComponent', () => {
  let component: CreatePricingPlanComponent;
  let fixture: ComponentFixture<CreatePricingPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePricingPlanComponent
      ],
      providers: [
        PricingPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePricingPlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
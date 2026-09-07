
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPricingPlanComponent } from './index.component';
import { PricingPlanService } from '../../../services/PricingPlan.service';

describe('IndexPricingPlanComponent', () => {
  let component: IndexPricingPlanComponent;
  let fixture: ComponentFixture<IndexPricingPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPricingPlanComponent
      ],
      providers: [
        PricingPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPricingPlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsurancePlanComponent } from './index.component';
import { InsurancePlanService } from '../../../services/InsurancePlan.service';

describe('IndexInsurancePlanComponent', () => {
  let component: IndexInsurancePlanComponent;
  let fixture: ComponentFixture<IndexInsurancePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsurancePlanComponent
      ],
      providers: [
        InsurancePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsurancePlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
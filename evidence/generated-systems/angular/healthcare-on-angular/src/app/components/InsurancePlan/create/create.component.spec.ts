
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsurancePlanComponent } from './create.component';
import { InsurancePlanService } from '../../../services/InsurancePlan.service';
import { Router } from '@angular/router';

describe('CreateInsurancePlanComponent', () => {
  let component: CreateInsurancePlanComponent;
  let fixture: ComponentFixture<CreateInsurancePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsurancePlanComponent
      ],
      providers: [
        InsurancePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsurancePlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
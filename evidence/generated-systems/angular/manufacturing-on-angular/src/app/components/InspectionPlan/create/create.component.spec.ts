
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInspectionPlanComponent } from './create.component';
import { InspectionPlanService } from '../../../services/InspectionPlan.service';
import { Router } from '@angular/router';

describe('CreateInspectionPlanComponent', () => {
  let component: CreateInspectionPlanComponent;
  let fixture: ComponentFixture<CreateInspectionPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInspectionPlanComponent
      ],
      providers: [
        InspectionPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInspectionPlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
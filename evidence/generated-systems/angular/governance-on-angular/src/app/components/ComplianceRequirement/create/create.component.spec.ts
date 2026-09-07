
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateComplianceRequirementComponent } from './create.component';
import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';
import { Router } from '@angular/router';

describe('CreateComplianceRequirementComponent', () => {
  let component: CreateComplianceRequirementComponent;
  let fixture: ComponentFixture<CreateComplianceRequirementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateComplianceRequirementComponent
      ],
      providers: [
        ComplianceRequirementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateComplianceRequirementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
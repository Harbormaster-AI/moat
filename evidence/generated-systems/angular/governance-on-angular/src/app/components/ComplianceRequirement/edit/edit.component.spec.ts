
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditComplianceRequirementComponent } from './edit.component';
import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';

describe('EditComplianceRequirementComponent', () => {
  let component: EditComplianceRequirementComponent;
  let fixture: ComponentFixture<EditComplianceRequirementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditComplianceRequirementComponent
      ],
      providers: [
        ComplianceRequirementService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditComplianceRequirementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
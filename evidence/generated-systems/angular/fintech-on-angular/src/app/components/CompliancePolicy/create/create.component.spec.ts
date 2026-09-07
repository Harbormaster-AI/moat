
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCompliancePolicyComponent } from './create.component';
import { CompliancePolicyService } from '../../../services/CompliancePolicy.service';
import { Router } from '@angular/router';

describe('CreateCompliancePolicyComponent', () => {
  let component: CreateCompliancePolicyComponent;
  let fixture: ComponentFixture<CreateCompliancePolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCompliancePolicyComponent
      ],
      providers: [
        CompliancePolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCompliancePolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
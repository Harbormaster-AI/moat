
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexComplianceRequirementComponent } from './index.component';
import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';

describe('IndexComplianceRequirementComponent', () => {
  let component: IndexComplianceRequirementComponent;
  let fixture: ComponentFixture<IndexComplianceRequirementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexComplianceRequirementComponent
      ],
      providers: [
        ComplianceRequirementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexComplianceRequirementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
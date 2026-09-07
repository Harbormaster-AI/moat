
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCompliancePolicyComponent } from './index.component';
import { CompliancePolicyService } from '../../../services/CompliancePolicy.service';

describe('IndexCompliancePolicyComponent', () => {
  let component: IndexCompliancePolicyComponent;
  let fixture: ComponentFixture<IndexCompliancePolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCompliancePolicyComponent
      ],
      providers: [
        CompliancePolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCompliancePolicyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateComplianceAlertComponent } from './create.component';
import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';
import { Router } from '@angular/router';

describe('CreateComplianceAlertComponent', () => {
  let component: CreateComplianceAlertComponent;
  let fixture: ComponentFixture<CreateComplianceAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateComplianceAlertComponent
      ],
      providers: [
        ComplianceAlertService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateComplianceAlertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
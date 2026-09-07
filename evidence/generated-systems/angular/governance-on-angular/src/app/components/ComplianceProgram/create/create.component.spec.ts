
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateComplianceProgramComponent } from './create.component';
import { ComplianceProgramService } from '../../../services/ComplianceProgram.service';
import { Router } from '@angular/router';

describe('CreateComplianceProgramComponent', () => {
  let component: CreateComplianceProgramComponent;
  let fixture: ComponentFixture<CreateComplianceProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateComplianceProgramComponent
      ],
      providers: [
        ComplianceProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateComplianceProgramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
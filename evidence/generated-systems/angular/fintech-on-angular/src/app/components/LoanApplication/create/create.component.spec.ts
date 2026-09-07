
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLoanApplicationComponent } from './create.component';
import { LoanApplicationService } from '../../../services/LoanApplication.service';
import { Router } from '@angular/router';

describe('CreateLoanApplicationComponent', () => {
  let component: CreateLoanApplicationComponent;
  let fixture: ComponentFixture<CreateLoanApplicationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLoanApplicationComponent
      ],
      providers: [
        LoanApplicationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLoanApplicationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
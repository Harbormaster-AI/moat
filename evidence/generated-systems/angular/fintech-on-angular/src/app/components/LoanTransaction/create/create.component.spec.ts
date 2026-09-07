
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLoanTransactionComponent } from './create.component';
import { LoanTransactionService } from '../../../services/LoanTransaction.service';
import { Router } from '@angular/router';

describe('CreateLoanTransactionComponent', () => {
  let component: CreateLoanTransactionComponent;
  let fixture: ComponentFixture<CreateLoanTransactionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLoanTransactionComponent
      ],
      providers: [
        LoanTransactionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLoanTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
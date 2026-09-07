
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLoanTransactionComponent } from './index.component';
import { LoanTransactionService } from '../../../services/LoanTransaction.service';

describe('IndexLoanTransactionComponent', () => {
  let component: IndexLoanTransactionComponent;
  let fixture: ComponentFixture<IndexLoanTransactionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLoanTransactionComponent
      ],
      providers: [
        LoanTransactionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLoanTransactionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
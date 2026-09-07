
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLoanApplicationComponent } from './index.component';
import { LoanApplicationService } from '../../../services/LoanApplication.service';

describe('IndexLoanApplicationComponent', () => {
  let component: IndexLoanApplicationComponent;
  let fixture: ComponentFixture<IndexLoanApplicationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLoanApplicationComponent
      ],
      providers: [
        LoanApplicationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLoanApplicationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
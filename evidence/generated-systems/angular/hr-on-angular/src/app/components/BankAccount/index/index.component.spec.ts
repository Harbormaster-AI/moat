
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBankAccountComponent } from './index.component';
import { BankAccountService } from '../../../services/BankAccount.service';

describe('IndexBankAccountComponent', () => {
  let component: IndexBankAccountComponent;
  let fixture: ComponentFixture<IndexBankAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBankAccountComponent
      ],
      providers: [
        BankAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBankAccountComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
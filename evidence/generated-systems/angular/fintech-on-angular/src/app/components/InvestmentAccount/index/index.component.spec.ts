
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInvestmentAccountComponent } from './index.component';
import { InvestmentAccountService } from '../../../services/InvestmentAccount.service';

describe('IndexInvestmentAccountComponent', () => {
  let component: IndexInvestmentAccountComponent;
  let fixture: ComponentFixture<IndexInvestmentAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInvestmentAccountComponent
      ],
      providers: [
        InvestmentAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInvestmentAccountComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
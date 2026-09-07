
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInvestmentPortfolioComponent } from './index.component';
import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';

describe('IndexInvestmentPortfolioComponent', () => {
  let component: IndexInvestmentPortfolioComponent;
  let fixture: ComponentFixture<IndexInvestmentPortfolioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInvestmentPortfolioComponent
      ],
      providers: [
        InvestmentPortfolioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInvestmentPortfolioComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
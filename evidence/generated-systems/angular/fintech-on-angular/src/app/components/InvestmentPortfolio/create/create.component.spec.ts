
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInvestmentPortfolioComponent } from './create.component';
import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';
import { Router } from '@angular/router';

describe('CreateInvestmentPortfolioComponent', () => {
  let component: CreateInvestmentPortfolioComponent;
  let fixture: ComponentFixture<CreateInvestmentPortfolioComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInvestmentPortfolioComponent
      ],
      providers: [
        InvestmentPortfolioService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInvestmentPortfolioComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
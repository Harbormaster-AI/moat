import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InvestmentPortfolioService } from './InvestmentPortfolio.service';

describe('InvestmentPortfolioService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InvestmentPortfolioService] });
	});

  it('should be created', () => {
    const service: InvestmentPortfolioService = TestBed.get(InvestmentPortfolioService);
    expect(service).toBeTruthy();
  });
});

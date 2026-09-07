import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';
import { InvestmentPortfolio } from '../../../models/InvestmentPortfolio';
import { SubBaseComponent } from '../../InvestmentPortfolio/sub.base.component';

@Component({
    selector: 'app-create-investmentPortfolio',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInvestmentPortfolioComponent extends SubBaseComponent implements OnInit {

    title = 'Add InvestmentPortfolio';

    investmentPortfolioForm: FormGroup;
    investmentPortfolio: InvestmentPortfolio;

    constructor( http: HttpClient,
        private investmentPortfolioService: InvestmentPortfolioService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.investmentPortfolioForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  portfolioCode: ['', Validators.required],
      baseCurrency: ['', Validators.required],
      createdAt: ['', Validators.required],
      Customer: ['', ],
      Accounts: ['', ],
      Orders: ['', ],
      Holdings: ['', ],
      Status: ['', ]
        });
    }

    
    addInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status): void {
        this.investmentPortfolioService
        .addInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInvestmentPortfolio']);
            });
    }

    ngOnInit(): void {
    }
}
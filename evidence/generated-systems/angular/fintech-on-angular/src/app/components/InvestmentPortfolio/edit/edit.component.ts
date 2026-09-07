import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InvestmentPortfolioService } from '../../../services/InvestmentPortfolio.service';
import { SubBaseComponent } from '../../InvestmentPortfolio/sub.base.component';


@Component({
    selector: 'app-edit-investmentPortfolio',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInvestmentPortfolioComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InvestmentPortfolio';

    investmentPortfolioForm: FormGroup;
    investmentPortfolio: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InvestmentPortfolioService,
        private fb: FormBuilder
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

    
    updateInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInvestmentPortfolio(portfolioCode, baseCurrency, createdAt, Customer, Accounts, Orders, Holdings, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInvestmentPortfolio']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInvestmentPortfolio(params['id']).subscribe(res => {
                this.investmentPortfolio = res;
            });
        });
    }
}
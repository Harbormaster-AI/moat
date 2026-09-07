import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InvestmentAccountService } from '../../../services/InvestmentAccount.service';
import { SubBaseComponent } from '../../InvestmentAccount/sub.base.component';


@Component({
    selector: 'app-edit-investmentAccount',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInvestmentAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InvestmentAccount';

    investmentAccountForm: FormGroup;
    investmentAccount: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InvestmentAccountService,
        private fb: FormBuilder
) {
        super(http);
        this.investmentAccountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  accountNumber: ['', Validators.required],
      baseCurrency: ['', Validators.required],
      balance: ['', Validators.required],
      Portfolio: ['', ],
      Trades: ['', ],
      Orders: ['', ],
      AccountType: ['', ],
      Status: ['', ]
        });
    }

    
    updateInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInvestmentAccount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInvestmentAccount(params['id']).subscribe(res => {
                this.investmentAccount = res;
            });
        });
    }
}
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InvestmentAccountService } from '../../../services/InvestmentAccount.service';
import { InvestmentAccount } from '../../../models/InvestmentAccount';
import { SubBaseComponent } from '../../InvestmentAccount/sub.base.component';

@Component({
    selector: 'app-create-investmentAccount',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInvestmentAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Add InvestmentAccount';

    investmentAccountForm: FormGroup;
    investmentAccount: InvestmentAccount;

    constructor( http: HttpClient,
        private investmentAccountService: InvestmentAccountService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status): void {
        this.investmentAccountService
        .addInvestmentAccount(accountNumber, baseCurrency, balance, Portfolio, Trades, Orders, AccountType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInvestmentAccount']);
            });
    }

    ngOnInit(): void {
    }
}
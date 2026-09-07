import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AccountService } from '../../../services/Account.service';
import { SubBaseComponent } from '../../Account/sub.base.component';


@Component({
    selector: 'app-edit-account',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Account';

    accountForm: FormGroup;
    account: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AccountService,
        private fb: FormBuilder
) {
        super(http);
        this.accountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  accountNumber: ['', Validators.required],
      iban: ['', Validators.required],
      bic: ['', Validators.required],
      openedDate: ['', Validators.required],
      currency: ['', Validators.required],
      balance: ['', Validators.required],
      availableBalance: ['', Validators.required],
      Customer: ['', ],
      Institution: ['', ],
      Transactions: ['', ],
      Cards: ['', ],
      Statements: ['', ],
      Mandates: ['', ],
      AccountType: ['', ],
      Status: ['', ]
        });
    }

    
    updateAccount(accountNumber, iban, bic, openedDate, currency, balance, availableBalance, Customer, Institution, Transactions, Cards, Statements, Mandates, AccountType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAccount(accountNumber, iban, bic, openedDate, currency, balance, availableBalance, Customer, Institution, Transactions, Cards, Statements, Mandates, AccountType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAccount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAccount(params['id']).subscribe(res => {
                this.account = res;
            });
        });
    }
}
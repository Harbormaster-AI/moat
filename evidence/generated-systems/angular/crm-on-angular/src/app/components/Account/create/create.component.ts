import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AccountService } from '../../../services/Account.service';
import { Account } from '../../../models/Account';
import { SubBaseComponent } from '../../Account/sub.base.component';

@Component({
    selector: 'app-create-account',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Add Account';

    accountForm: FormGroup;
    account: Account;

    constructor( http: HttpClient,
        private accountService: AccountService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.accountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      accountNumber: ['', Validators.required],
      industry: ['', Validators.required],
      billingAddress: ['', Validators.required],
      shippingAddress: ['', Validators.required],
      website: ['', Validators.required],
      phone: ['', Validators.required],
      asActive: ['', Validators.required],
      Organization: ['', ],
      ParentAccount: ['', ],
      ChildAccounts: ['', ],
      Contacts: ['', ],
      Opportunities: ['', ],
      Cases: ['', ],
      Owner: ['', ],
      Territory: ['', ],
      Activities: ['', ],
      Campaigns: ['', ],
      Quotes: ['', ],
      Orders: ['', ],
      Contracts: ['', ],
      Notes: ['', ],
      EmailMessages: ['', ],
      AccountType: ['', ],
      LifecycleStage: ['', ]
        });
    }

    
    addAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage): void {
        this.accountService
        .addAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage)
            .subscribe(() => {
                this.router.navigate(['/indexAccount']);
            });
    }

    ngOnInit(): void {
    }
}
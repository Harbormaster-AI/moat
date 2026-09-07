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

    
    updateAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAccount(name, accountNumber, industry, billingAddress, shippingAddress, website, phone, asActive, Organization, ParentAccount, ChildAccounts, Contacts, Opportunities, Cases, Owner, Territory, Activities, Campaigns, Quotes, Orders, Contracts, Notes, EmailMessages, AccountType, LifecycleStage, params['id'])
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
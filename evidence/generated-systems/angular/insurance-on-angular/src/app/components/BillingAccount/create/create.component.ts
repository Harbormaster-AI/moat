import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BillingAccountService } from '../../../services/BillingAccount.service';
import { BillingAccount } from '../../../models/BillingAccount';
import { SubBaseComponent } from '../../BillingAccount/sub.base.component';

@Component({
    selector: 'app-create-billingAccount',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBillingAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Add BillingAccount';

    billingAccountForm: FormGroup;
    billingAccount: BillingAccount;

    constructor( http: HttpClient,
        private billingAccountService: BillingAccountService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.billingAccountForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  accountNumber: ['', Validators.required],
      balance: ['', Validators.required],
      Customer: ['', ],
      Policies: ['', ],
      Invoices: ['', ],
      Payments: ['', ],
      Status: ['', ]
        });
    }

    
    addBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status): void {
        this.billingAccountService
        .addBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status)
            .subscribe(() => {
                this.router.navigate(['/indexBillingAccount']);
            });
    }

    ngOnInit(): void {
    }
}
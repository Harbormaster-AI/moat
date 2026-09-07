import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BillingAccountService } from '../../../services/BillingAccount.service';
import { SubBaseComponent } from '../../BillingAccount/sub.base.component';


@Component({
    selector: 'app-edit-billingAccount',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBillingAccountComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BillingAccount';

    billingAccountForm: FormGroup;
    billingAccount: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BillingAccountService,
        private fb: FormBuilder
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

    
    updateBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBillingAccount(accountNumber, balance, Customer, Policies, Invoices, Payments, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBillingAccount']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBillingAccount(params['id']).subscribe(res => {
                this.billingAccount = res;
            });
        });
    }
}
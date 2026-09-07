import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CustomerService } from '../../../services/Customer.service';
import { SubBaseComponent } from '../../Customer/sub.base.component';


@Component({
    selector: 'app-edit-customer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCustomerComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Customer';

    customerForm: FormGroup;
    customer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CustomerService,
        private fb: FormBuilder
) {
        super(http);
        this.customerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      dateOfBirth: ['', Validators.required],
      email: ['', Validators.required],
      phone: ['', Validators.required],
      address: ['', Validators.required],
      taxId: ['', Validators.required],
      riskScore: ['', Validators.required],
      Institution: ['', ],
      Accounts: ['', ],
      Wallets: ['', ],
      Cards: ['', ],
      KycProfiles: ['', ],
      Consents: ['', ],
      Agreements: ['', ],
      LoanApplications: ['', ],
      Loans: ['', ],
      Portfolios: ['', ],
      Disputes: ['', ],
      CustomerType: ['', ]
        });
    }

    
    updateCustomer(firstName, lastName, dateOfBirth, email, phone, address, taxId, riskScore, Institution, Accounts, Wallets, Cards, KycProfiles, Consents, Agreements, LoanApplications, Loans, Portfolios, Disputes, CustomerType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCustomer(firstName, lastName, dateOfBirth, email, phone, address, taxId, riskScore, Institution, Accounts, Wallets, Cards, KycProfiles, Consents, Agreements, LoanApplications, Loans, Portfolios, Disputes, CustomerType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCustomer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCustomer(params['id']).subscribe(res => {
                this.customer = res;
            });
        });
    }
}
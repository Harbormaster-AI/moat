import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentService } from '../../../services/Payment.service';
import { SubBaseComponent } from '../../Payment/sub.base.component';


@Component({
    selector: 'app-edit-payment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Payment';

    paymentForm: FormGroup;
    payment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentService,
        private fb: FormBuilder
) {
        super(http);
        this.paymentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  paymentReference: ['', Validators.required],
      amount: ['', Validators.required],
      paymentDate: ['', Validators.required],
      Invoice: ['', ],
      BillingAccount: ['', ],
      Policy: ['', ],
      Method: ['', ],
      Status: ['', ]
        });
    }

    
    updatePayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPayment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPayment(params['id']).subscribe(res => {
                this.payment = res;
            });
        });
    }
}
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentService } from '../../../services/Payment.service';
import { Payment } from '../../../models/Payment';
import { SubBaseComponent } from '../../Payment/sub.base.component';

@Component({
    selector: 'app-create-payment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentComponent extends SubBaseComponent implements OnInit {

    title = 'Add Payment';

    paymentForm: FormGroup;
    payment: Payment;

    constructor( http: HttpClient,
        private paymentService: PaymentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status): void {
        this.paymentService
        .addPayment(paymentReference, amount, paymentDate, Invoice, BillingAccount, Policy, Method, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPayment']);
            });
    }

    ngOnInit(): void {
    }
}
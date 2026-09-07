import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentMethodService } from '../../../services/PaymentMethod.service';
import { PaymentMethod } from '../../../models/PaymentMethod';
import { SubBaseComponent } from '../../PaymentMethod/sub.base.component';

@Component({
    selector: 'app-create-paymentMethod',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentMethodComponent extends SubBaseComponent implements OnInit {

    title = 'Add PaymentMethod';

    paymentMethodForm: FormGroup;
    paymentMethod: PaymentMethod;

    constructor( http: HttpClient,
        private paymentMethodService: PaymentMethodService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.paymentMethodForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  last4: ['', Validators.required],
      cardholderName: ['', Validators.required],
      billingAddress: ['', Validators.required],
      BillingProfile: ['', ],
      MethodType: ['', ]
        });
    }

    
    addPaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType): void {
        this.paymentMethodService
        .addPaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType)
            .subscribe(() => {
                this.router.navigate(['/indexPaymentMethod']);
            });
    }

    ngOnInit(): void {
    }
}
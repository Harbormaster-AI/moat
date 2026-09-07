import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentMethodService } from '../../../services/PaymentMethod.service';
import { SubBaseComponent } from '../../PaymentMethod/sub.base.component';


@Component({
    selector: 'app-edit-paymentMethod',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentMethodComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PaymentMethod';

    paymentMethodForm: FormGroup;
    paymentMethod: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentMethodService,
        private fb: FormBuilder
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

    
    updatePaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePaymentMethod(last4, cardholderName, billingAddress, BillingProfile, MethodType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPaymentMethod']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPaymentMethod(params['id']).subscribe(res => {
                this.paymentMethod = res;
            });
        });
    }
}
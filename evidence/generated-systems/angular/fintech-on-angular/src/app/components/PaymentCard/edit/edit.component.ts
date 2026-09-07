import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PaymentCardService } from '../../../services/PaymentCard.service';
import { SubBaseComponent } from '../../PaymentCard/sub.base.component';


@Component({
    selector: 'app-edit-paymentCard',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPaymentCardComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PaymentCard';

    paymentCardForm: FormGroup;
    paymentCard: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PaymentCardService,
        private fb: FormBuilder
) {
        super(http);
        this.paymentCardForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  cardToken: ['', Validators.required],
      maskedPan: ['', Validators.required],
      expiryMonth: ['', Validators.required],
      expiryYear: ['', Validators.required],
      cardholderName: ['', Validators.required],
      Customer: ['', ],
      Account: ['', ],
      Tokenizations: ['', ],
      Disputes: ['', ],
      Scheme: ['', ],
      Status: ['', ]
        });
    }

    
    updatePaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPaymentCard']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPaymentCard(params['id']).subscribe(res => {
                this.paymentCard = res;
            });
        });
    }
}
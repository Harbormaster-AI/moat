import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PaymentCardService } from '../../../services/PaymentCard.service';
import { PaymentCard } from '../../../models/PaymentCard';
import { SubBaseComponent } from '../../PaymentCard/sub.base.component';

@Component({
    selector: 'app-create-paymentCard',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePaymentCardComponent extends SubBaseComponent implements OnInit {

    title = 'Add PaymentCard';

    paymentCardForm: FormGroup;
    paymentCard: PaymentCard;

    constructor( http: HttpClient,
        private paymentCardService: PaymentCardService,
        private fb: FormBuilder,
        private router: Router
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

    
    addPaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status): void {
        this.paymentCardService
        .addPaymentCard(cardToken, maskedPan, expiryMonth, expiryYear, cardholderName, Customer, Account, Tokenizations, Disputes, Scheme, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPaymentCard']);
            });
    }

    ngOnInit(): void {
    }
}
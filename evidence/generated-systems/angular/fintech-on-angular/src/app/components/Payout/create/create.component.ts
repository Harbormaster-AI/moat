import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PayoutService } from '../../../services/Payout.service';
import { Payout } from '../../../models/Payout';
import { SubBaseComponent } from '../../Payout/sub.base.component';

@Component({
    selector: 'app-create-payout',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePayoutComponent extends SubBaseComponent implements OnInit {

    title = 'Add Payout';

    payoutForm: FormGroup;
    payout: Payout;

    constructor( http: HttpClient,
        private payoutService: PayoutService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.payoutForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  payoutReference: ['', Validators.required],
      amount: ['', Validators.required],
      currency: ['', Validators.required],
      scheduledDate: ['', Validators.required],
      paidDate: ['', Validators.required],
      Merchant: ['', ],
      SettlementBatch: ['', ],
      DestinationAccount: ['', ],
      Status: ['', ]
        });
    }

    
    addPayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status): void {
        this.payoutService
        .addPayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPayout']);
            });
    }

    ngOnInit(): void {
    }
}
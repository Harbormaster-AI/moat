import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PayoutService } from '../../../services/Payout.service';
import { SubBaseComponent } from '../../Payout/sub.base.component';


@Component({
    selector: 'app-edit-payout',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPayoutComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Payout';

    payoutForm: FormGroup;
    payout: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PayoutService,
        private fb: FormBuilder
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

    
    updatePayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePayout(payoutReference, amount, currency, scheduledDate, paidDate, Merchant, SettlementBatch, DestinationAccount, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPayout']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPayout(params['id']).subscribe(res => {
                this.payout = res;
            });
        });
    }
}
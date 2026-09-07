import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FXDealService } from '../../../services/FXDeal.service';
import { SubBaseComponent } from '../../FXDeal/sub.base.component';


@Component({
    selector: 'app-edit-fXDeal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFXDealComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FXDeal';

    fXDealForm: FormGroup;
    fXDeal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FXDealService,
        private fb: FormBuilder
) {
        super(http);
        this.fXDealForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  dealReference: ['', Validators.required],
      baseCurrency: ['', Validators.required],
      quoteCurrency: ['', Validators.required],
      rate: ['', Validators.required],
      amount: ['', Validators.required],
      settlementDate: ['', Validators.required],
      Quote: ['', ],
      PaymentOrders: ['', ],
      Status: ['', ]
        });
    }

    
    updateFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFXDeal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFXDeal(params['id']).subscribe(res => {
                this.fXDeal = res;
            });
        });
    }
}
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FXDealService } from '../../../services/FXDeal.service';
import { FXDeal } from '../../../models/FXDeal';
import { SubBaseComponent } from '../../FXDeal/sub.base.component';

@Component({
    selector: 'app-create-fXDeal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFXDealComponent extends SubBaseComponent implements OnInit {

    title = 'Add FXDeal';

    fXDealForm: FormGroup;
    fXDeal: FXDeal;

    constructor( http: HttpClient,
        private fXDealService: FXDealService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status): void {
        this.fXDealService
        .addFXDeal(dealReference, baseCurrency, quoteCurrency, rate, amount, settlementDate, Quote, PaymentOrders, Status)
            .subscribe(() => {
                this.router.navigate(['/indexFXDeal']);
            });
    }

    ngOnInit(): void {
    }
}
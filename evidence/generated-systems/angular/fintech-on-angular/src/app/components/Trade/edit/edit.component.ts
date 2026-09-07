import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TradeService } from '../../../services/Trade.service';
import { SubBaseComponent } from '../../Trade/sub.base.component';


@Component({
    selector: 'app-edit-trade',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTradeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Trade';

    tradeForm: FormGroup;
    trade: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TradeService,
        private fb: FormBuilder
) {
        super(http);
        this.tradeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  executedAt: ['', Validators.required],
      quantity: ['', Validators.required],
      price: ['', Validators.required],
      fees: ['', Validators.required],
      settlementDate: ['', Validators.required],
      Order: ['', ],
      Security: ['', ],
      InvestmentAccount: ['', ]
        });
    }

    
    updateTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTrade']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTrade(params['id']).subscribe(res => {
                this.trade = res;
            });
        });
    }
}
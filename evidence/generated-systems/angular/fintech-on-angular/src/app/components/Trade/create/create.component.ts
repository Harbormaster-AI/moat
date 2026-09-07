import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TradeService } from '../../../services/Trade.service';
import { Trade } from '../../../models/Trade';
import { SubBaseComponent } from '../../Trade/sub.base.component';

@Component({
    selector: 'app-create-trade',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTradeComponent extends SubBaseComponent implements OnInit {

    title = 'Add Trade';

    tradeForm: FormGroup;
    trade: Trade;

    constructor( http: HttpClient,
        private tradeService: TradeService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount): void {
        this.tradeService
        .addTrade(executedAt, quantity, price, fees, settlementDate, Order, Security, InvestmentAccount)
            .subscribe(() => {
                this.router.navigate(['/indexTrade']);
            });
    }

    ngOnInit(): void {
    }
}
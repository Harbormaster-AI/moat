import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TradeOrderService } from '../../../services/TradeOrder.service';
import { TradeOrder } from '../../../models/TradeOrder';
import { SubBaseComponent } from '../../TradeOrder/sub.base.component';

@Component({
    selector: 'app-create-tradeOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTradeOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add TradeOrder';

    tradeOrderForm: FormGroup;
    tradeOrder: TradeOrder;

    constructor( http: HttpClient,
        private tradeOrderService: TradeOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.tradeOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderId: ['', Validators.required],
      quantity: ['', Validators.required],
      limitPrice: ['', Validators.required],
      placedAt: ['', Validators.required],
      Portfolio: ['', ],
      Security: ['', ],
      Trades: ['', ],
      Side: ['', ],
      Type: ['', ],
      Status: ['', ],
      TimeInForce: ['', ]
        });
    }

    
    addTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce): void {
        this.tradeOrderService
        .addTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce)
            .subscribe(() => {
                this.router.navigate(['/indexTradeOrder']);
            });
    }

    ngOnInit(): void {
    }
}
import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TradeOrderService } from '../../../services/TradeOrder.service';
import { SubBaseComponent } from '../../TradeOrder/sub.base.component';


@Component({
    selector: 'app-edit-tradeOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTradeOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TradeOrder';

    tradeOrderForm: FormGroup;
    tradeOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TradeOrderService,
        private fb: FormBuilder
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

    
    updateTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTradeOrder(orderId, quantity, limitPrice, placedAt, Portfolio, Security, Trades, Side, Type, Status, TimeInForce, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTradeOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTradeOrder(params['id']).subscribe(res => {
                this.tradeOrder = res;
            });
        });
    }
}
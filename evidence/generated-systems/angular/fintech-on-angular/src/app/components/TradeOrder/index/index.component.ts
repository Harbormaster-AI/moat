
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TradeOrderService } from '../../../services/TradeOrder.service';
import { TradeOrder } from '../../../models/TradeOrder';

@Component({
    selector: 'app-index-tradeOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTradeOrderComponent implements OnInit {

    tradeOrders: TradeOrder[] = [];

    constructor(
        private router: Router,
        private service: TradeOrderService
) {}

    ngOnInit(): void {
        this.getTradeOrders();
}

    getTradeOrders(): void {
        this.service.getTradeOrders().subscribe((res) => {
        this.tradeOrders = res;
    });
}

    deleteTradeOrder(id: any): void {
        this.service.deleteTradeOrder(id)
            .subscribe(() => {
                this.getTradeOrders();
            });
    }
}
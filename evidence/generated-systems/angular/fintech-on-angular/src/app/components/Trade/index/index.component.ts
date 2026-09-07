
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TradeService } from '../../../services/Trade.service';
import { Trade } from '../../../models/Trade';

@Component({
    selector: 'app-index-trade',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTradeComponent implements OnInit {

    trades: Trade[] = [];

    constructor(
        private router: Router,
        private service: TradeService
) {}

    ngOnInit(): void {
        this.getTrades();
}

    getTrades(): void {
        this.service.getTrades().subscribe((res) => {
        this.trades = res;
    });
}

    deleteTrade(id: any): void {
        this.service.deleteTrade(id)
            .subscribe(() => {
                this.getTrades();
            });
    }
}
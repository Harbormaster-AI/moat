
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PayoutService } from '../../../services/Payout.service';
import { Payout } from '../../../models/Payout';

@Component({
    selector: 'app-index-payout',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPayoutComponent implements OnInit {

    payouts: Payout[] = [];

    constructor(
        private router: Router,
        private service: PayoutService
) {}

    ngOnInit(): void {
        this.getPayouts();
}

    getPayouts(): void {
        this.service.getPayouts().subscribe((res) => {
        this.payouts = res;
    });
}

    deletePayout(id: any): void {
        this.service.deletePayout(id)
            .subscribe(() => {
                this.getPayouts();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DealService } from '../../../services/Deal.service';
import { Deal } from '../../../models/Deal';

@Component({
    selector: 'app-index-deal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDealComponent implements OnInit {

    deals: Deal[] = [];

    constructor(
        private router: Router,
        private service: DealService
) {}

    ngOnInit(): void {
        this.getDeals();
}

    getDeals(): void {
        this.service.getDeals().subscribe((res) => {
        this.deals = res;
    });
}

    deleteDeal(id: any): void {
        this.service.deleteDeal(id)
            .subscribe(() => {
                this.getDeals();
            });
    }
}
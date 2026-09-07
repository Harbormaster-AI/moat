
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FXDealService } from '../../../services/FXDeal.service';
import { FXDeal } from '../../../models/FXDeal';

@Component({
    selector: 'app-index-fXDeal',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFXDealComponent implements OnInit {

    fXDeals: FXDeal[] = [];

    constructor(
        private router: Router,
        private service: FXDealService
) {}

    ngOnInit(): void {
        this.getFXDeals();
}

    getFXDeals(): void {
        this.service.getFXDeals().subscribe((res) => {
        this.fXDeals = res;
    });
}

    deleteFXDeal(id: any): void {
        this.service.deleteFXDeal(id)
            .subscribe(() => {
                this.getFXDeals();
            });
    }
}
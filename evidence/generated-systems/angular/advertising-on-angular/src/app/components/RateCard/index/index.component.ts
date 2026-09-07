
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RateCardService } from '../../../services/RateCard.service';
import { RateCard } from '../../../models/RateCard';

@Component({
    selector: 'app-index-rateCard',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRateCardComponent implements OnInit {

    rateCards: RateCard[] = [];

    constructor(
        private router: Router,
        private service: RateCardService
) {}

    ngOnInit(): void {
        this.getRateCards();
}

    getRateCards(): void {
        this.service.getRateCards().subscribe((res) => {
        this.rateCards = res;
    });
}

    deleteRateCard(id: any): void {
        this.service.deleteRateCard(id)
            .subscribe(() => {
                this.getRateCards();
            });
    }
}
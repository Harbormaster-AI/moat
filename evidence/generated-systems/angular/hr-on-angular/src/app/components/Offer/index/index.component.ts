
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OfferService } from '../../../services/Offer.service';
import { Offer } from '../../../models/Offer';

@Component({
    selector: 'app-index-offer',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOfferComponent implements OnInit {

    offers: Offer[] = [];

    constructor(
        private router: Router,
        private service: OfferService
) {}

    ngOnInit(): void {
        this.getOffers();
}

    getOffers(): void {
        this.service.getOffers().subscribe((res) => {
        this.offers = res;
    });
}

    deleteOffer(id: any): void {
        this.service.deleteOffer(id)
            .subscribe(() => {
                this.getOffers();
            });
    }
}
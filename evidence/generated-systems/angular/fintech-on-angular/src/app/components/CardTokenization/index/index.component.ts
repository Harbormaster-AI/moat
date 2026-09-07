
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CardTokenizationService } from '../../../services/CardTokenization.service';
import { CardTokenization } from '../../../models/CardTokenization';

@Component({
    selector: 'app-index-cardTokenization',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCardTokenizationComponent implements OnInit {

    cardTokenizations: CardTokenization[] = [];

    constructor(
        private router: Router,
        private service: CardTokenizationService
) {}

    ngOnInit(): void {
        this.getCardTokenizations();
}

    getCardTokenizations(): void {
        this.service.getCardTokenizations().subscribe((res) => {
        this.cardTokenizations = res;
    });
}

    deleteCardTokenization(id: any): void {
        this.service.deleteCardTokenization(id)
            .subscribe(() => {
                this.getCardTokenizations();
            });
    }
}
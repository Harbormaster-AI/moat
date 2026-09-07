
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RateService } from '../../../services/Rate.service';
import { Rate } from '../../../models/Rate';

@Component({
    selector: 'app-index-rate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRateComponent implements OnInit {

    rates: Rate[] = [];

    constructor(
        private router: Router,
        private service: RateService
) {}

    ngOnInit(): void {
        this.getRates();
}

    getRates(): void {
        this.service.getRates().subscribe((res) => {
        this.rates = res;
    });
}

    deleteRate(id: any): void {
        this.service.deleteRate(id)
            .subscribe(() => {
                this.getRates();
            });
    }
}
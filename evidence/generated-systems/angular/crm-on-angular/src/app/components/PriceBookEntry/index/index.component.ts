
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PriceBookEntryService } from '../../../services/PriceBookEntry.service';
import { PriceBookEntry } from '../../../models/PriceBookEntry';

@Component({
    selector: 'app-index-priceBookEntry',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPriceBookEntryComponent implements OnInit {

    priceBookEntrys: PriceBookEntry[] = [];

    constructor(
        private router: Router,
        private service: PriceBookEntryService
) {}

    ngOnInit(): void {
        this.getPriceBookEntrys();
}

    getPriceBookEntrys(): void {
        this.service.getPriceBookEntrys().subscribe((res) => {
        this.priceBookEntrys = res;
    });
}

    deletePriceBookEntry(id: any): void {
        this.service.deletePriceBookEntry(id)
            .subscribe(() => {
                this.getPriceBookEntrys();
            });
    }
}
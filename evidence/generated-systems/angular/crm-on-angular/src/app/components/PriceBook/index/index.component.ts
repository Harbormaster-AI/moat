
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PriceBookService } from '../../../services/PriceBook.service';
import { PriceBook } from '../../../models/PriceBook';

@Component({
    selector: 'app-index-priceBook',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPriceBookComponent implements OnInit {

    priceBooks: PriceBook[] = [];

    constructor(
        private router: Router,
        private service: PriceBookService
) {}

    ngOnInit(): void {
        this.getPriceBooks();
}

    getPriceBooks(): void {
        this.service.getPriceBooks().subscribe((res) => {
        this.priceBooks = res;
    });
}

    deletePriceBook(id: any): void {
        this.service.deletePriceBook(id)
            .subscribe(() => {
                this.getPriceBooks();
            });
    }
}
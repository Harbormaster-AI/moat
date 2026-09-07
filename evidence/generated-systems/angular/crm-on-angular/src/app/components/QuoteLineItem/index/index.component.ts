
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';
import { QuoteLineItem } from '../../../models/QuoteLineItem';

@Component({
    selector: 'app-index-quoteLineItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQuoteLineItemComponent implements OnInit {

    quoteLineItems: QuoteLineItem[] = [];

    constructor(
        private router: Router,
        private service: QuoteLineItemService
) {}

    ngOnInit(): void {
        this.getQuoteLineItems();
}

    getQuoteLineItems(): void {
        this.service.getQuoteLineItems().subscribe((res) => {
        this.quoteLineItems = res;
    });
}

    deleteQuoteLineItem(id: any): void {
        this.service.deleteQuoteLineItem(id)
            .subscribe(() => {
                this.getQuoteLineItems();
            });
    }
}
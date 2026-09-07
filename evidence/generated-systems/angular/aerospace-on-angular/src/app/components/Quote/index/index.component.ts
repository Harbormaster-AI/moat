
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QuoteService } from '../../../services/Quote.service';
import { Quote } from '../../../models/Quote';

@Component({
    selector: 'app-index-quote',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQuoteComponent implements OnInit {

    quotes: Quote[] = [];

    constructor(
        private router: Router,
        private service: QuoteService
) {}

    ngOnInit(): void {
        this.getQuotes();
}

    getQuotes(): void {
        this.service.getQuotes().subscribe((res) => {
        this.quotes = res;
    });
}

    deleteQuote(id: any): void {
        this.service.deleteQuote(id)
            .subscribe(() => {
                this.getQuotes();
            });
    }
}
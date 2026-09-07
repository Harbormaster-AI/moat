
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FXQuoteService } from '../../../services/FXQuote.service';
import { FXQuote } from '../../../models/FXQuote';

@Component({
    selector: 'app-index-fXQuote',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFXQuoteComponent implements OnInit {

    fXQuotes: FXQuote[] = [];

    constructor(
        private router: Router,
        private service: FXQuoteService
) {}

    ngOnInit(): void {
        this.getFXQuotes();
}

    getFXQuotes(): void {
        this.service.getFXQuotes().subscribe((res) => {
        this.fXQuotes = res;
    });
}

    deleteFXQuote(id: any): void {
        this.service.deleteFXQuote(id)
            .subscribe(() => {
                this.getFXQuotes();
            });
    }
}
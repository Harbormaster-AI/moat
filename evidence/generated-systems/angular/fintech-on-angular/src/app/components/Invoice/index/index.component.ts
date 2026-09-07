
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InvoiceService } from '../../../services/Invoice.service';
import { Invoice } from '../../../models/Invoice';

@Component({
    selector: 'app-index-invoice',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInvoiceComponent implements OnInit {

    invoices: Invoice[] = [];

    constructor(
        private router: Router,
        private service: InvoiceService
) {}

    ngOnInit(): void {
        this.getInvoices();
}

    getInvoices(): void {
        this.service.getInvoices().subscribe((res) => {
        this.invoices = res;
    });
}

    deleteInvoice(id: any): void {
        this.service.deleteInvoice(id)
            .subscribe(() => {
                this.getInvoices();
            });
    }
}
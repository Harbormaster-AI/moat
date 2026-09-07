
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SalesOrderLineService } from '../../../services/SalesOrderLine.service';
import { SalesOrderLine } from '../../../models/SalesOrderLine';

@Component({
    selector: 'app-index-salesOrderLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSalesOrderLineComponent implements OnInit {

    salesOrderLines: SalesOrderLine[] = [];

    constructor(
        private router: Router,
        private service: SalesOrderLineService
) {}

    ngOnInit(): void {
        this.getSalesOrderLines();
}

    getSalesOrderLines(): void {
        this.service.getSalesOrderLines().subscribe((res) => {
        this.salesOrderLines = res;
    });
}

    deleteSalesOrderLine(id: any): void {
        this.service.deleteSalesOrderLine(id)
            .subscribe(() => {
                this.getSalesOrderLines();
            });
    }
}
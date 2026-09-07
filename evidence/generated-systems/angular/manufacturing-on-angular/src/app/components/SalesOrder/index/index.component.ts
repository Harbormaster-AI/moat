
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SalesOrderService } from '../../../services/SalesOrder.service';
import { SalesOrder } from '../../../models/SalesOrder';

@Component({
    selector: 'app-index-salesOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSalesOrderComponent implements OnInit {

    salesOrders: SalesOrder[] = [];

    constructor(
        private router: Router,
        private service: SalesOrderService
) {}

    ngOnInit(): void {
        this.getSalesOrders();
}

    getSalesOrders(): void {
        this.service.getSalesOrders().subscribe((res) => {
        this.salesOrders = res;
    });
}

    deleteSalesOrder(id: any): void {
        this.service.deleteSalesOrder(id)
            .subscribe(() => {
                this.getSalesOrders();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PurchaseOrderService } from '../../../services/PurchaseOrder.service';
import { PurchaseOrder } from '../../../models/PurchaseOrder';

@Component({
    selector: 'app-index-purchaseOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPurchaseOrderComponent implements OnInit {

    purchaseOrders: PurchaseOrder[] = [];

    constructor(
        private router: Router,
        private service: PurchaseOrderService
) {}

    ngOnInit(): void {
        this.getPurchaseOrders();
}

    getPurchaseOrders(): void {
        this.service.getPurchaseOrders().subscribe((res) => {
        this.purchaseOrders = res;
    });
}

    deletePurchaseOrder(id: any): void {
        this.service.deletePurchaseOrder(id)
            .subscribe(() => {
                this.getPurchaseOrders();
            });
    }
}
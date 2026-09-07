
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductionOrderService } from '../../../services/ProductionOrder.service';
import { ProductionOrder } from '../../../models/ProductionOrder';

@Component({
    selector: 'app-index-productionOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductionOrderComponent implements OnInit {

    productionOrders: ProductionOrder[] = [];

    constructor(
        private router: Router,
        private service: ProductionOrderService
) {}

    ngOnInit(): void {
        this.getProductionOrders();
}

    getProductionOrders(): void {
        this.service.getProductionOrders().subscribe((res) => {
        this.productionOrders = res;
    });
}

    deleteProductionOrder(id: any): void {
        this.service.deleteProductionOrder(id)
            .subscribe(() => {
                this.getProductionOrders();
            });
    }
}
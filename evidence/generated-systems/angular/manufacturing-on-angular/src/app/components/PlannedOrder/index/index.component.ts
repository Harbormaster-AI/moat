
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PlannedOrderService } from '../../../services/PlannedOrder.service';
import { PlannedOrder } from '../../../models/PlannedOrder';

@Component({
    selector: 'app-index-plannedOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPlannedOrderComponent implements OnInit {

    plannedOrders: PlannedOrder[] = [];

    constructor(
        private router: Router,
        private service: PlannedOrderService
) {}

    ngOnInit(): void {
        this.getPlannedOrders();
}

    getPlannedOrders(): void {
        this.service.getPlannedOrders().subscribe((res) => {
        this.plannedOrders = res;
    });
}

    deletePlannedOrder(id: any): void {
        this.service.deletePlannedOrder(id)
            .subscribe(() => {
                this.getPlannedOrders();
            });
    }
}
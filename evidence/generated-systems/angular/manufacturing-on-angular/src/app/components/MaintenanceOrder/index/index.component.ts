
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';
import { MaintenanceOrder } from '../../../models/MaintenanceOrder';

@Component({
    selector: 'app-index-maintenanceOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMaintenanceOrderComponent implements OnInit {

    maintenanceOrders: MaintenanceOrder[] = [];

    constructor(
        private router: Router,
        private service: MaintenanceOrderService
) {}

    ngOnInit(): void {
        this.getMaintenanceOrders();
}

    getMaintenanceOrders(): void {
        this.service.getMaintenanceOrders().subscribe((res) => {
        this.maintenanceOrders = res;
    });
}

    deleteMaintenanceOrder(id: any): void {
        this.service.deleteMaintenanceOrder(id)
            .subscribe(() => {
                this.getMaintenanceOrders();
            });
    }
}
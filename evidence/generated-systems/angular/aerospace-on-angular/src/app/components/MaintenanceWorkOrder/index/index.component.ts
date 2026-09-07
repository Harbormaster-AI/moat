
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';
import { MaintenanceWorkOrder } from '../../../models/MaintenanceWorkOrder';

@Component({
    selector: 'app-index-maintenanceWorkOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMaintenanceWorkOrderComponent implements OnInit {

    maintenanceWorkOrders: MaintenanceWorkOrder[] = [];

    constructor(
        private router: Router,
        private service: MaintenanceWorkOrderService
) {}

    ngOnInit(): void {
        this.getMaintenanceWorkOrders();
}

    getMaintenanceWorkOrders(): void {
        this.service.getMaintenanceWorkOrders().subscribe((res) => {
        this.maintenanceWorkOrders = res;
    });
}

    deleteMaintenanceWorkOrder(id: any): void {
        this.service.deleteMaintenanceWorkOrder(id)
            .subscribe(() => {
                this.getMaintenanceWorkOrders();
            });
    }
}
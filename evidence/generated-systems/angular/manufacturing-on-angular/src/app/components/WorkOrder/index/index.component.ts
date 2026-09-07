
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WorkOrderService } from '../../../services/WorkOrder.service';
import { WorkOrder } from '../../../models/WorkOrder';

@Component({
    selector: 'app-index-workOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWorkOrderComponent implements OnInit {

    workOrders: WorkOrder[] = [];

    constructor(
        private router: Router,
        private service: WorkOrderService
) {}

    ngOnInit(): void {
        this.getWorkOrders();
}

    getWorkOrders(): void {
        this.service.getWorkOrders().subscribe((res) => {
        this.workOrders = res;
    });
}

    deleteWorkOrder(id: any): void {
        this.service.deleteWorkOrder(id)
            .subscribe(() => {
                this.getWorkOrders();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';
import { LaboratoryOrder } from '../../../models/LaboratoryOrder';

@Component({
    selector: 'app-index-laboratoryOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLaboratoryOrderComponent implements OnInit {

    laboratoryOrders: LaboratoryOrder[] = [];

    constructor(
        private router: Router,
        private service: LaboratoryOrderService
) {}

    ngOnInit(): void {
        this.getLaboratoryOrders();
}

    getLaboratoryOrders(): void {
        this.service.getLaboratoryOrders().subscribe((res) => {
        this.laboratoryOrders = res;
    });
}

    deleteLaboratoryOrder(id: any): void {
        this.service.deleteLaboratoryOrder(id)
            .subscribe(() => {
                this.getLaboratoryOrders();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MedicationOrderService } from '../../../services/MedicationOrder.service';
import { MedicationOrder } from '../../../models/MedicationOrder';

@Component({
    selector: 'app-index-medicationOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMedicationOrderComponent implements OnInit {

    medicationOrders: MedicationOrder[] = [];

    constructor(
        private router: Router,
        private service: MedicationOrderService
) {}

    ngOnInit(): void {
        this.getMedicationOrders();
}

    getMedicationOrders(): void {
        this.service.getMedicationOrders().subscribe((res) => {
        this.medicationOrders = res;
    });
}

    deleteMedicationOrder(id: any): void {
        this.service.deleteMedicationOrder(id)
            .subscribe(() => {
                this.getMedicationOrders();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ClinicalOrderService } from '../../../services/ClinicalOrder.service';
import { ClinicalOrder } from '../../../models/ClinicalOrder';

@Component({
    selector: 'app-index-clinicalOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexClinicalOrderComponent implements OnInit {

    clinicalOrders: ClinicalOrder[] = [];

    constructor(
        private router: Router,
        private service: ClinicalOrderService
) {}

    ngOnInit(): void {
        this.getClinicalOrders();
}

    getClinicalOrders(): void {
        this.service.getClinicalOrders().subscribe((res) => {
        this.clinicalOrders = res;
    });
}

    deleteClinicalOrder(id: any): void {
        this.service.deleteClinicalOrder(id)
            .subscribe(() => {
                this.getClinicalOrders();
            });
    }
}
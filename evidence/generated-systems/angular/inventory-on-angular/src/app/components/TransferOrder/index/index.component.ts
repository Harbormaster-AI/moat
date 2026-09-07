
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TransferOrderService } from '../../../services/TransferOrder.service';
import { TransferOrder } from '../../../models/TransferOrder';

@Component({
    selector: 'app-index-transferOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTransferOrderComponent implements OnInit {

    transferOrders: TransferOrder[] = [];

    constructor(
        private router: Router,
        private service: TransferOrderService
) {}

    ngOnInit(): void {
        this.getTransferOrders();
}

    getTransferOrders(): void {
        this.service.getTransferOrders().subscribe((res) => {
        this.transferOrders = res;
    });
}

    deleteTransferOrder(id: any): void {
        this.service.deleteTransferOrder(id)
            .subscribe(() => {
                this.getTransferOrders();
            });
    }
}
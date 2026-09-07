
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OrderService } from '../../../services/Order.service';
import { Order } from '../../../models/Order';

@Component({
    selector: 'app-index-order',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOrderComponent implements OnInit {

    orders: Order[] = [];

    constructor(
        private router: Router,
        private service: OrderService
) {}

    ngOnInit(): void {
        this.getOrders();
}

    getOrders(): void {
        this.service.getOrders().subscribe((res) => {
        this.orders = res;
    });
}

    deleteOrder(id: any): void {
        this.service.deleteOrder(id)
            .subscribe(() => {
                this.getOrders();
            });
    }
}
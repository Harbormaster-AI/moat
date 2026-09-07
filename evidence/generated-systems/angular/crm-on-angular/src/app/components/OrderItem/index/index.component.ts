
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { OrderItemService } from '../../../services/OrderItem.service';
import { OrderItem } from '../../../models/OrderItem';

@Component({
    selector: 'app-index-orderItem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexOrderItemComponent implements OnInit {

    orderItems: OrderItem[] = [];

    constructor(
        private router: Router,
        private service: OrderItemService
) {}

    ngOnInit(): void {
        this.getOrderItems();
}

    getOrderItems(): void {
        this.service.getOrderItems().subscribe((res) => {
        this.orderItems = res;
    });
}

    deleteOrderItem(id: any): void {
        this.service.deleteOrderItem(id)
            .subscribe(() => {
                this.getOrderItems();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsertionOrderService } from '../../../services/InsertionOrder.service';
import { InsertionOrder } from '../../../models/InsertionOrder';

@Component({
    selector: 'app-index-insertionOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsertionOrderComponent implements OnInit {

    insertionOrders: InsertionOrder[] = [];

    constructor(
        private router: Router,
        private service: InsertionOrderService
) {}

    ngOnInit(): void {
        this.getInsertionOrders();
}

    getInsertionOrders(): void {
        this.service.getInsertionOrders().subscribe((res) => {
        this.insertionOrders = res;
    });
}

    deleteInsertionOrder(id: any): void {
        this.service.deleteInsertionOrder(id)
            .subscribe(() => {
                this.getInsertionOrders();
            });
    }
}
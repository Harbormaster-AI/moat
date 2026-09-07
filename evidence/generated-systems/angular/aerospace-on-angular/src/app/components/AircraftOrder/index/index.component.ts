
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AircraftOrderService } from '../../../services/AircraftOrder.service';
import { AircraftOrder } from '../../../models/AircraftOrder';

@Component({
    selector: 'app-index-aircraftOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAircraftOrderComponent implements OnInit {

    aircraftOrders: AircraftOrder[] = [];

    constructor(
        private router: Router,
        private service: AircraftOrderService
) {}

    ngOnInit(): void {
        this.getAircraftOrders();
}

    getAircraftOrders(): void {
        this.service.getAircraftOrders().subscribe((res) => {
        this.aircraftOrders = res;
    });
}

    deleteAircraftOrder(id: any): void {
        this.service.deleteAircraftOrder(id)
            .subscribe(() => {
                this.getAircraftOrders();
            });
    }
}
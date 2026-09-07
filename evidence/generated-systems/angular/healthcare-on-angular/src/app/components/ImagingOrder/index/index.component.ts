
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ImagingOrderService } from '../../../services/ImagingOrder.service';
import { ImagingOrder } from '../../../models/ImagingOrder';

@Component({
    selector: 'app-index-imagingOrder',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexImagingOrderComponent implements OnInit {

    imagingOrders: ImagingOrder[] = [];

    constructor(
        private router: Router,
        private service: ImagingOrderService
) {}

    ngOnInit(): void {
        this.getImagingOrders();
}

    getImagingOrders(): void {
        this.service.getImagingOrders().subscribe((res) => {
        this.imagingOrders = res;
    });
}

    deleteImagingOrder(id: any): void {
        this.service.deleteImagingOrder(id)
            .subscribe(() => {
                this.getImagingOrders();
            });
    }
}
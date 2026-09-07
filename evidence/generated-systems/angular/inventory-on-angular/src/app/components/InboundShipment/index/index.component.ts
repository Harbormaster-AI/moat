
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InboundShipmentService } from '../../../services/InboundShipment.service';
import { InboundShipment } from '../../../models/InboundShipment';

@Component({
    selector: 'app-index-inboundShipment',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInboundShipmentComponent implements OnInit {

    inboundShipments: InboundShipment[] = [];

    constructor(
        private router: Router,
        private service: InboundShipmentService
) {}

    ngOnInit(): void {
        this.getInboundShipments();
}

    getInboundShipments(): void {
        this.service.getInboundShipments().subscribe((res) => {
        this.inboundShipments = res;
    });
}

    deleteInboundShipment(id: any): void {
        this.service.deleteInboundShipment(id)
            .subscribe(() => {
                this.getInboundShipments();
            });
    }
}
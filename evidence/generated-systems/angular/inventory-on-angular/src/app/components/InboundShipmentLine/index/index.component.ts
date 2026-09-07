
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';
import { InboundShipmentLine } from '../../../models/InboundShipmentLine';

@Component({
    selector: 'app-index-inboundShipmentLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInboundShipmentLineComponent implements OnInit {

    inboundShipmentLines: InboundShipmentLine[] = [];

    constructor(
        private router: Router,
        private service: InboundShipmentLineService
) {}

    ngOnInit(): void {
        this.getInboundShipmentLines();
}

    getInboundShipmentLines(): void {
        this.service.getInboundShipmentLines().subscribe((res) => {
        this.inboundShipmentLines = res;
    });
}

    deleteInboundShipmentLine(id: any): void {
        this.service.deleteInboundShipmentLine(id)
            .subscribe(() => {
                this.getInboundShipmentLines();
            });
    }
}
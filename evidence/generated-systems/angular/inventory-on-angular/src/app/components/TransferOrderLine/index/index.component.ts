
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';
import { TransferOrderLine } from '../../../models/TransferOrderLine';

@Component({
    selector: 'app-index-transferOrderLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTransferOrderLineComponent implements OnInit {

    transferOrderLines: TransferOrderLine[] = [];

    constructor(
        private router: Router,
        private service: TransferOrderLineService
) {}

    ngOnInit(): void {
        this.getTransferOrderLines();
}

    getTransferOrderLines(): void {
        this.service.getTransferOrderLines().subscribe((res) => {
        this.transferOrderLines = res;
    });
}

    deleteTransferOrderLine(id: any): void {
        this.service.deleteTransferOrderLine(id)
            .subscribe(() => {
                this.getTransferOrderLines();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SettlementBatchService } from '../../../services/SettlementBatch.service';
import { SettlementBatch } from '../../../models/SettlementBatch';

@Component({
    selector: 'app-index-settlementBatch',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSettlementBatchComponent implements OnInit {

    settlementBatchs: SettlementBatch[] = [];

    constructor(
        private router: Router,
        private service: SettlementBatchService
) {}

    ngOnInit(): void {
        this.getSettlementBatchs();
}

    getSettlementBatchs(): void {
        this.service.getSettlementBatchs().subscribe((res) => {
        this.settlementBatchs = res;
    });
}

    deleteSettlementBatch(id: any): void {
        this.service.deleteSettlementBatch(id)
            .subscribe(() => {
                this.getSettlementBatchs();
            });
    }
}
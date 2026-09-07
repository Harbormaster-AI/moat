
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProductionLineService } from '../../../services/ProductionLine.service';
import { ProductionLine } from '../../../models/ProductionLine';

@Component({
    selector: 'app-index-productionLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProductionLineComponent implements OnInit {

    productionLines: ProductionLine[] = [];

    constructor(
        private router: Router,
        private service: ProductionLineService
) {}

    ngOnInit(): void {
        this.getProductionLines();
}

    getProductionLines(): void {
        this.service.getProductionLines().subscribe((res) => {
        this.productionLines = res;
    });
}

    deleteProductionLine(id: any): void {
        this.service.deleteProductionLine(id)
            .subscribe(() => {
                this.getProductionLines();
            });
    }
}
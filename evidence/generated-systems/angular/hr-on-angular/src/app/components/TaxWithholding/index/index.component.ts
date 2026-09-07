
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TaxWithholdingService } from '../../../services/TaxWithholding.service';
import { TaxWithholding } from '../../../models/TaxWithholding';

@Component({
    selector: 'app-index-taxWithholding',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTaxWithholdingComponent implements OnInit {

    taxWithholdings: TaxWithholding[] = [];

    constructor(
        private router: Router,
        private service: TaxWithholdingService
) {}

    ngOnInit(): void {
        this.getTaxWithholdings();
}

    getTaxWithholdings(): void {
        this.service.getTaxWithholdings().subscribe((res) => {
        this.taxWithholdings = res;
    });
}

    deleteTaxWithholding(id: any): void {
        this.service.deleteTaxWithholding(id)
            .subscribe(() => {
                this.getTaxWithholdings();
            });
    }
}
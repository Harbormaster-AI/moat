
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WarrantyService } from '../../../services/Warranty.service';
import { Warranty } from '../../../models/Warranty';

@Component({
    selector: 'app-index-warranty',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWarrantyComponent implements OnInit {

    warrantys: Warranty[] = [];

    constructor(
        private router: Router,
        private service: WarrantyService
) {}

    ngOnInit(): void {
        this.getWarrantys();
}

    getWarrantys(): void {
        this.service.getWarrantys().subscribe((res) => {
        this.warrantys = res;
    });
}

    deleteWarranty(id: any): void {
        this.service.deleteWarranty(id)
            .subscribe(() => {
                this.getWarrantys();
            });
    }
}
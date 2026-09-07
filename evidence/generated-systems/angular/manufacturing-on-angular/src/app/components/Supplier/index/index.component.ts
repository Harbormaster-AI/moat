
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SupplierService } from '../../../services/Supplier.service';
import { Supplier } from '../../../models/Supplier';

@Component({
    selector: 'app-index-supplier',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSupplierComponent implements OnInit {

    suppliers: Supplier[] = [];

    constructor(
        private router: Router,
        private service: SupplierService
) {}

    ngOnInit(): void {
        this.getSuppliers();
}

    getSuppliers(): void {
        this.service.getSuppliers().subscribe((res) => {
        this.suppliers = res;
    });
}

    deleteSupplier(id: any): void {
        this.service.deleteSupplier(id)
            .subscribe(() => {
                this.getSuppliers();
            });
    }
}
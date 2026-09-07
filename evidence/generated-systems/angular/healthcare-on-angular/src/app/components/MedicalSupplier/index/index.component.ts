
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MedicalSupplierService } from '../../../services/MedicalSupplier.service';
import { MedicalSupplier } from '../../../models/MedicalSupplier';

@Component({
    selector: 'app-index-medicalSupplier',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMedicalSupplierComponent implements OnInit {

    medicalSuppliers: MedicalSupplier[] = [];

    constructor(
        private router: Router,
        private service: MedicalSupplierService
) {}

    ngOnInit(): void {
        this.getMedicalSuppliers();
}

    getMedicalSuppliers(): void {
        this.service.getMedicalSuppliers().subscribe((res) => {
        this.medicalSuppliers = res;
    });
}

    deleteMedicalSupplier(id: any): void {
        this.service.deleteMedicalSupplier(id)
            .subscribe(() => {
                this.getMedicalSuppliers();
            });
    }
}
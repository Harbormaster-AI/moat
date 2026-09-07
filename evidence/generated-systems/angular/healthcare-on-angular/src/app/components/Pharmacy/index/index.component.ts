
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PharmacyService } from '../../../services/Pharmacy.service';
import { Pharmacy } from '../../../models/Pharmacy';

@Component({
    selector: 'app-index-pharmacy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPharmacyComponent implements OnInit {

    pharmacys: Pharmacy[] = [];

    constructor(
        private router: Router,
        private service: PharmacyService
) {}

    ngOnInit(): void {
        this.getPharmacys();
}

    getPharmacys(): void {
        this.service.getPharmacys().subscribe((res) => {
        this.pharmacys = res;
    });
}

    deletePharmacy(id: any): void {
        this.service.deletePharmacy(id)
            .subscribe(() => {
                this.getPharmacys();
            });
    }
}
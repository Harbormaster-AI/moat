
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InsurancePayerService } from '../../../services/InsurancePayer.service';
import { InsurancePayer } from '../../../models/InsurancePayer';

@Component({
    selector: 'app-index-insurancePayer',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInsurancePayerComponent implements OnInit {

    insurancePayers: InsurancePayer[] = [];

    constructor(
        private router: Router,
        private service: InsurancePayerService
) {}

    ngOnInit(): void {
        this.getInsurancePayers();
}

    getInsurancePayers(): void {
        this.service.getInsurancePayers().subscribe((res) => {
        this.insurancePayers = res;
    });
}

    deleteInsurancePayer(id: any): void {
        this.service.deleteInsurancePayer(id)
            .subscribe(() => {
                this.getInsurancePayers();
            });
    }
}
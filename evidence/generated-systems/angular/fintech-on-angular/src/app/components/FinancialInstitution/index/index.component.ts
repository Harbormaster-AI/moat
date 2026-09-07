
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FinancialInstitutionService } from '../../../services/FinancialInstitution.service';
import { FinancialInstitution } from '../../../models/FinancialInstitution';

@Component({
    selector: 'app-index-financialInstitution',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFinancialInstitutionComponent implements OnInit {

    financialInstitutions: FinancialInstitution[] = [];

    constructor(
        private router: Router,
        private service: FinancialInstitutionService
) {}

    ngOnInit(): void {
        this.getFinancialInstitutions();
}

    getFinancialInstitutions(): void {
        this.service.getFinancialInstitutions().subscribe((res) => {
        this.financialInstitutions = res;
    });
}

    deleteFinancialInstitution(id: any): void {
        this.service.deleteFinancialInstitution(id)
            .subscribe(() => {
                this.getFinancialInstitutions();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgreementService } from '../../../services/Agreement.service';
import { Agreement } from '../../../models/Agreement';

@Component({
    selector: 'app-index-agreement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAgreementComponent implements OnInit {

    agreements: Agreement[] = [];

    constructor(
        private router: Router,
        private service: AgreementService
) {}

    ngOnInit(): void {
        this.getAgreements();
}

    getAgreements(): void {
        this.service.getAgreements().subscribe((res) => {
        this.agreements = res;
    });
}

    deleteAgreement(id: any): void {
        this.service.deleteAgreement(id)
            .subscribe(() => {
                this.getAgreements();
            });
    }
}
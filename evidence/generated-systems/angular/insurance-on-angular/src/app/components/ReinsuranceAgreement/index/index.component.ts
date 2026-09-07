
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReinsuranceAgreementService } from '../../../services/ReinsuranceAgreement.service';
import { ReinsuranceAgreement } from '../../../models/ReinsuranceAgreement';

@Component({
    selector: 'app-index-reinsuranceAgreement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexReinsuranceAgreementComponent implements OnInit {

    reinsuranceAgreements: ReinsuranceAgreement[] = [];

    constructor(
        private router: Router,
        private service: ReinsuranceAgreementService
) {}

    ngOnInit(): void {
        this.getReinsuranceAgreements();
}

    getReinsuranceAgreements(): void {
        this.service.getReinsuranceAgreements().subscribe((res) => {
        this.reinsuranceAgreements = res;
    });
}

    deleteReinsuranceAgreement(id: any): void {
        this.service.deleteReinsuranceAgreement(id)
            .subscribe(() => {
                this.getReinsuranceAgreements();
            });
    }
}
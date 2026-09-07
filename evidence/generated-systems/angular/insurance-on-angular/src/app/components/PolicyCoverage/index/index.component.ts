
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';
import { PolicyCoverage } from '../../../models/PolicyCoverage';

@Component({
    selector: 'app-index-policyCoverage',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPolicyCoverageComponent implements OnInit {

    policyCoverages: PolicyCoverage[] = [];

    constructor(
        private router: Router,
        private service: PolicyCoverageService
) {}

    ngOnInit(): void {
        this.getPolicyCoverages();
}

    getPolicyCoverages(): void {
        this.service.getPolicyCoverages().subscribe((res) => {
        this.policyCoverages = res;
    });
}

    deletePolicyCoverage(id: any): void {
        this.service.deletePolicyCoverage(id)
            .subscribe(() => {
                this.getPolicyCoverages();
            });
    }
}
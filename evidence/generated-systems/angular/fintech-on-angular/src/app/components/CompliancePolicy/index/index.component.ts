
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompliancePolicyService } from '../../../services/CompliancePolicy.service';
import { CompliancePolicy } from '../../../models/CompliancePolicy';

@Component({
    selector: 'app-index-compliancePolicy',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCompliancePolicyComponent implements OnInit {

    compliancePolicys: CompliancePolicy[] = [];

    constructor(
        private router: Router,
        private service: CompliancePolicyService
) {}

    ngOnInit(): void {
        this.getCompliancePolicys();
}

    getCompliancePolicys(): void {
        this.service.getCompliancePolicys().subscribe((res) => {
        this.compliancePolicys = res;
    });
}

    deleteCompliancePolicy(id: any): void {
        this.service.deleteCompliancePolicy(id)
            .subscribe(() => {
                this.getCompliancePolicys();
            });
    }
}
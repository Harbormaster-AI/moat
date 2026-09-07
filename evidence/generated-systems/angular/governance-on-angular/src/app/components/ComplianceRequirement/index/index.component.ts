
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';
import { ComplianceRequirement } from '../../../models/ComplianceRequirement';

@Component({
    selector: 'app-index-complianceRequirement',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexComplianceRequirementComponent implements OnInit {

    complianceRequirements: ComplianceRequirement[] = [];

    constructor(
        private router: Router,
        private service: ComplianceRequirementService
) {}

    ngOnInit(): void {
        this.getComplianceRequirements();
}

    getComplianceRequirements(): void {
        this.service.getComplianceRequirements().subscribe((res) => {
        this.complianceRequirements = res;
    });
}

    deleteComplianceRequirement(id: any): void {
        this.service.deleteComplianceRequirement(id)
            .subscribe(() => {
                this.getComplianceRequirements();
            });
    }
}
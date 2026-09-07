
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ComplianceProgramService } from '../../../services/ComplianceProgram.service';
import { ComplianceProgram } from '../../../models/ComplianceProgram';

@Component({
    selector: 'app-index-complianceProgram',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexComplianceProgramComponent implements OnInit {

    compliancePrograms: ComplianceProgram[] = [];

    constructor(
        private router: Router,
        private service: ComplianceProgramService
) {}

    ngOnInit(): void {
        this.getCompliancePrograms();
}

    getCompliancePrograms(): void {
        this.service.getCompliancePrograms().subscribe((res) => {
        this.compliancePrograms = res;
    });
}

    deleteComplianceProgram(id: any): void {
        this.service.deleteComplianceProgram(id)
            .subscribe(() => {
                this.getCompliancePrograms();
            });
    }
}
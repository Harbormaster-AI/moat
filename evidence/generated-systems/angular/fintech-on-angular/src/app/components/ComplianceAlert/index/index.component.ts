
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';
import { ComplianceAlert } from '../../../models/ComplianceAlert';

@Component({
    selector: 'app-index-complianceAlert',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexComplianceAlertComponent implements OnInit {

    complianceAlerts: ComplianceAlert[] = [];

    constructor(
        private router: Router,
        private service: ComplianceAlertService
) {}

    ngOnInit(): void {
        this.getComplianceAlerts();
}

    getComplianceAlerts(): void {
        this.service.getComplianceAlerts().subscribe((res) => {
        this.complianceAlerts = res;
    });
}

    deleteComplianceAlert(id: any): void {
        this.service.deleteComplianceAlert(id)
            .subscribe(() => {
                this.getComplianceAlerts();
            });
    }
}
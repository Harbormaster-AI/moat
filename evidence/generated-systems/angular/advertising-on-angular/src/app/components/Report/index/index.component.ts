
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReportService } from '../../../services/Report.service';
import { Report } from '../../../models/Report';

@Component({
    selector: 'app-index-report',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexReportComponent implements OnInit {

    reports: Report[] = [];

    constructor(
        private router: Router,
        private service: ReportService
) {}

    ngOnInit(): void {
        this.getReports();
}

    getReports(): void {
        this.service.getReports().subscribe((res) => {
        this.reports = res;
    });
}

    deleteReport(id: any): void {
        this.service.deleteReport(id)
            .subscribe(() => {
                this.getReports();
            });
    }
}
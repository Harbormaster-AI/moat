
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ImagingReportService } from '../../../services/ImagingReport.service';
import { ImagingReport } from '../../../models/ImagingReport';

@Component({
    selector: 'app-index-imagingReport',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexImagingReportComponent implements OnInit {

    imagingReports: ImagingReport[] = [];

    constructor(
        private router: Router,
        private service: ImagingReportService
) {}

    ngOnInit(): void {
        this.getImagingReports();
}

    getImagingReports(): void {
        this.service.getImagingReports().subscribe((res) => {
        this.imagingReports = res;
    });
}

    deleteImagingReport(id: any): void {
        this.service.deleteImagingReport(id)
            .subscribe(() => {
                this.getImagingReports();
            });
    }
}
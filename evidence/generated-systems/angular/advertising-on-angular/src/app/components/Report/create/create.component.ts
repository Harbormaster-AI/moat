import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ReportService } from '../../../services/Report.service';
import { Report } from '../../../models/Report';
import { SubBaseComponent } from '../../Report/sub.base.component';

@Component({
    selector: 'app-create-report',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateReportComponent extends SubBaseComponent implements OnInit {

    title = 'Add Report';

    reportForm: FormGroup;
    report: Report;

    constructor( http: HttpClient,
        private reportService: ReportService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.reportForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  reportName: ['', Validators.required],
      generatedAt: ['', Validators.required],
      fileUrl: ['', Validators.required],
      AdAccount: ['', ],
      Campaign: ['', ],
      LineItem: ['', ],
      ReportType: ['', ]
        });
    }

    
    addReport(reportName, generatedAt, fileUrl, AdAccount, Campaign, LineItem, ReportType): void {
        this.reportService
        .addReport(reportName, generatedAt, fileUrl, AdAccount, Campaign, LineItem, ReportType)
            .subscribe(() => {
                this.router.navigate(['/indexReport']);
            });
    }

    ngOnInit(): void {
    }
}
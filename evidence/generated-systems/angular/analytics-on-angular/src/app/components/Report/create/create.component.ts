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
                  title: ['', Validators.required],
      audience: ['', Validators.required],
      Workspace: ['', ],
      Visualizations: ['', ],
      Datasets: ['', ],
      SemanticModels: ['', ],
      Queries: ['', ],
      Tags: ['', ],
      Status: ['', ]
        });
    }

    
    addReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status): void {
        this.reportService
        .addReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status)
            .subscribe(() => {
                this.router.navigate(['/indexReport']);
            });
    }

    ngOnInit(): void {
    }
}
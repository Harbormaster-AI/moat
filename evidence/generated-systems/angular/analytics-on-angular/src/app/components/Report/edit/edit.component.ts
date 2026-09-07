import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ReportService } from '../../../services/Report.service';
import { SubBaseComponent } from '../../Report/sub.base.component';


@Component({
    selector: 'app-edit-report',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditReportComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Report';

    reportForm: FormGroup;
    report: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ReportService,
        private fb: FormBuilder
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

    
    updateReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexReport']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getReport(params['id']).subscribe(res => {
                this.report = res;
            });
        });
    }
}
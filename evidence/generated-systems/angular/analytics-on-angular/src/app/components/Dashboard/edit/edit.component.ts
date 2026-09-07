import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DashboardService } from '../../../services/Dashboard.service';
import { SubBaseComponent } from '../../Dashboard/sub.base.component';


@Component({
    selector: 'app-edit-dashboard',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDashboardComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Dashboard';

    dashboardForm: FormGroup;
    dashboard: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DashboardService,
        private fb: FormBuilder
) {
        super(http);
        this.dashboardForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      theme: ['', Validators.required],
      Workspace: ['', ],
      Visualizations: ['', ],
      Reports: ['', ],
      Datasets: ['', ],
      Alerts: ['', ],
      Queries: ['', ],
      Tags: ['', ],
      Status: ['', ]
        });
    }

    
    updateDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDashboard']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDashboard(params['id']).subscribe(res => {
                this.dashboard = res;
            });
        });
    }
}
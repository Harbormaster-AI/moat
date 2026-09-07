import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DashboardService } from '../../../services/Dashboard.service';
import { Dashboard } from '../../../models/Dashboard';
import { SubBaseComponent } from '../../Dashboard/sub.base.component';

@Component({
    selector: 'app-create-dashboard',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDashboardComponent extends SubBaseComponent implements OnInit {

    title = 'Add Dashboard';

    dashboardForm: FormGroup;
    dashboard: Dashboard;

    constructor( http: HttpClient,
        private dashboardService: DashboardService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status): void {
        this.dashboardService
        .addDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status)
            .subscribe(() => {
                this.router.navigate(['/indexDashboard']);
            });
    }

    ngOnInit(): void {
    }
}
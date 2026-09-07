
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DashboardService } from '../../../services/Dashboard.service';
import { Dashboard } from '../../../models/Dashboard';

@Component({
    selector: 'app-index-dashboard',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDashboardComponent implements OnInit {

    dashboards: Dashboard[] = [];

    constructor(
        private router: Router,
        private service: DashboardService
) {}

    ngOnInit(): void {
        this.getDashboards();
}

    getDashboards(): void {
        this.service.getDashboards().subscribe((res) => {
        this.dashboards = res;
    });
}

    deleteDashboard(id: any): void {
        this.service.deleteDashboard(id)
            .subscribe(() => {
                this.getDashboards();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AnalyticsWorkspaceService } from '../../../services/AnalyticsWorkspace.service';
import { AnalyticsWorkspace } from '../../../models/AnalyticsWorkspace';

@Component({
    selector: 'app-index-analyticsWorkspace',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAnalyticsWorkspaceComponent implements OnInit {

    analyticsWorkspaces: AnalyticsWorkspace[] = [];

    constructor(
        private router: Router,
        private service: AnalyticsWorkspaceService
) {}

    ngOnInit(): void {
        this.getAnalyticsWorkspaces();
}

    getAnalyticsWorkspaces(): void {
        this.service.getAnalyticsWorkspaces().subscribe((res) => {
        this.analyticsWorkspaces = res;
    });
}

    deleteAnalyticsWorkspace(id: any): void {
        this.service.deleteAnalyticsWorkspace(id)
            .subscribe(() => {
                this.getAnalyticsWorkspaces();
            });
    }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApplicationService } from '../../../services/Application.service';
import { Application } from '../../../models/Application';

@Component({
    selector: 'app-index-application',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexApplicationComponent implements OnInit {

    applications: Application[] = [];

    constructor(
        private router: Router,
        private service: ApplicationService
) {}

    ngOnInit(): void {
        this.getApplications();
}

    getApplications(): void {
        this.service.getApplications().subscribe((res) => {
        this.applications = res;
    });
}

    deleteApplication(id: any): void {
        this.service.deleteApplication(id)
            .subscribe(() => {
                this.getApplications();
            });
    }
}
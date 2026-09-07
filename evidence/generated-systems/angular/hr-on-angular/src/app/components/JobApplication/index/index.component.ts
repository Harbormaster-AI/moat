
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobApplicationService } from '../../../services/JobApplication.service';
import { JobApplication } from '../../../models/JobApplication';

@Component({
    selector: 'app-index-jobApplication',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexJobApplicationComponent implements OnInit {

    jobApplications: JobApplication[] = [];

    constructor(
        private router: Router,
        private service: JobApplicationService
) {}

    ngOnInit(): void {
        this.getJobApplications();
}

    getJobApplications(): void {
        this.service.getJobApplications().subscribe((res) => {
        this.jobApplications = res;
    });
}

    deleteJobApplication(id: any): void {
        this.service.deleteJobApplication(id)
            .subscribe(() => {
                this.getJobApplications();
            });
    }
}
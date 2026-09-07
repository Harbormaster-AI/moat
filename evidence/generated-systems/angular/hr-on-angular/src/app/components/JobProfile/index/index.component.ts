
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobProfileService } from '../../../services/JobProfile.service';
import { JobProfile } from '../../../models/JobProfile';

@Component({
    selector: 'app-index-jobProfile',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexJobProfileComponent implements OnInit {

    jobProfiles: JobProfile[] = [];

    constructor(
        private router: Router,
        private service: JobProfileService
) {}

    ngOnInit(): void {
        this.getJobProfiles();
}

    getJobProfiles(): void {
        this.service.getJobProfiles().subscribe((res) => {
        this.jobProfiles = res;
    });
}

    deleteJobProfile(id: any): void {
        this.service.deleteJobProfile(id)
            .subscribe(() => {
                this.getJobProfiles();
            });
    }
}
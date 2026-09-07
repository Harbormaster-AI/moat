
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobFamilyService } from '../../../services/JobFamily.service';
import { JobFamily } from '../../../models/JobFamily';

@Component({
    selector: 'app-index-jobFamily',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexJobFamilyComponent implements OnInit {

    jobFamilys: JobFamily[] = [];

    constructor(
        private router: Router,
        private service: JobFamilyService
) {}

    ngOnInit(): void {
        this.getJobFamilys();
}

    getJobFamilys(): void {
        this.service.getJobFamilys().subscribe((res) => {
        this.jobFamilys = res;
    });
}

    deleteJobFamily(id: any): void {
        this.service.deleteJobFamily(id)
            .subscribe(() => {
                this.getJobFamilys();
            });
    }
}